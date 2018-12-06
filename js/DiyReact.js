/** @jsx DiyReact.createElement */

const DiyReact = importFromBelow();

const randomLikes = () => Math.ceil(Math.random() * 100);

const stories = [
  {name: "DiyReact介绍", url: "http://google.com", likes: randomLikes()},
  {name: "Rendering DOM elements ", url: "http://google.com", likes: randomLikes()},
  {name: "Element creation and JSX", url: "http://google.com", likes: randomLikes()},
  {name: "Instances and reconciliation", url: "http://google.com", likes: randomLikes()},
  {name: "Components and state", url: "http://google.com", likes: randomLikes()}
];

class App extends DiyReact.Component {
  render() {
    return (
      <div>
        <h1>DiyReact Stories</h1>
        <ul>
          {this.props.stories.map(story => {
            return <Story name={story.name} url={story.url} />;
          })}
        </ul>
      </div>
    );
  }
  
  componentWillMount() {
    console.log('execute componentWillMount');
  }
  
  componentDidMount() {
    console.log('execute componentDidMount');
  }
  
  componentWillUnmount() {
    console.log('execute componentWillUnmount');
  }
}

class Story extends DiyReact.Component {
  constructor(props) {
    super(props);
    this.state = { likes: Math.ceil(Math.random() * 100) };
  }
  like() {
    this.setState({
      likes: this.state.likes + 1
    });
  }
  render() {
    const { name, url } = this.props;
    const { likes } = this.state;
    const likesElement = <span />;
    return (
      <li>
        <button onClick={e => this.like()}>{likes}<b>❤️</b></button>
        <a href={url}>{name}</a>
      </li>
    );
  }
  
  // shouldcomponentUpdate() {
  //   return true;
  // }
  
  componentWillUpdate() {
    console.log('execute componentWillUpdate');
  }
  
  componentDidUpdate() {
    console.log('execute componentDidUpdate');
  }

}

DiyReact.render(<App stories={stories} />, document.getElementById("root"));

/* 🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼🌼 */

