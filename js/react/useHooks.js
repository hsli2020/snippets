useHooks

// Hooks are a new addition in React 16.8 that lets you use state and other React
// features without writing a class. This website provides easy to understand code
// examples to help you learn how hooks work and hopefully inspire you to take
// advantage of them in your next project. Be sure to check out the official docs.
// You can also submit post ideas in our Github repo.

// useWhyDidYouUpdate

// This hook makes it easy to see which prop changes are causing a component to
// re-render. If a function is particularly expensive to run and you know it
// renders the same results given the same props you can use the React.memo
// higher order component, as we've done with the Counter component in the
// below example. In this case if you're still seeing re-renders that seem
// uneccessary you can drop in the useWhyDidYouUpdate hook and check your
// console to see which props changed between renders and view their
// previous/current values. Pretty nifty huh?

import { useState, useEffect, useRef } from 'react';

// Let's pretend this <Counter> component is expensive to re-render so ...
// ... we wrap with React.memo, but we're still seeing performance issues :/
// So we add useWhyDidYouUpdate and check our console to see what's going on.

const Counter = React.memo(props => {
  useWhyDidYouUpdate('Counter', props);
  return <div style={props.style}>{props.count}</div>;
});

function App() {

  const [count, setCount] = useState(0);
  const [userId, setUserId] = useState(0);

  // Our console output tells use that the style prop for <Counter> ...
  // ... changes on every render, even when we only change userId state by ...
  // ... clicking the "switch user" button. Oh of course! That's because the
  // ... counterStyle object is being re-created on every render.
  // Thanks to our hook we figured this out and realized we should probably ...
  // ... move this object outside of the component body.

  const counterStyle = {
    fontSize: '3rem',
    color: 'red'
  };

  return (
    <div>
      <div className="counter">
        <Counter count={count} style={counterStyle} />
        <button onClick={() => setCount(count + 1)}>Increment</button>
      </div>
      <div className="user">
        <img src={`http://i.pravatar.cc/80?img=${userId}`} />
        <button onClick={() => setUserId(userId + 1)}>Switch User</button>
      </div>
    </div>
  );
}

// Hook
function useWhyDidYouUpdate(name, props) {

  // Get a mutable ref object where we can store props ...
  // ... for comparison next time this hook runs.

  const previousProps = useRef();

  useEffect(() => {

    if (previousProps.current) {

      // Get all keys from previous and current props
      const allKeys = Object.keys({ ...previousProps.current, ...props });

      // Use this object to keep track of changed props
      const changesObj = {};

      // Iterate through keys
      allKeys.forEach(key => {

        // If previous is different from current
        if (previousProps.current[key] !== props[key]) {

          // Add to changesObj
          changesObj[key] = {
            from: previousProps.current[key],
            to: props[key]
          };
        }
      });

      // If changesObj not empty then output to console
      if (Object.keys(changesObj).length) {
        console.log('[why-did-you-update]', name, changesObj);
      }
    }

    // Finally update previousProps with current props for next hook call
    previousProps.current = props;
  });
}

// useDarkMode

// This hook handles all the stateful logic required to add a ☾ dark mode
// toggle to your website. It utilizes localStorage to remember the user's
// chosen mode, defaults to their browser or OS level setting using the
// prefers-color-scheme media query and manages the setting of a .dark-mode
// className on body to apply your styles.

// This post also helps illustrate the power of hook composition. The syncing
// of state to localStorage is handled by our useLocalStorage hook. Detecting
// the user's dark mode preference is handled by our useMedia hook. Both of
// these hooks were created for other use-cases, but here we've composed them
// to create a super useful hook in relatively few lines of code. It's almost
// as if hooks bring the compositional power of React components to stateful
// logic! 

// Usage

function App() {
  const [darkMode, setDarkMode] = useDarkMode();

  return (
    <div>
      <div className="navbar">
        <Toggle darkMode={darkMode} setDarkMode={setDarkMode} />
      </div>
      <Content />
    </div>
  );
}

// Hook
function useDarkMode() {

  // Use our useLocalStorage hook to persist state through a page refresh.
  // Read the recipe for this hook to learn more: usehooks.com/useLocalStorage

  const [enabledState, setEnabledState] = useLocalStorage('dark-mode-enabled');

  // See if user has set a browser or OS preference for dark mode.
  // The usePrefersDarkMode hook composes a useMedia hook (see code below).
  const prefersDarkMode = usePrefersDarkMode();

  // If enabledState is defined use it, otherwise fallback to prefersDarkMode.
  // This allows user to override OS level setting on our website.
  const enabled =
    typeof enabledState !== 'undefined' ? enabledState : prefersDarkMode;

  // Fire off effect that add/removes dark mode class
  useEffect(
    () => {
      const className = 'dark-mode';
      const element = window.document.body;

      if (enabled) {
        element.classList.add(className);
      } else {
        element.classList.remove(className);
      }
    },

    [enabled] // Only re-call effect when value changes
  );

  // Return enabled state and setter
  return [enabled, setEnabledState];
}

