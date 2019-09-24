// Example 0
function useState(initialValue) {
  var _val = initialValue // _val is a local variable created by useState
  function state() {
    // state is an inner function, a closure
    return _val // state() uses _val, declared by parent funciton
  }
  function setState(newVal) {
    // same
    _val = newVal // setting _val without exposing _val
  }
  return [state, setState] // exposing functions for external use
}

var [foo, setFoo] = useState(0) // using array destructuring

console.log(foo()) // logs 0 - the initialValue we gave
setFoo(1) // sets _val inside useState's scope
console.log(foo()) // logs 1 - new initialValue, despite exact same call
