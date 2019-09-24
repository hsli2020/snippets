// Example 3
const MyReact = (function() {
  let _val, _deps // hold our state and dependencies in scope
  return {
    render(Component) {
      const Comp = Component()
      Comp.render()
      return Comp
    },
    useEffect(callback, depArray) {
      const hasNoDeps = !depArray
      const hasChangedDeps = _deps ? !depArray.every((el, i) => el === _deps[i]) : true
      if (hasNoDeps || hasChangedDeps) {
        callback()
        _deps = depArray
      }
    },
    useState(initialValue) {
      _val = _val || initialValue
      function setState(newVal) {
        _val = newVal
      }
      return [_val, setState]
    }
  }
})()

// usage
function Counter() {
  const [count, setCount] = MyReact.useState(0)
  MyReact.useEffect(() => {
    console.log('effect', count)
  }, [count])
  return {
    click: () => setCount(count + 1),
    noop: () => setCount(count),
    render: () => console.log('render', { count })
  }
}

let App
App = MyReact.render(Counter)
// effect 0
// render {count: 0}

App.click()
App = MyReact.render(Counter)
// effect 1
// render {count: 1}

App.noop()
App = MyReact.render(Counter)
// // no effect run
// render {count: 1}

App.click()
App = MyReact.render(Counter)
// effect 2
// render {count: 2}
