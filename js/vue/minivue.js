function MiniVue(options) {
    let evalInContext = function (code, arguments) {
        arguments = arguments || {};
        with (this) {
            with (arguments) {
                return eval(code);
            }
        }
    }.bind(this);

    let evaluateProperty = function (key, value) {
        const prefix = key.slice(0, key.indexOf(":") + 1) || (key[0] === "@" ? "@" : "");
        key = key.slice(prefix.length);
        if (prefix === "v-bind:" || prefix === ":" || key === "v-if") {
            value = evalInContext(value);
        } else if (prefix === "v-on:" || prefix === "@") {
            key = "on" + key;
            let code = value; // otherwise value will remain the last
            value = (event) => evalInContext(code, { $event: event });
        }
        if (key === "class") key = "className";
        return [key, value];
    };

    this.buildVirtualElement = function (domNode) {
        if (domNode.nodeType == Node.ELEMENT_NODE) {
            let element = {
                tag: domNode.tagName,
                attrs: {},
                children: [],
            };
            for (let attr of domNode.attributes) {
                const [key, value] = evaluateProperty(attr.name, attr.value);
                if (key == "class") console.log(key, value);
                element.attrs[key] = value;
                if (key === "v-if" && !value) {
                    return null; // will not be added if v-if is false
                }
            }
            for (let child of domNode.childNodes) {
                const childElement = this.buildVirtualElement(child);
                if (childElement) {
                    element["children"].push(childElement);
                }
            }
            return element;
        } else if (domNode.nodeType == Node.TEXT_NODE) {
            return { tag: "TEXT", text: domNode.wholeText, children: [] };
        }
    };

    this.reconcile = function (parentDom, element, instance) {
        if (!instance) {
            instance = this.instantiate(element);
            parentDom.appendChild(instance.dom);
        } else if (!element) {
            instance.dom.remove();
            delete instance;
            return;
        } else if (element.tag === "TEXT" || instance.element.tag !== element.tag) {
            let newInstance = this.instantiate(element);
            parentDom.replaceChild(newInstance.dom, instance.dom);
            instance.element = newInstance.element;
            instance.dom = newInstance.dom;
            instance.children = newInstance.children; // got me!
        } else {
            updateDomProperties(element, instance.dom);
        }
        for (let i = 0; i < element.children.length; i++) {
            let [childElement, childInstance] = [element.children[i], instance.children[i]];
            this.reconcile(instance.dom, childElement, childInstance);
        }
        return instance;
    };

    let updateDomProperties = function (element, dom) {
        for (let key in element.attrs) {
            dom[key] = element.attrs[key];
        }
    };

    this.instantiate = function (element) {
        let instance = { element: element, dom: null, children: [] };
        if (element.tag !== "TEXT") {
            const dom = document.createElement(element.tag);
            updateDomProperties(element, dom);
            for (let child of element.children) {
                const childInstance = this.instantiate(child);
                instance.children.push(childInstance);
                dom.appendChild(childInstance.dom);
            }
            instance.dom = dom;
        } else {
            const pattern = /\{\{([^\}]+?)\}\}/g;
            const expressions = element.text.match(pattern) || [];
            for (let expression of expressions) {
                const unwrapped = expression.replace(pattern, "$1");
                const value = evalInContext(unwrapped);
                element.text = element.text.split(expression).join(value); // replace expression with value
            }
            instance.dom = document.createTextNode(element.text);
        }
        return instance;
    };

    const data = options.data || {};
    for (let key in data) {
        Object.defineProperty(this, key, {
            get() {
                return data[key];
            },
            set(value) {
                data[key] = value;
                this.updateDOM();
            },
        });
    }

    const methods = options.methods || {};
    this.methods = {};
    for (let key in methods) {
        this[key] = methods[key];
    }

    this.updateDOM = function () {
        let virtualDOM = this.buildVirtualElement(this.template);
        this.currentInstance = this.reconcile(this.container, virtualDOM, this.currentInstance);
    };

    this.container = document.querySelector(options["el"]);
    this.template = this.container.cloneNode(true);
    this.currentInstance = null;
    this.container.innerHTML = "";
    this.updateDOM();

    if (options.mounted) {
        options.mounted.bind(this)();
    }
}
