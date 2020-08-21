/**
 * You're going to be building a small app with the following requirements:
 *   - Three components: App, Heading, and Button
 *   - Every time the button is clicked, the heading will update with the current click count
 *   - Every time the button is clicked, the color of the button will change
 */

////////////////////////////////////////////////////////////////////////////////
// Svelte

// App.svelte,

<script>
  import Button from './Button.svelte';
  import Heading from './Heading.svelte';

  let count = 0;
  let color = '#000000';

  const colors = ['#00ff00', '#ff0000', '#0000ff'];
  
  let handleClick = () => {
    count++;
    color = colors[Math.floor(Math.random() * 3)];
  }
</script> 

<main>
  <Heading count={count} />
  <Button color={color} handleClick={handleClick} />
</main>

// Heading.svelte

<script>
  export let count; // prop
</script>

<h1>Hello, I am a Svelte App!</h1>
<h2>The following button has been clicked {count} times.</h2>

// Button.svelte.

<script>
  export let handleClick; // prop
  export let color; // prop
</script>

<button style="--color: {color}" on:click={handleClick}>
  Click me!
</button>

<style>
  button {
    color: white;
    background-color: var(--color);
  }
</style>

////////////////////////////////////////////////////////////////////////////////
// React

// App.js

import Heading from './Heading.js';
import Button from './Button.js';
import { useState } from 'react';

function App() {
  const [count, setCount] = useState(0);
  const [color, setColor] = useState('#000000');

  const colors = ['#00ff00', '#ff0000', '#0000ff'];
  
  let handleClick = () => {
    setCount(count+1);
    setColor(colors[Math.floor(Math.random() * 3)]);
  }

  return (
    <main>
      <Heading count={count} />     
      <Button color={color} handleClick={handleClick} />
    </main>
  )
}

export default App;

// Heading.js

function Heading({ count }) { // prop
  return (
    <div>
      <h1>Hello, I am a React App!</h1>
      <h2>The following button has been clicked {count} times.</h2>
    </div>
  )
}

export default Heading;

// Button.js

function Button({ color, handleClick }) { // prop
  const styles = {
    backgroundColor: color,
    color: '#ffffff'
  }
  return (
    <button style={styles} onClick={handleClick}>
      Click me!
    </button>
  )
}

export default Button;