// Compose our useMedia hook to detect dark mode preference.
// The API for useMedia looks a bit weird, but that's because ...
// ... it was designed to support multiple media queries and return values.
// Thanks to hook composition we can hide away that extra complexity!
// Read the recipe for useMedia to learn more: usehooks.com/useMedia
function usePrefersDarkMode() {
  return useMedia(['(prefers-color-scheme: dark)'], [true], false);
}

// useMedia

// This hook makes it super easy to utilize media queries in your component
// logic. In our example below we render a different number of columns
// depending on which media query matches the current screen width, and then
// distribute images amongst the columns in a way that limits column height
// difference (we don't want one column way longer than the rest).

// You could create a hook that directly measures screen width instead of using
// media queries, but this method is nice because it makes it easy to share
// media queries between JS and your stylesheet. See it in action in the
// CodeSandbox Demo.

import { useState, useEffect } from 'react';

function App() {

  const columnCount = useMedia(
    // Media queries
    ['(min-width: 1500px)', '(min-width: 1000px)', '(min-width: 600px)'],

    // Column counts (relates to above media queries by array index)
    [5, 4, 3],

    // Default column count
    2
  );

  // Create array of column heights (start at 0)
  let columnHeights = new Array(columnCount).fill(0);

  // Create array of arrays that will hold each column's items
  let columns = new Array(columnCount).fill().map(() => []);

  data.forEach(item => {
    // Get index of shortest column
    const shortColumnIndex = columnHeights.indexOf(Math.min(...columnHeights));

    // Add item
    columns[shortColumnIndex].push(item);

    // Update height
    columnHeights[shortColumnIndex] += item.height;
  });

  // Render columns and items
  return (
    <div className="App">
      <div className="columns is-mobile">
        {columns.map(column => (
          <div className="column">
            {column.map(item => (
              <div
                className="image-container"
                style={{
                  // Size image container to aspect ratio of image
                  paddingTop: (item.height / item.width) * 100 + '%'
                }}
              >
                <img src={item.image} alt="" />
              </div>
            ))}
          </div>
        ))}
      </div>
    </div>
  );
}

// Hook
function useMedia(queries, values, defaultValue) {

  // Array containing a media query list for each query
  const mediaQueryLists = queries.map(q => window.matchMedia(q));

  // Function that gets value based on matching media query
  const getValue = () => {

    // Get index of first media query that matches
    const index = mediaQueryLists.findIndex(mql => mql.matches);

    // Return related value or defaultValue if none
    return typeof values[index] !== 'undefined' ? values[index] : defaultValue;
  };

  // State and setter for matched value
  const [value, setValue] = useState(getValue);

  useEffect(
    () => {
      // Event listener callback
      // Note: By defining getValue outside of useEffect we ensure that it has ...
      // ... current values of hook args (as this hook callback is created once on mount).
      const handler = () => setValue(getValue);

      // Set a listener for each media query with above handler as callback.
      mediaQueryLists.forEach(mql => mql.addListener(handler));

      // Remove listeners on cleanup
      return () => mediaQueryLists.forEach(mql => mql.removeListener(handler));
    },
    [] // Empty array ensures effect is only run on mount and unmount
  );

  return value;
}

// useLockBodyScroll

// Sometimes you want to prevent your users from being able to scroll the body
// of your page while a particular component is absolutely positioned over your
// page (think modal or full-screen mobile menu). It can be confusing to see
// the background content scroll underneath a modal, especially if you intended
// to scroll an area within the modal. Well, this hook solves that! Simply call
// the useLockBodyScroll hook in any component and body scrolling will be
// locked until that component unmounts. See it in action in the CodeSandbox
// Demo.

import { useState, useLayoutEffect } from 'react';

// Usage

function App(){

  // State for our modal
  const [modalOpen, setModalOpen] = useState(false);

  return (
    <div>
      <button onClick={() => setModalOpen(true)}>Show Modal</button>
      <Content />
      {modalOpen && (
        <Modal
          title="Try scrolling"
          content="I bet you you can't! Muahahaha 😈"
          onClose={() => setModalOpen(false)}
        />
      )}
    </div>
  );
}

function Modal({ title, content, onClose }){

  // Call hook to lock body scroll
  useLockBodyScroll();

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal">
        <h2>{title}</h2>
        <p>{content}</p>
      </div>
    </div>
  );
}

// Hook
function useLockBodyScroll() {

  useLayoutEffect(() => {

   // Get original body overflow
   const originalStyle = window.getComputedStyle(document.body).overflow;  

   // Prevent scrolling on mount
   document.body.style.overflow = 'hidden';

   // Re-enable scrolling when component unmounts
   return () => document.body.style.overflow = originalStyle;
   }, []); // Empty array ensures effect is only run on mount and unmount
}

// useTheme

