// ----------------------------------------------------------
// App.svelte
// ----------------------------------------------------------
<script>
	import { crossfade, scale } from 'svelte/transition';
	import images from './images.js';

	const [send, receive] = crossfade({
		duration: 200,
		fallback: scale
	});

	let selected = null;
	let loading = null;

	const ASSETS = `https://sveltejs.github.io/assets/crossfade`;

	const load = image => {
		const timeout = setTimeout(() => loading = image, 100);

		const img = new Image();

		img.onload = () => {
			selected = image;
			clearTimeout(timeout);
			loading = null;
		};

		img.src = `${ASSETS}/${image.id}.jpg`;
	};
</script>

<div class="container">
	<div class="phone">
		<h1>Photo gallery</h1>

		<div class="grid">
			{#each images as image}
				<div class="square">
					{#if selected !== image}
						<button
							style="background-color: {image.color};"
							on:click="{() => load(image)}"
							in:receive={{key:image.id}}
							out:send={{key:image.id}}
						>{loading === image ? '...' : image.id}</button>
					{/if}
				</div>
			{/each}
		</div>

		{#if selected}
			{#await selected then d}
				<div class="photo" in:receive={{key:d.id}} out:send={{key:d.id}}>
					<img
						alt={d.alt}
						src="{ASSETS}/{d.id}.jpg"
						on:click="{() => selected = null}"
					>

					<p class='credit'>
						<a target="_blank" href="https://www.flickr.com/photos/{d.path}">via Flickr</a> &ndash;
						<a target="_blank" href={d.license.url}>{d.license.name}</a>
					</p>
				</div>
			{/await}
		{/if}
	</div>
</div>

<style>
	.container {
		position: absolute;
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100%;
		top: 0;
		left: 0;
	}

	.phone {
		position: relative;
		display: flex;
		flex-direction: column;
		width: 52vmin;
		height: 76vmin;
		border: 2vmin solid #ccc;
		border-bottom-width: 10vmin;
		padding: 3vmin;
		border-radius: 2vmin;
	}

	h1 {
		font-weight: 300;
		text-transform: uppercase;
		font-size: 5vmin;
		margin: 0.2em 0 0.5em 0;
	}

	.grid {
		display: grid;
		flex: 1;
		grid-template-columns: repeat(3, 1fr);
		grid-template-rows: repeat(4, 1fr);
		grid-gap: 2vmin;
	}

	button {
		width: 100%;
		height: 100%;
		color: white;
		font-size: 5vmin;
		border: none;
		margin: 0;
		will-change: transform;
	}

	.photo, img {
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		overflow: hidden;
	}

	.photo {
		display: flex;
		align-items: stretch;
		justify-content: flex-end;
		flex-direction: column;
		will-change: transform;
	}

	img {
		object-fit: cover;
		cursor: pointer;
	}

	.credit {
		text-align: right;
		font-size: 2.5vmin;
		padding: 1em;
		margin: 0;
		color: white;
		font-weight: bold;
		opacity: 0.6;
		background: rgba(0,0,0,0.4);
	}

	.credit a, .credit a:visited {
		color: white;
	}
</style>

// ----------------------------------------------------------
// images.js
// ----------------------------------------------------------
const BY = {
	name: 'CC BY 2.0',
	url: 'https://creativecommons.org/licenses/by/2.0/'
};

const BY_SA = {
	name: 'CC BY-SA 2.0',
	url: 'https://creativecommons.org/licenses/by-sa/2.0/'
};

const BY_ND = {
	name: 'CC BY-ND 2.0',
	url: 'https://creativecommons.org/licenses/by-nd/2.0/'
};

// via http://labs.tineye.com/multicolr
export default [
	{
		color: '#001f3f',
		id: '1',
		alt: 'Crepuscular rays',
		path: '43428526@N03/7863279376',
		license: BY
	},
	{
		color: '#0074D9',
		id: '2',
		alt: 'Lapland winter scene',
		path: '25507134@N00/6527537485',
		license: BY
	},
	{
		color: '#7FDBFF',
		id: '3',
		alt: 'Jellyfish',
		path: '37707866@N00/3354331318',
		license: BY
	},
	{
		color: '#39CCCC',
		id: '4',
		alt: 'A man scuba diving',
		path: '32751486@N00/4608886209',
		license: BY_SA
	},
	{
		color: '#3D9970',
		id: '5',
		alt: 'Underwater scene',
		path: '25483059@N08/5548569010',
		license: BY
	},
	{
		color: '#2ECC40',
		id: '6',
		alt: 'Ferns',
		path: '8404611@N06/2447470760',
		license: BY
	},
	{
		color: '#01FF70',
		id: '7',
		alt: 'Posters in a bar',
		path: '33917831@N00/114428206',
		license: BY_SA
	},
	{
		color: '#FFDC00',
		id: '8',
		alt: 'Daffodil',
		path: '46417125@N04/4818617089',
		license: BY_ND
	},
	{
		color: '#FF851B',
		id: '9',
		alt: 'Dust storm in Sydney',
		path: '56068058@N00/3945496657',
		license: BY
	},
	{
		color: '#FF4136',
		id: '10',
		alt: 'Postbox',
		path: '31883499@N05/4216820032',
		license: BY
	},
	{
		color: '#85144b',
		id: '11',
		alt: 'Fireworks',
		path: '8484971@N07/2625506561',
		license: BY_ND
	},
	{
		color: '#B10DC9',
		id: '12',
		alt: 'The Stereophonics',
		path: '58028312@N00/5385464371',
		license: BY_ND
	}
];