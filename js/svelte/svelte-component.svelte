// ----------------------------------------------------------
// App.svelte
// ----------------------------------------------------------
<script>
	import RedThing from './RedThing.svelte';
	import GreenThing from './GreenThing.svelte';
	import BlueThing from './BlueThing.svelte';

	const options = [
		{ color: 'red',   component: RedThing   },
		{ color: 'green', component: GreenThing },
		{ color: 'blue',  component: BlueThing  },
	];

	let selected = options[0];
</script>

<select bind:value={selected}>
	{#each options as option}
		<option value={option}>{option.color}</option>
	{/each}
</select>

<svelte:component this={selected.component}/>

// ----------------------------------------------------------
// BlueThing.svelte
// ----------------------------------------------------------
<style>
	strong { color: blue; }
</style>

<strong>blue thing</strong>

// ----------------------------------------------------------
// RedThing.svelte
// ----------------------------------------------------------
<style>
	strong { color: red; }
</style>

<strong>red thing</strong>

// ----------------------------------------------------------
// GreenThing.svelte
// ----------------------------------------------------------
<style>
	strong { color: green; }
</style>

<strong>green thing</strong>