function importFromBelow() {
  const TEXT_ELEMENT = 'TEXT_ELEMENT';
  
  function updateDomProperties(dom, prevProps, nextProps) {
    const isEvent = name => name.startsWith("on");
    const isAttribute = name => !isEvent(name) && name != "children";

    // Remove event listeners
    Object.keys(prevProps).filter(isEvent).forEach(name => {
      const eventType = name.toLowerCase().substring(2);
      dom.removeEventListener(eventType, prevProps[name]);
    });

    // Remove attributes
    Object.keys(prevProps).filter(isAttribute).forEach(name => {
      dom[name] = null;
    });

    // Set attributes
    Object.keys(nextProps).filter(isAttribute).forEach(name => {
      dom[name] = nextProps[name];
    });

    // Add event listeners
    Object.keys(nextProps).filter(isEvent).forEach(name => {
      const eventType = name.toLowerCase().substring(2);
      dom.addEventListener(eventType, nextProps[name]);
    });
  }
  
  let rootInstance = null;
  function render(element, parentDom) {
    const prevInstance = rootInstance;
    const nextInstance = reconcile(parentDom, prevInstance, element);
    rootInstance = nextInstance;
  }
  
  function reconcile(parentDom, instance, element) {
    if (instance === null) {
      const newInstance = instantiate(element);
      // componentWillMount
      newInstance.publicInstance
        && newInstance.publicInstance.componentWillMount
        && newInstance.publicInstance.componentWillMount();
      parentDom.appendChild(newInstance.dom);
      // componentDidMount
      newInstance.publicInstance
        && newInstance.publicInstance.componentDidMount
        && newInstance.publicInstance.componentDidMount();
      return newInstance;
    } else if (element === null) {
      // componentWillUnmount
      instance.publicInstance
        && instance.publicInstance.componentWillUnmount
        && instance.publicInstance.componentWillUnmount();
      parentDom.removeChild(instance.dom);
      return null;
    } else if (instance.element.type !== element.type) {
      const newInstance = instantiate(element);
      // componentDidMount
      newInstance.publicInstance
        && newInstance.publicInstance.componentDidMount
        && newInstance.publicInstance.componentDidMount();
      parentDom.replaceChild(newInstance.dom, instance.dom);
      return newInstance;
    } else if (typeof element.type === 'string') {
      updateDomProperties(instance.dom, instance.element.props, element.props);
      instance.childInstances = reconcileChildren(instance, element);
      instance.element = element;
      return instance;
    } else {
      if (instance.publicInstance
          && instance.publicInstance.shouldcomponentUpdate) {
        if (!instance.publicInstance.shouldcomponentUpdate()) {
          return;
        }
      }
      // componentWillUpdate
      instance.publicInstance
        && instance.publicInstance.componentWillUpdate
        && instance.publicInstance.componentWillUpdate();
      instance.publicInstance.props = element.props;
      const newChildElement = instance.publicInstance.render();
      const oldChildInstance = instance.childInstance;
      const newChildInstance = reconcile(parentDom, oldChildInstance, newChildElement);
      // componentDidUpdate
      instance.publicInstance
        && instance.publicInstance.componentDidUpdate
        && instance.publicInstance.componentDidUpdate();
      instance.dom = newChildInstance.dom;
      instance.childInstance = newChildInstance;
      instance.element = element;
      return instance;
    }
  }

  function reconcileChildren(instance, element) {
    const {dom, childInstances} = instance;
    const newChildElements = element.props.children || [];
    const count = Math.max(childInstances.length, newChildElements.length);
    const newChildInstances = [];
    for (let i = 0; i < count; i++) {
      newChildInstances[i] = reconcile(dom, childInstances[i], newChildElements[i]);
    }
    return newChildInstances.filter(instance => instance !== null);
  }
  
  function instantiate(element) {
    const {type, props = {}} = element;
    
    const isDomElement = typeof type === 'string';
    
    if (isDomElement) {
      // 创建dom
      const isTextElement = type === TEXT_ELEMENT;
      const dom = isTextElement ? document.createTextNode('') : document.createElement(type);

      // 设置dom的事件、数据属性
      updateDomProperties(dom, [], element.props);
      const children = props.children || [];
      const childInstances = children.map(instantiate);
      const childDoms = childInstances.map(childInstance => childInstance.dom);
      childDoms.forEach(childDom => dom.appendChild(childDom));
      const instance = {element, dom, childInstances};
      return instance;
    } else {
      const instance = {};
      const publicInstance = createPublicInstance(element, instance);
      const childElement = publicInstance.render();
      const childInstance = instantiate(childElement);
      Object.assign(instance, {dom: childInstance.dom, element, childInstance, publicInstance});
      return instance;
    }
  }
  
  function createTextElement(value) {
    return createElement(TEXT_ELEMENT, {nodeValue: value});
  }
  
  function createElement(type, props, ...children) {
    props = Object.assign({}, props);
    props.children = [].concat(...children)
      .filter(child => child != null && child !== false)
      .map(child => child instanceof Object ? child : createTextElement(child));
    return {type, props};
  }
  
  function createPublicInstance(element, instance) {
    const {type, props} = element;
    const publicInstance = new type(props);
    publicInstance.__internalInstance = instance;
    return publicInstance;
  }
  
  class Component {
    constructor(props) {
      this.props = props;
      this.state = this.state || {};
    }
    
    setState(partialState) {
      this.state = Object.assign({}, this.state, partialState);
      // update instance
      const parentDom = this.__internalInstance.dom.parentNode;
      const element = this.__internalInstance.element;
      reconcile(parentDom, this.__internalInstance, element);
    }
  }
  
  return {
    render,
    createElement,
    Component
  };
}