// This hook makes it easy to dynamically change the appearance of your app
// using CSS variables. You simply pass in an object containing key/value pairs
// of the CSS variables you'd like to update and the hook updates each variable
// in the document's root element. This is useful in situations where you can't
// define styles inline (no psudeoclass support) and there are too many style
// permutations to include each theme in your stylesheet (such as a web app
// that lets users customize the look of their profile). It's worth noting that
// many css-in-js libraries support dynamic styles out of the box, but it's
// interesting to experiment with how this can be done with just CSS variables
// and a React Hook. The example below is intentionally very simple, but you
// could imagine the theme object being stored in state or fetched from an API.
// Be sure to check out the CodeSandbox demo for a more interesting example and
// to see the accompanying stylesheet.

import { useLayoutEffect } from 'react';

import './styles.scss'; // -> https://codesandbox.io/s/15mko9187

// Usage

const theme = {
  'button-padding': '16px',
  'button-font-size': '14px',
  'button-border-radius': '4px',
  'button-border': 'none',
  'button-color': '#FFF',
  'button-background': '#6772e5',
  'button-hover-border': 'none',
  'button-hover-color': '#FFF'
};

function App() {

  useTheme(theme);

  return (
    <div>
      <button className="button">Button</button>
    </div>
  );
}

// Hook

function useTheme(theme) {

  useLayoutEffect(

    () => {
      // Iterate through each value in theme object
      for (const key in theme) {
        // Update css variables in document's root element
        document.documentElement.style.setProperty(`--${key}`, theme[key]);
      }
    },

    [theme] // Only call again if theme object reference changes
  );
}

// useSpring

// This hook is part of the react-spring animation library which allows for
// highly performant physics-based animations. I try to avoid including
// dependencies in these recipes, but once in awhile I'm going to make an
// exception for hooks that expose the functionality of really useful
// libraries. One nice thing about react-spring is that it allows you to
// completely skip the React render cycle when applying animations, often
// giving a pretty substantial performance boost. In our recipe below we render
// a row of cards and apply a springy animation effect related to the mouse
// position over any given card. To make this work we call the useSpring hook
// with an array of values we want to animate, render an animated.div component
// (exported by react-spring), get the mouse position over a card with the
// onMouseMove event, then call setAnimatedProps (function returned by the
// hook) to update that set of values based on the mouse position. Read through
// the comments in the recipe below for more details or jump right over to the
// CodeSandbox demo. I liked this effect so much I ended up using it on my
// startup's landing page 😎

import { useState, useRef } from 'react';
import { useSpring, animated } from 'react-spring';

// Displays a row of cards
// Usage of hook is within <Card> component below

function App() {
  return (
    <div className="container">
      <div className="row">
        {cards.map((card, i) => (
          <div className="column">
            <Card>
              <div className="card-title">{card.title}</div>
              <div className="card-body">{card.description}</div>
              <img className="card-image" src={card.image} />
            </Card>
          </div>
        ))}
      </div>
    </div>
  );
}

function Card({ children }) {

  // We add this ref to card element and use in onMouseMove event ...
  // ... to get element's offset and dimensions.
  const ref = useRef();

  // Keep track of whether card is hovered so we can increment ...
  // ... zIndex to ensure it shows up above other cards when animation causes overlap.
  const [isHovered, setHovered] = useState(false);

  // The useSpring hook
  const [animatedProps, setAnimatedProps] = useSpring({
    // Array containing [rotateX, rotateY, and scale] values.
    // We store under a single key (xys) instead of separate keys ...
    // ... so that we can use animatedProps.xys.interpolate() to ...
    // ... easily generate the css transform value below.
    xys: [0, 0, 1],

    // Setup physics
    config: { mass: 10, tension: 400, friction: 40, precision: 0.00001 }
  });

  return (
    <animated.div
      ref={ref}
      className="card"
      onMouseEnter={() => setHovered(true)}
      onMouseMove={({ clientX, clientY }) => {
        // Get mouse x position within card
        const x =
          clientX -
          (ref.current.offsetLeft -
            (window.scrollX || window.pageXOffset || document.body.scrollLeft));

        // Get mouse y position within card
        const y =
          clientY -
          (ref.current.offsetTop -
            (window.scrollY || window.pageYOffset || document.body.scrollTop));

        // Set animated values based on mouse position and card dimensions
        const dampen = 50; // Lower the number the less rotation
        const xys = [
          -(y - ref.current.clientHeight / 2) / dampen, // rotateX
          (x - ref.current.clientWidth / 2) / dampen, // rotateY
          1.07 // Scale
        ];

        // Update values to animate to
        setAnimatedProps({ xys: xys });
      }}

      onMouseLeave={() => {
        setHovered(false);
        // Set xys back to original
        setAnimatedProps({ xys: [0, 0, 1] });
      }}

      style={{
        // If hovered we want it to overlap other cards when it scales up
        zIndex: isHovered ? 2 : 1,

        // Interpolate function to handle css changes
        transform: animatedProps.xys.interpolate(
          (x, y, s) =>
            `perspective(600px) rotateX(${x}deg) rotateY(${y}deg) scale(${s})`
        )
      }}
    >
      {children}
    </animated.div>
  );
}

