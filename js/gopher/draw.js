"use strict";

if(window.animationLoop) {
	window.cancelAnimationFrame(window.animationLoop)
}

class HSL {
	constructor(h, s, l) {
		this.h = h
		this.s = s
		this.l = l
	}

	toString() {
		return "hsl(" +
			(this.h * 360).toFixed(2) + "," +
			(this.s * 100).toFixed(2) + "%," +
			(this.l * 100).toFixed(2) + "%)";
	}
}

const tau = Math.PI * 2

let view = document.getElementById("view")
let context = view.getContext("2d")
updateViewSize()

let gopher = {
	body: new HSL(205.0/360, 1.0,  0.81),
	bodyDark: new HSL(205.0/360, 0.55,  0.71),
	dark: new HSL(217.0/360, 0.19, 0.18),
	light: new HSL(217.0/360, 0.0, 1.00),
	lightDark: new HSL(217.0/360, 0.0, 0.90),
	limb: new HSL(46.0/360, 0.38, 0.80),

	head: {x: 0, y:0},
	gaze: {x: 0, y:0},
};

window.animationLoop = tick();

window.addEventListener("resize", updateViewSize)

function updateViewSize() {
	//view.width = window.innerWidth;
	//view.height = window.innerHeight;

	view.width = view.clientWidth;
	view.height = view.clientHeight;
}

view.addEventListener("mousemove", (ev) => {
	let screenSize = {x: view.width, y: view.height}

	gopher.gaze.x = ev.clientX * 2.0 / view.width - 1;
	gopher.gaze.y = ev.clientY * 2.0 / view.height - 1;

	gopher.head.x = gopher.gaze.x * 0.5;
	gopher.head.y = gopher.gaze.y * 0.5;

	ev.preventDefault()
});

view.addEventListener("touchmove", (ev) => {
	if(ev.targetTouches.length == 0) return;

	let screenSize = {x: view.width, y: view.height}
	let touch = ev.targetTouches[0]

	gopher.gaze.x = touch.clientX * 2.0 / view.width - 1;
	gopher.gaze.y = touch.clientY * 2.0 / view.height - 1;

	gopher.head.x = gopher.gaze.x * 0.5;
	gopher.head.y = gopher.gaze.y * 0.5;

	ev.preventDefault()
});

function tick(t) {
	context.clearRect(0,0,10000,10000);
	context.save()

	// context.fillStyle = "#000";
	// context.fillText(stringify(gopher.head), 20, 20)
	// context.fillText(stringify(gopher.gaze), 20, 30)

	let screenSize = {x: view.width, y: view.height}
	render(gopher, context, screenSize)
	context.restore()
	return window.requestAnimationFrame(tick)
}

function stringify(v) {
	return JSON.stringify(v, (key, val) => {
		if(typeof val == "number"){
			return val.toFixed(2)
		}
		return val
	})
}

