let obj = {};
let handler = {
  get(target, property) {
   console.log(`${property} 被读取`);
   return property in target ? target[property] : 3;
  },
  set(target, property, value) {
   console.log(`${property} 被设置为 ${value}`);
   target[property] = value;
  }
}

let p = new Proxy(obj, handler);
p.name = 'tom' //name 被设置为 tom
p.age; //age 被读取 3
//console.log(p);
console.log(obj);