// useHistory

// This hook makes it really easy to add undo/redo functionality to your app.
// Our recipe is a simple drawing app. It generates a grid of blocks, allows
// you to click any block to toggle its color, and uses the useHistory hook so
// we can undo, redo, or clear all changes to the canvas. Check out our
// CodeSandbox demo. Within our hook we're using useReducer to store state
// instead of useState, which should look familiar to anyone that's used redux
// (read more about useReducer in the official docs). The hook code was copied,
// with minor changes, from the excellent use-undo library, so if you'd like to
// pull this into your project you can also use that library via npm.

import { useReducer, useCallback } from 'react';

// Usage

function App() {
  const { state, set, undo, redo, clear, canUndo, canRedo } = useHistory({});

  return (
    <div className="container">
      <div className="controls">
        <div className="title">👩‍🎨 Click squares to draw</div>
        <button onClick={undo} disabled={!canUndo}>
          Undo
        </button>
        <button onClick={redo} disabled={!canRedo}>
          Redo
        </button>
        <button onClick={clear}>Clear</button>
      </div>

      <div className="grid">
        {((blocks, i, len) => {
          // Generate a grid of blocks
          while (++i <= len) {
            const index = i;
            blocks.push(
              <div
                // Give block "active" class if true in state object
                className={'block' + (state[index] ? ' active' : '')}
                // Toggle boolean value of clicked block and merge into current state
                onClick={() => set({ ...state, [index]: !state[index] })}
                key={i}
              />
            );
          }
          return blocks;
        })([], 0, 625)}
      </div>
    </div>
  );
}

// Initial state that we pass into useReducer
const initialState = {
  // Array of previous state values updated each time we push a new state
  past: [],

  // Current state value
  present: null,

  // Will contain "future" state values if we undo (so we can redo)
  future: []
};

// Our reducer function to handle state changes based on action
const reducer = (state, action) => {

  const { past, present, future } = state;

  switch (action.type) {
    case 'UNDO':
      const previous = past[past.length - 1];
      const newPast = past.slice(0, past.length - 1);

      return {
        past: newPast,
        present: previous,
        future: [present, ...future]
      };

    case 'REDO':
      const next = future[0];
      const newFuture = future.slice(1);

      return {
        past: [...past, present],
        present: next,
        future: newFuture
      };

    case 'SET':
      const { newPresent } = action;

      if (newPresent === present) {
        return state;
      }

      return {
        past: [...past, present],
        present: newPresent,
        future: []
      };

    case 'CLEAR':
      const { initialPresent } = action;
      return {
        ...initialState,
        present: initialPresent
      };
  }
};

// Hook
const useHistory = initialPresent => {
  const [state, dispatch] = useReducer(reducer, {
    ...initialState,
    present: initialPresent
  });

  const canUndo = state.past.length !== 0;
  const canRedo = state.future.length !== 0;

  // Setup our callback functions
  // We memoize with useCallback to prevent unecessary re-renders

  const undo = useCallback(
    () => {
      if (canUndo) {
        dispatch({ type: 'UNDO' });
      }
    },
    [canUndo, dispatch]
  );

  const redo = useCallback(
    () => {
      if (canRedo) {
        dispatch({ type: 'REDO' });
      }
    },
    [canRedo, dispatch]
  );

  const set = useCallback(newPresent => dispatch({ type: 'SET', newPresent }), [
    dispatch
  ]);

  const clear = useCallback(() => dispatch({ type: 'CLEAR', initialPresent }), [
    dispatch
  ]);

  // If needed we could also return past and future state
  return { state: state.present, set, undo, redo, clear, canUndo, canRedo };
};

// useScript

// This hook makes it super easy to dynamically load an external script and
// know when its loaded. This is useful when you need to interact with a 3rd
// party libary (Stripe, Google Analytics, etc) and you'd prefer to load the
// script when needed rather then include it in the document head for every
// page request. In the example below we wait until the script has loaded
// successfully before calling a function declared in the script. If you're
// interested in seeing how this would look if implemented as a Higher Order
// Component then check out the source of react-script-loader-hoc. I personally
// find it much more readable as a hook. Another advantage is because it's so
// easy to call the same hook multiple times to load multiple different
// scripts, unlike the HOC implementation, we can skip adding support for
// passing in multiple src strings.

import { useState, useEffect } from 'react';

// Usage
function App() {
  const [loaded, error] = useScript(
    'https://pm28k14qlj.codesandbox.io/test-external-script.js'
  );

  return (
    <div>
      <div>
        Script loaded: <b>{loaded.toString()}</b>
      </div>
      {loaded && !error && (
        <div>
          Script function call response: <b>{TEST_SCRIPT.start()}</b>
        </div>
      )}
    </div>
  );
}

// Hook
let cachedScripts = [];