function render(gopher, draw, screenSize) {
	let bodyRadius = Math.min(screenSize.x, screenSize.y) / 2
	bodyRadius *= 0.8

	let eyeRadius = bodyRadius * 0.4
	let earRadius = eyeRadius * 0.53
	let lineWidth = bodyRadius * 0.045
	let highlightSize = eyeRadius * 0.06

	draw.lineWidth = lineWidth
	draw.strokeStyle = gopher.dark.toString()

	// The center of the drawing is in the bottom-middle.
	draw.translate(screenSize.x/2, screenSize.y)

	for(let ear = -1; ear <= 1; ear+=2){
		saved(draw, (draw) => { // ear
			let earOffset = headOffset(gopher, -ear, bodyRadius)
			earOffset.x *= -1
			earOffset.x += ear*earRadius*0.3
			earOffset.y = -bodyRadius*1.8 - Math.sin(gopher.head.y)*earRadius * 0.5

			draw.translate(earOffset.x, earOffset.y)
			draw.scale(ear, 1)
			draw.rotate(tau / 10)
			draw.translate(0, earRadius)

			let earPath = new Path2D()

			earPath.moveTo(-earRadius, earRadius)
			earPath.lineTo(-earRadius, -earRadius)
			earPath.bezierCurveTo(
				-earRadius, -earRadius-earRadius*1.3,
				earRadius, -earRadius-earRadius*1.3,
				earRadius, -earRadius)
			earPath.lineTo(earRadius, earRadius)

			draw.fillStyle = gopher.body.toString()
			draw.fill(earPath)
			draw.stroke(earPath)

			let earInline = new Path2D()
			earInline.moveTo(0, 0)
			earInline.quadraticCurveTo(
				-earRadius*0.6, -earRadius,
				earRadius*0.2, -earRadius*1.4
			);
			earInline.moveTo(0, 0)
			earInline.quadraticCurveTo(
				-earRadius*0.6, -earRadius,
				earRadius*0.3, -earRadius*1.0
			);

			draw.lineWidth = lineWidth * 0.7
			draw.lineCap = "round"
			draw.stroke(earInline)
		})
	}

	saved(draw, (draw) => { // body
		let bodyPath = new Path2D()

		bodyPath.moveTo(-bodyRadius, 0)
		bodyPath.lineTo(-bodyRadius, -bodyRadius)
		bodyPath.bezierCurveTo(
			-bodyRadius, -bodyRadius-bodyRadius*1.2,
			bodyRadius, -bodyRadius-bodyRadius*1.2,
			bodyRadius, -bodyRadius)
		bodyPath.lineTo(bodyRadius, 0)

		draw.fillStyle = gopher.bodyDark.toString()
		draw.fill(bodyPath)
		saved(draw, (draw) => {
			draw.clip(bodyPath)
			draw.translate(bodyRadius * 0.1, 0)
			draw.fillStyle = gopher.body.toString()
			draw.fill(bodyPath)
		})
		draw.stroke(bodyPath)
	})

	for(let eye = -1; eye <= 1; eye+=2){
		let gaze = gopher.gaze;

		saved(draw, (draw) => { // eye
			let eyeOffset = headOffset(gopher, eye, bodyRadius);
			draw.translate(eyeOffset.x, eyeOffset.y);

			let eyePath = new Path2D();
			eyePath.arc(0, 0, eyeRadius, 0, tau, true);

			draw.fillStyle = gopher.lightDark.toString()
			draw.fill(eyePath)

			saved(draw, (draw) => {
				draw.clip(eyePath)
				draw.translate(eyeRadius*0.2, -eyeRadius*0.15)
				draw.fillStyle = gopher.light.toString()
				draw.fill(eyePath)
			})

			draw.stroke(eyePath)

			let pupil = { x: gaze.x, y: gaze.y}
			let mag = Math.sqrt(pupil.x*pupil.x + pupil.y*pupil.y)
			if(mag > 1){
				pupil.x /= mag + 0.0001;
				pupil.y /= mag + 0.0001;
			}
			pupil.x = Math.sin(pupil.x + eye * 0.4) * eyeRadius * 0.6
			pupil.y = Math.sin(pupil.y) * eyeRadius * 0.6
			let pupilSize = eyeRadius*0.25

			let pupilPath = new Path2D();
			pupilPath.arc(pupil.x, pupil.y, pupilSize, 0, tau, true)
			draw.fillStyle = gopher.dark.toString()
			draw.fill(pupilPath)

			// pupil highlight
			let highlighPath = new Path2D()
			highlighPath.arc(
				pupil.x + pupilSize * 0.3 - gaze.x * highlightSize * 0.3,
				pupil.y - pupilSize * 0.3 - gaze.y * highlightSize * 0.3, highlightSize, 0, tau, true)
			draw.fillStyle = gopher.light.toString()
			draw.fill(highlighPath)
		})
	}

	saved(draw, (draw) => { // nose
		let tipSize = eyeRadius*0.2;
		let toothSize = tipSize;
		let noseSize = tipSize * 2.2;

		let noseOffset = headOffset(gopher, 0, bodyRadius)
		noseOffset.y += eyeRadius * 0.8;
		draw.translate(noseOffset.x, noseOffset.y);

		saved(draw, (draw) => { // tooth
			draw.translate(0, tipSize*1.2 - gopher.head.y * tipSize * 0.3)
			draw.beginPath()
			draw.moveTo(-toothSize, 0)
			draw.lineTo(-toothSize, toothSize*1.5)
			draw.bezierCurveTo(
				-toothSize, toothSize*2,
				toothSize, toothSize*2,
				toothSize, toothSize*1.5,
			)
			draw.lineTo(toothSize, 0)
			draw.fillStyle = gopher.light.toString()
			draw.fill()
			draw.lineWidth = lineWidth * 0.7;
			draw.stroke()
		})

		saved(draw, (draw) => { // nose
			draw.translate(0, tipSize*1.2)
			draw.beginPath()

			draw.moveTo(-noseSize, 0)
			draw.bezierCurveTo(
				-noseSize, -noseSize,
				noseSize, -noseSize,
				noseSize, 0
			)
			draw.bezierCurveTo(
				noseSize, noseSize*0.4,
				-noseSize, noseSize*0.4,
				-noseSize, 0
			)
			draw.closePath()

			draw.fillStyle = gopher.limb.toString()
			draw.fill()
			draw.stroke()
		})

		saved(draw, (draw) => { // tip of the nose
			let tip = {
				x: Math.sin(gopher.head.x) * tipSize,
				y: Math.sin(gopher.head.y) * tipSize * 0.5 - tipSize*0.2,
			};

			draw.beginPath()
			draw.scale(1.2, 0.9)
			draw.arc(tip.x, tip.y, tipSize, 0, tau, true)
			draw.fillStyle = gopher.dark.toString()
			draw.fill()

			// tip highlight
			draw.beginPath()
			draw.arc(
				tip.x + tipSize * 0.3 - gopher.head.x * highlightSize * 0.3,
				tip.y - tipSize * 0.3 - gopher.head.y * highlightSize * 0.3, highlightSize, 0, tau, true)
			draw.fillStyle = gopher.light.toString()
			draw.fill()
		})
	})
}

function headOffset(gopher, eye, bodyRadius) {
	let x = eye*bodyRadius*0.65
	let y = -bodyRadius*1.1

	x += Math.sin(gopher.head.x - eye*0.4) * bodyRadius*0.3
	y += Math.sin(gopher.head.y*tau/4) * bodyRadius*0.15

	return {x: x, y: y}
}

function saved(draw, fn) {
	draw.save()
	fn(draw)
	draw.restore()
}
