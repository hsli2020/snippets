// ----------------------------------------------------------
// App.svelte
// ----------------------------------------------------------

<script>
	import Outer from './Outer.svelte';

	function handleMessage(event) {
		alert(event.detail.text);
	}
</script>

<Outer on:message={handleMessage}/>

// ----------------------------------------------------------
// Inner.svelte
// ----------------------------------------------------------

<script>
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function sayHello() {
		dispatch('message', {
			text: 'Hello!'
		});
	}
</script>

<button on:click={sayHello}>
	Click to say hello
</button>

// ----------------------------------------------------------
// Outer.svelte
// ----------------------------------------------------------

<script>
	import Inner from './Inner.svelte';
</script>

<Inner on:message/>