function useScript(src) {
  // Keeping track of script loaded and error state
  const [state, setState] = useState({
    loaded: false,
    error: false
  });

  useEffect(
    () => {
      // If cachedScripts array already includes src that means another instance ...
      // ... of this hook already loaded this script, so no need to load again.

      if (cachedScripts.includes(src)) {
        setState({
          loaded: true,
          error: false
        });
      } else {
        cachedScripts.push(src);

        // Create script
        let script = document.createElement('script');
        script.src = src;
        script.async = true;

        // Script event listener callbacks for load and error
        const onScriptLoad = () => {
          setState({
            loaded: true,
            error: false
          });
        };

        const onScriptError = () => {
          // Remove from cachedScripts we can try loading again
          const index = cachedScripts.indexOf(src);

          if (index >= 0) cachedScripts.splice(index, 1);

          script.remove();

          setState({
            loaded: true,
            error: true
          });
        };

        script.addEventListener('load', onScriptLoad);
        script.addEventListener('error', onScriptError);

        // Add script to document body
        document.body.appendChild(script);

        // Remove event listeners on cleanup
        return () => {
          script.removeEventListener('load', onScriptLoad);
          script.removeEventListener('error', onScriptError);
        };
      }
    },
    [src] // Only re-run effect if script src changes
  );

  return [state.loaded, state.error];
}

// useKeyPress

// This hook makes it easy to detect when the user is pressing a specific key
// on their keyboard. The recipe is fairly simple, as I want to show how little
// code is required, but I challenge any readers to create a more advanced
// version of this hook. Detecting when multiple keys are held down at the same
// time would be a nice addition. Bonus points: also require they be held down
// in a specified order. Feel free to share anything you've created in this
// recipe's gist.

import { useState, useEffect } from 'react';

// Usage

function App() {

  // Call our hook for each key that we'd like to monitor

  const happyPress = useKeyPress('h');
  const sadPress = useKeyPress('s');
  const robotPress = useKeyPress('r');
  const foxPress = useKeyPress('f');

  return (
    <div>
      <div>h, s, r, f</div>
      <div>
        {happyPress && '😊'}
        {sadPress && '😢'}
        {robotPress && '🤖'}
        {foxPress && '🦊'}
      </div>
    </div>
  );
}

// Hook
function useKeyPress(targetKey) {

  // State for keeping track of whether key is pressed
  const [keyPressed, setKeyPressed] = useState(false);

  // If pressed key is our target key then set to true
  function downHandler({ key }) {
    if (key === targetKey) {
      setKeyPressed(true);
    }
  }

  // If released key is our target key then set to false
  const upHandler = ({ key }) => {
    if (key === targetKey) {
      setKeyPressed(false);
    }
  };

  // Add event listeners
  useEffect(() => {
    window.addEventListener('keydown', downHandler);
    window.addEventListener('keyup', upHandler);

    // Remove event listeners on cleanup
    return () => {
      window.removeEventListener('keydown', downHandler);
      window.removeEventListener('keyup', upHandler);
    };
  }, []); // Empty array ensures that effect is only run on mount and unmount

  return keyPressed;
}

// useMemo

// React has a built-in hook called useMemo that allows you to memoize
// expensive functions so that you can avoid calling them on every render. You
// simple pass in a function and an array of inputs and useMemo will only
// recompute the memoized value when one of the inputs has changed. In our
// example below we have an expensive function called computeLetterCount (for
// demo purposes we make it slow by including a large and completely
// unnecessary loop). When the current selected word changes you'll notice a
// delay as it has to recall computeLetterCount on the new word. We also have a
// separate counter that gets incremented everytime the increment button is
// clicked. When that counter is incremented you'll notice that there is zero
// lag between renders. This is because computeLetterCount is not called again.
// The input word hasn't changed and thus the cached value is returned. You'll
// probably want to check out the CodeSandbox demo so you can see for yourself.

import { useState, useMemo } from 'react';

// Usage

function App() {

  // State for our counter
  const [count, setCount] = useState(0);

  // State to keep track of current word in array we want to show
  const [wordIndex, setWordIndex] = useState(0);

  // Words we can flip through and view letter count
  const words = ['hey', 'this', 'is', 'cool'];
  const word = words[wordIndex];

  // Returns number of letters in a word
  // We make it slow by including a large and completely unnecessary loop
  const computeLetterCount = word => {
    let i = 0;
    while (i < 1000000000) i++;
    return word.length;
  };

  // Memoize computeLetterCount so it uses cached return value if input array ...
  // ... values are the same as last time the function was run.
  const letterCount = useMemo(() => computeLetterCount(word), [word]);

  // This would result in lag when incrementing the counter because ...
  // ... we'd have to wait for expensive function when re-rendering.
  //const letterCount = computeLetterCount(word);

  return (
    <div style={{ padding: '15px' }}>
      <h2>Compute number of letters (slow 🐌)</h2>
      <p>"{word}" has {letterCount} letters</p>
      <button
        onClick={() => {
          const next = wordIndex + 1 === words.length ? 0 : wordIndex + 1;
          setWordIndex(next);
        }}
      >
        Next word
      </button>

      <h2>Increment a counter (fast ⚡️)</h2>
      <p>Counter: {count}</p>
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </div>
  );
}

