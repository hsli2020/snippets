// ----------------------------------------------------------
// App.svelte
// ----------------------------------------------------------
<script>
	import { onMount } from 'svelte';
	import List from './List.svelte';
	import Item from './Item.svelte';

	let item;
	let page;

	async function hashchange() {
		// the poor man's router!
		const path = window.location.hash.slice(1);

		if (path.startsWith('/item')) {
			const id = path.slice(6);
			item = await fetch(`https://node-hnapi.herokuapp.com/item/${id}`).then(r => r.json());

			window.scrollTo(0,0);
		} else if (path.startsWith('/top')) {
			page = +path.slice(5);
			item = null;
		} else {
			window.location.hash = '/top/1';
		}
	}

	onMount(hashchange);
</script>

<style>
	main {
		position: relative;
		max-width: 800px;
		margin: 0 auto;
		min-height: 101vh;
		padding: 1em;
	}

	main :global(.meta) {
		color: #999;
		font-size: 12px;
		margin: 0 0 1em 0;
	}

	main :global(a) {
		color: rgb(0,0,150);
	}
</style>

<svelte:window on:hashchange={hashchange}/>

<main>
	{#if item}
		<Item {item} returnTo="#/top/{page}"/>
	{:else if page}
		<List {page}/>
	{/if}
</main>

// ----------------------------------------------------------
// Comment.svelte
// ----------------------------------------------------------
<script>
	export let comment;
</script>

<style>
	article {
		border-top: 1px solid #eee;
		margin: 1em 0 0 0;
		padding: 1em 0 0 0;
		font-size: 14px;
	}

	.meta {
		color: #999;
	}

	.replies {
		padding: 0 0 0 1em;
	}
</style>

<article>
	<p class="meta">{comment.user} {comment.time_ago}</p>

	{@html comment.content}

	<div class="replies">
		{#each comment.comments as child}
			<svelte:self comment={child}/>
		{/each}
	</div>
</article>

// ----------------------------------------------------------
// Item.svelte
// ----------------------------------------------------------
<script>
	import Comment from "./Comment.svelte";

	export let item;
	export let returnTo;
</script>

<style>
	article {
		margin: 0 0 1em 0;
	}

	a {
		display: block;
		margin: 0 0 1em 0;
	}

	h1 {
		font-size: 1.4em;
		margin: 0;
	}
</style>

<a href={returnTo}>&laquo; back</a>

<article>
	<a href="{item.url}">
		<h1>{item.title}</h1>
		<small>{item.domain}</small>
	</a>

	<p class="meta">submitted by {item.user} {item.time_ago}
</article>

<div class="comments">
	{#each item.comments as comment}
		<Comment {comment}/>
	{/each}
</div>

// ----------------------------------------------------------
// List.svelte
// ----------------------------------------------------------
<script>
	import { beforeUpdate } from "svelte";
	import Summary from "./Summary.svelte";

	const PAGE_SIZE = 20;

	export let page;

	let items;
	let offset;

	$: fetch(`https://node-hnapi.herokuapp.com/news?page=${page}`)
		.then(r => r.json())
		.then(data => {
			items = data;
			offset = PAGE_SIZE * (page - 1);
			window.scrollTo(0, 0);
		});
</script>

<style>
	a {
		padding: 2em;
		display: block;
	}

	.loading {
		opacity: 0;
		animation: 0.4s 0.8s forwards fade-in;
	}

	@keyframes fade-in {
		from { opacity: 0; }
		to { opacity: 1; }
	}
</style>

{#if items}
	{#each items as item, i}
		<Summary {item} {i} {offset}/>
	{/each}

	<a href="#/top/{page + 1}">page {page + 1}</a>
{:else}
	<p class="loading">loading...</p>
{/if}

// ----------------------------------------------------------
// Summary.svelte
// ----------------------------------------------------------
<script>
	export let item;
	export let i;
	export let offset;

	function comment_text() {
		const c = item.comments_count;
		return `${c} ${c === 1 ? 'comment' : 'comments'}`;
	}
</script>

<style>
	article {
		position: relative;
		padding: 0 0 0 2em;
		border-bottom: 1px solid #eee;
	}

	h2 {
		font-size: 1em;
		margin: 0.5em 0;
	}

	span {
		position: absolute;
		left: 0;
	}

	a {
		color: #333;
	}
</style>

<article>
	<span>{i + offset + 1}</span>
	<h2><a target="_blank" href={item.url}>{item.title}</a></h2>
	<p class="meta"><a href="#/item/{item.id}">{comment_text()}</a> by {item.user} {item.time_ago}</p>
</article>