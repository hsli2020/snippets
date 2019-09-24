// ----------------------------------------------------------
// App.svelte
// ----------------------------------------------------------
<script>
	import { scaleLinear } from 'd3-scale';
	import points from './data.js';

	const yTicks = [0, 2, 4, 6, 8];
	const xTicks = [1980, 1990, 2000, 2010];
	const padding = { top: 20, right: 15, bottom: 20, left: 25 };

	let width = 500;
	let height = 200;

	$: xScale = scaleLinear()
		.domain([minX, maxX])
		.range([padding.left, width - padding.right]);

	$: yScale = scaleLinear()
		.domain([Math.min.apply(null, yTicks), Math.max.apply(null, yTicks)])
		.range([height - padding.bottom, padding.top]);

	$: minX = points[0].x;
	$: maxX = points[points.length - 1].x;
	$: path = `M${points.map(p => `${xScale(p.x)},${yScale(p.y)}`).join('L')}`;
	$: area = `${path}L${xScale(maxX)},${yScale(0)}L${xScale(minX)},${yScale(0)}Z`;

	function formatMobile (tick) {
		return "'" + tick % 100;
	}
</script>

<h2>Arctic sea ice minimum</h2>

<div class="chart" bind:clientWidth={width} bind:clientHeight={height}>
	<svg>
		<!-- y axis -->
		<g class="axis y-axis" transform="translate(0, {padding.top})">
			{#each yTicks as tick}
				<g class="tick tick-{tick}" transform="translate(0, {yScale(tick) - padding.bottom})">
					<line x2="100%"></line>
					<text y="-4">{tick} {tick === 8 ? ' million sq km' : ''}</text>
				</g>
			{/each}
		</g>

		<!-- x axis -->
		<g class="axis x-axis">
			{#each xTicks as tick}
				<g class="tick tick-{ tick }" transform="translate({xScale(tick)},{height})">
					<line y1="-{height}" y2="-{padding.bottom}" x1="0" x2="0"></line>
					<text y="-2">{width > 380 ? tick : formatMobile(tick)}</text>
				</g>
			{/each}
		</g>

		<!-- data -->
		<path class="path-area" d={area}></path>
		<path class="path-line" d={path}></path>
	</svg>
</div>

<p>Average September extent. Source: <a href='https://climate.nasa.gov/vital-signs/arctic-sea-ice/'>NSIDC/NASA</a></p>

<style>
	.chart, h2, p {
		width: 100%;
		max-width: 500px;
		margin-left: auto;
		margin-right: auto;
	}

	svg {
		position: relative;
		width: 100%;
		height: 200px;
		overflow: visible;
	}

	.tick {
		font-size: .725em;
		font-weight: 200;
	}

	.tick line {
		stroke: #aaa;
		stroke-dasharray: 2;
	}

	.tick text {
		fill: #666;
		text-anchor: start;
	}

	.tick.tick-0 line {
		stroke-dasharray: 0;
	}

	.x-axis .tick text {
		text-anchor: middle;
	}

	.path-line {
		fill: none;
		stroke: rgb(0,100,100);
		stroke-linejoin: round;
		stroke-linecap: round;
		stroke-width: 2;
	}

	.path-area {
		fill: rgba(0,100,100,0.2);
	}
</style>

// ----------------------------------------------------------
// data.js
// ----------------------------------------------------------
export default [
	{ x: 1979, y: 7.19 },
	{ x: 1980, y: 7.83 },
	{ x: 1981, y: 7.24 },
	{ x: 1982, y: 7.44 },
	{ x: 1983, y: 7.51 },
	{ x: 1984, y: 7.10 },
	{ x: 1985, y: 6.91 },
	{ x: 1986, y: 7.53 },
	{ x: 1987, y: 7.47 },
	{ x: 1988, y: 7.48 },
	{ x: 1989, y: 7.03 },
	{ x: 1990, y: 6.23 },
	{ x: 1991, y: 6.54 },
	{ x: 1992, y: 7.54 },
	{ x: 1993, y: 6.50 },
	{ x: 1994, y: 7.18 },
	{ x: 1995, y: 6.12 },
	{ x: 1996, y: 7.87 },
	{ x: 1997, y: 6.73 },
	{ x: 1998, y: 6.55 },
	{ x: 1999, y: 6.23 },
	{ x: 2000, y: 6.31 },
	{ x: 2001, y: 6.74 },
	{ x: 2002, y: 5.95 },
	{ x: 2003, y: 6.13 },
	{ x: 2004, y: 6.04 },
	{ x: 2005, y: 5.56 },
	{ x: 2006, y: 5.91 },
	{ x: 2007, y: 4.29 },
	{ x: 2008, y: 4.72 },
	{ x: 2009, y: 5.38 },
	{ x: 2010, y: 4.92 },
	{ x: 2011, y: 4.61 },
	{ x: 2012, y: 3.62 },
	{ x: 2013, y: 5.35 },
	{ x: 2014, y: 5.28 },
	{ x: 2015, y: 4.63 },
	{ x: 2016, y: 4.72 }
];