// useDebounce

// This hook allows you to debounce any fast changing value. The debounced
// value will only reflect the latest value when the useDebounce hook has not
// been called for the specified time period. When used in conjuction with
// useEffect, as we do in the recipe below, you can easily ensure that
// expensive operations like API calls are not executed too frequently. The
// example below allows you to search the Marvel Comic API and uses useDebounce
// to prevent API calls from being fired on every keystroke. Be sure to theck
// out the CodeSandbox demo for this one. Hook code and inspiration from
// github.com/xnimorz/use-debounce.

import { useState, useEffect, useRef } from 'react';

// Usage

function App() {

  // State and setters for ...
  // Search term
  const [searchTerm, setSearchTerm] = useState('');

  // API search results
  const [results, setResults] = useState([]);

  // Searching status (whether there is pending API request)
  const [isSearching, setIsSearching] = useState(false);

  // Debounce search term so that it only gives us latest value ...
  // ... if searchTerm has not been updated within last 500ms.
  // The goal is to only have the API call fire when user stops typing ...
  // ... so that we aren't hitting our API rapidly.
  const debouncedSearchTerm = useDebounce(searchTerm, 500);

  // Effect for API call 
  useEffect(
    () => {
      if (debouncedSearchTerm) {
        setIsSearching(true);
        searchCharacters(debouncedSearchTerm).then(results => {
          setIsSearching(false);
          setResults(results);
        });
      } else {
        setResults([]);
      }
    },
    [debouncedSearchTerm] // Only call effect if debounced search term changes
  );

  return (
    <div>
      <input
        placeholder="Search Marvel Comics"
        onChange={e => setSearchTerm(e.target.value)}
      />

      {isSearching && <div>Searching ...</div>}

      {results.map(result => (
        <div key={result.id}>
          <h4>{result.title}</h4>
          <img
            src={`${result.thumbnail.path}/portrait_incredible.${
              result.thumbnail.extension
            }`}
          />
        </div>
      ))}
    </div>
  );
}

// API search function

function searchCharacters(search) {

  const apiKey = 'f9dfb1e8d466d36c27850bedd2047687';

  return fetch(
    `https://gateway.marvel.com/v1/public/comics?apikey=${apiKey}&titleStartsWith=${search}`,
    {
      method: 'GET'
    }
  )  
    .then(r => r.json())
    .then(r => r.data.results)
    .catch(error => {
      console.error(error);
      return [];
    });
}

// Hook

function useDebounce(value, delay) {

  // State and setters for debounced value
  const [debouncedValue, setDebouncedValue] = useState(value);

  useEffect(
    () => {
      // Update debounced value after delay
      const handler = setTimeout(() => {
        setDebouncedValue(value);
      }, delay);

      // Cancel the timeout if value changes (also on delay change or unmount)
      // This is how we prevent debounced value from updating if value is changed ...
      // .. within the delay period. Timeout gets cleared and restarted.
      return () => {
        clearTimeout(handler);
      };
    },
    [value, delay] // Only re-call effect if value or delay changes
  );

  return debouncedValue;
}

// useOnScreen

// This hook allows you to easily detect when an element is visible on the
// screen as well as specify how much of the element should be visible before
// being considered on screen. Perfect for lazy loading images or triggering
// animations when the user has scrolled down to a particular section.

import { useState, useEffect, useRef } from 'react';

// Usage

function App() {

  // Ref for the element that we want to detect whether on screen
  const ref = useRef();

  // Call the hook passing in ref and root margin
  // In this case it would only be considered onScreen if more ...
  // ... than 300px of element is visible.
  const onScreen = useOnScreen(ref, '-300px');

  return (
    <div>
      <div style={{ height: '100vh' }}>
        <h1>Scroll down to next section 👇</h1>
      </div>
      <div
        ref={ref}
        style={{
          height: '100vh',
          backgroundColor: onScreen ? '#23cebd' : '#efefef'
        }}
      >
        {onScreen ? (
          <div>
            <h1>Hey I'm on the screen</h1>
            <img src="https://i.giphy.com/media/ASd0Ukj0y3qMM/giphy.gif" />
          </div>
        ) : (
          <h1>Scroll down 300px from the top of this section 👇</h1>
        )}
      </div>
    </div>
  );
}

// Hook

function useOnScreen(ref, rootMargin = '0px') {

  // State and setter for storing whether element is visible
  const [isIntersecting, setIntersecting] = useState(false);

  useEffect(() => {

    const observer = new IntersectionObserver(
      ([entry]) => {
        // Update our state when observer callback fires
        setIntersecting(entry.isIntersecting);
      },
      {
        rootMargin
      }
    );

    if (ref.current) {
      observer.observe(ref.current);
    }

    return () => {
      observer.unobserve(ref.current);
    };
  }, []); // Empty array ensures that effect is only run on mount and unmount

  return isIntersecting;
}

// usePrevious

