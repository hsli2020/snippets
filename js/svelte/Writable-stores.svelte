// ----------------------------------------------------------
// App.svelte
// ----------------------------------------------------------
<script>
	import { count } from './stores.js';
	import Incrementer from './Incrementer.svelte';
	import Decrementer from './Decrementer.svelte';
	import Resetter from './Resetter.svelte';
</script>

<h1>The count is {$count}</h1>

<Incrementer/>
<Decrementer/>
<Resetter/>

// ----------------------------------------------------------
// Incrementer.svelte
// ----------------------------------------------------------
<script>
	import { count } from './stores.js';

	function decrement() {
		count.update(n => n - 1);
	}
</script>

<button on:click={decrement}> - </button>

// ----------------------------------------------------------
// Decrementer.svelte
// ----------------------------------------------------------
<script>
	import { count } from './stores.js';

	function increment() {
		count.update(n => n + 1);
	}
</script>

<button on:click={increment}> + </button>

// ----------------------------------------------------------
// Resetter.svelte
// ----------------------------------------------------------
<script>
	import { count } from './stores.js';

	function reset() {
		count.set(0);
	}
</script>

<button on:click={reset}> reset </button>

// ----------------------------------------------------------
// store.js
// ----------------------------------------------------------
import { writable } from 'svelte/store';

export const count = writable(0);