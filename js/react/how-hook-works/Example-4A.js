// Example 4 - Custom Hooks
const MyReact = (function() {
  let hooks = [],
      currentHook = 0 // array of hooks, and an iterator!
  return {
    render(Component) {
      const Comp = Component() // run effects
      Comp.render()
      currentHook = 0 // reset for next render
      return Comp
    },
    useEffect(callback, depArray) {
      const hasNoDeps = !depArray
      const deps = hooks[currentHook] // type: array | undefined
      const hasChangedDeps = deps ? !depArray.every((el, i) => el === deps[i]) : true
      if (hasNoDeps || hasChangedDeps) {
        callback()
        hooks[currentHook] = depArray
      }
      currentHook++ // done with this hook
    },
    useState(initialValue) {
      hooks[currentHook] = hooks[currentHook] || initialValue // type: any
      const setStateHookIndex = currentHook // for setState's closure!
      const setState = newState => (hooks[setStateHookIndex] = newState)
      return [hooks[currentHook++], setState]
    }
  }
})()

// Example 4, revisited
function Component() {
  const [text, setText] = useSplitURL('www.netlify.com')
  return {
    type: txt => setText(txt),
    render: () => console.log({ text })
  }
}

function useSplitURL(str) {
  const [text, setText] = MyReact.useState(str)
  const masked = text.split('.')
  return [masked, setText]
}

let App
App = MyReact.render(Component)
// { text: [ 'www', 'netlify', 'com' ] }

App.type('www.reactjs.org')
App = MyReact.render(Component)
// { text: [ 'www', 'reactjs', 'org' ] }}