// One question that comes up a lot is "When using hooks how do I get the
// previous value of props or state?". With React class components you have the
// componentDidUpdate method which receives previous props and state as
// arguments or you can update an instance variable (this.previous = value) and
// reference it later to get the previous value. So how can we do this inside a
// functional component that doesn't have lifecycle methods or an instance to
// store values on? Hooks to the rescue! We can create a custom hook that uses
// the useRef hook internally for storing the previous value. See the recipe
// below with inline comments. You can also find this example in the official
// React Hooks FAQ.

import { useState, useEffect, useRef } from 'react';

// Usage

function App() {

  // State value and setter for our example
  const [count, setCount] = useState(0);

  // Get the previous value (was passed into hook on last render)
  const prevCount = usePrevious(count);

  // Display both current and previous count value
  return (
    <div>
      <h1>Now: {count}, before: {prevCount}</h1>
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </div>
   );
}

// Hook

function usePrevious(value) {

  // The ref object is a generic container whose current property is mutable ...
  // ... and can hold any value, similar to an instance property on a class
  const ref = useRef();

  // Store current value in ref
  useEffect(() => {
    ref.current = value;
  }, [value]); // Only re-run if value changes

  // Return previous value (happens before update in useEffect above)
  return ref.current;
}

// useOnClickOutside

// This hook allows you to detect clicks outside of a specified element. In the
// example below we use it to close a modal when any element outside of the
// modal is clicked. By abstracting this logic out into a hook we can easily
// use it across all of our components that need this kind of functionality
// (dropdown menus, tooltips, etc).

import { useState, useEffect, useRef } from 'react';

// Usage

function App() {

  // Create a ref that we add to the element for which we want to detect outside clicks
  const ref = useRef();

  // State for our modal
  const [isModalOpen, setModalOpen] = useState(false);

  // Call hook passing in the ref and a function to call on outside click
  useOnClickOutside(ref, () => setModalOpen(false));

  return (
    <div>
      {isModalOpen ? (
        <div ref={ref}>
          👋 Hey, I'm a modal. Click anywhere outside of me to close.
        </div>
      ) : (
        <button onClick={() => setModalOpen(true)}>Open Modal</button>
      )}
    </div>
  );
}

// Hook

function useOnClickOutside(ref, handler) {

  useEffect(
    () => {
      const listener = event => {
        // Do nothing if clicking ref's element or descendent elements
        if (!ref.current || ref.current.contains(event.target)) {
          return;
        }

        handler(event);
      };

      document.addEventListener('mousedown', listener);
      document.addEventListener('touchstart', listener);

      return () => {
        document.removeEventListener('mousedown', listener);
        document.removeEventListener('touchstart', listener);
      };
    },

    // Add ref and handler to effect dependencies
    // It's worth noting that because passed in handler is a new ...
    // ... function on every render that will cause this effect ...
    // ... callback/cleanup to run every render. It's not a big deal ...
    // ... but to optimize you can wrap handler in useCallback before ...
    // ... passing it into this hook.
    [ref, handler]
  );
}

// useAnimation

// This hook allows you to smoothly animate any value using an easing function
// (linear, elastic, etc). In the example we call the useAnimation hook three
// times to animated three balls on to the screen at different intervals.
// Additionally we show how easy it is to compose hooks. Our useAnimation hook
// doesn't actual make use of useState or useEffect itself, but instead serves
// as a wrapper around the useAnimationTimer hook. Having the timer logic
// abstracted out into its own hook gives us better code readability and the
// ability to use timer logic in other contexts. Be sure to check out the
// CodeSandbox Demo for this one.

import { useState, useEffect } from 'react';

// Usage

function App() {

  // Call hook multiple times to get animated values with different start delays
  const animation1 = useAnimation('elastic', 600, 0);
  const animation2 = useAnimation('elastic', 600, 150);
  const animation3 = useAnimation('elastic', 600, 300);

  return (
    <div style={{ display: 'flex', justifyContent: 'center' }}>
      <Ball
        innerStyle={{
          marginTop: animation1 * 200 - 100
        }}
      />

      <Ball
        innerStyle={{
          marginTop: animation2 * 200 - 100
        }}
      />

      <Ball
        innerStyle={{
          marginTop: animation3 * 200 - 100
        }}
      />
    </div>
  );
}

const Ball = ({ innerStyle }) => (
  <div
    style={{
      width: 100,
      height: 100,
      marginRight: '40px',
      borderRadius: '50px',
      backgroundColor: '#4dd5fa',
      ...innerStyle
    }}
  />
);

// Hook 

function useAnimation(
  easingName = 'linear',
  duration = 500,
  delay = 0
) {
  // The useAnimationTimer hook calls useState every animation frame ...
  // ... giving us elapsed time and causing a rerender as frequently ...
  // ... as possible for a smooth animation.
  const elapsed = useAnimationTimer(duration, delay);

  // Amount of specified duration elapsed on a scale from 0 - 1
  const n = Math.min(1, elapsed / duration);

  // Return altered value based on our specified easing function
  return easing[easingName](n);
}

// Some easing functions copied from:
// https://github.com/streamich/ts-easing/blob/master/src/index.ts
// Hardcode here or pull in a dependency

const easing = {
  linear: n => n,
  elastic: n =>
    n * (33 * n * n * n * n - 106 * n * n * n + 126 * n * n - 67 * n + 15),
  inExpo: n => Math.pow(2, 10 * (n - 1))
};

function useAnimationTimer(duration = 1000, delay = 0) {

  const [elapsed, setTime] = useState(0);

  useEffect(
    () => {
      let animationFrame, timerStop, start;

      // Function to be executed on each animation frame
      function onFrame() {
        setTime(Date.now() - start);
        loop();
      }

      // Call onFrame() on next animation frame
      function loop() {
        animationFrame = requestAnimationFrame(onFrame);
      }

      function onStart() {
        // Set a timeout to stop things when duration time elapses
        timerStop = setTimeout(() => {
          cancelAnimationFrame(animationFrame);
          setTime(Date.now() - start);
        }, duration);

        // Start the loop
        start = Date.now();
        loop();
      }

      // Start after specified delay (defaults to 0)
      const timerDelay = setTimeout(onStart, delay);

      // Clean things up
      return () => {
        clearTimeout(timerStop);
        clearTimeout(timerDelay);
        cancelAnimationFrame(animationFrame);
      };
    },
    [duration, delay] // Only re-run effect if duration or delay changes
  );

  return elapsed;
}

// useWindowSize

// A really common need is to get the current size of the browser window. This
// hook returns an object containing the window's width and height. If executed
// server-side (no window object) the value of width and height will be undefined.

import { useState, useEffect } from 'react';

// Usage

function App() {

  const size = useWindowSize();

  return (
    <div>
      {size.width}px / {size.height}px
    </div>
  );
}

// Hook

function useWindowSize() {

  const isClient = typeof window === 'object';

  function getSize() {
    return {
      width: isClient ? window.innerWidth : undefined,
      height: isClient ? window.innerHeight : undefined
    };
  }

  const [windowSize, setWindowSize] = useState(getSize);

  useEffect(() => {
    if (!isClient) {
      return false;
    }

    function handleResize() {
      setWindowSize(getSize());
    }

    window.addEventListener('resize', handleResize);

    return () => window.removeEventListener('resize', handleResize);

  }, []); // Empty array ensures that effect is only run on mount and unmount

  return windowSize;
}

// useHover

// Detect whether the mouse is hovering an element. The hook returns a ref and
// a boolean value indicating whether the element with that ref is currently
// being hovered. So just add the returned ref to any element whose hover state
// you want to monitor.

import { useRef, useState, useEffect } from 'react';

// Usage

function App() {

  const [hoverRef, isHovered] = useHover();

  return (
    <div ref={hoverRef}>
      {isHovered ? '😁' : '☹️'}
    </div>
  );
}

// Hook

function useHover() {

  const [value, setValue] = useState(false);
  const ref = useRef(null);
  const handleMouseOver = () => setValue(true);
  const handleMouseOut = () => setValue(false);

  useEffect(
    () => {
      const node = ref.current;

      if (node) {
        node.addEventListener('mouseover', handleMouseOver);
        node.addEventListener('mouseout', handleMouseOut);

        return () => {
          node.removeEventListener('mouseover', handleMouseOver);
          node.removeEventListener('mouseout', handleMouseOut);
        };
      }
    },
    [ref.current] // Recall only if ref changes
  );

  return [ref, value];
}

// useLocalStorage

// Sync state to local storage so that it persists through a page refresh.
// Usage is similar to useState except we pass in a local storage key so that
// we can default to that value on page load instead of the specified initial
// value.

import { useState } from 'react';

// Usage

function App() {

  // Similar to useState but first arg is key to the value in local storage.
  const [name, setName] = useLocalStorage('name', 'Bob');

  return (
    <div>
      <input
        type="text"
        placeholder="Enter your name"
        value={name}
        onChange={e => setName(e.target.value)}
      />
    </div>
  );
}

// Hook

function useLocalStorage(key, initialValue) {

  // State to store our value
  // Pass initial state function to useState so logic is only executed once

  const [storedValue, setStoredValue] = useState(() => {
    try {
      // Get from local storage by key
      const item = window.localStorage.getItem(key);

      // Parse stored json or if none return initialValue
      return item ? JSON.parse(item) : initialValue;
    } catch (error) {
      // If error also return initialValue
      console.log(error);
      return initialValue;
    }
  });

  // Return a wrapped version of useState's setter function that ...
  // ... persists the new value to localStorage.

  const setValue = value => {
    try {
      // Allow value to be a function so we have same API as useState
      const valueToStore =
        value instanceof Function ? value(storedValue) : value;

      // Save state
      setStoredValue(valueToStore);

      // Save to local storage
      window.localStorage.setItem(key, JSON.stringify(valueToStore));
    } catch (error) {
      // A more advanced implementation would handle the error case
      console.log(error);
    }
  };

  return [storedValue, setValue];
}
