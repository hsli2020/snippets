﻿/* WGo.js 2.01 */

(function(window, undefined) {

"use strict";

var scripts= document.getElementsByTagName('script');
var path= scripts[scripts.length-1].src.split('?')[0];      // remove any ?query
var mydir= path.split('/').slice(0, -1).join('/')+'/';

/**
 * Main namespace - it initializes WGo in first run and then execute main function.
 * You must call WGo.init() if you want to use library, without calling WGo.
 */
var WGo = {
	// constants for colors (rather use WGo.B or WGo.W)
	B: 1,
	W: -1,

	// if true errors will be shown in dialog window, otherwise they will be ignored
	ERROR_REPORT: true,
	DIR: mydir,
}

// helping function for class inheritance
WGo.extendClass = function(parent, child) {
	child.prototype = Object.create(parent.prototype);
	child.prototype.constructor = child;
	child.prototype.super = parent;

	return child;
};

// helping function for class inheritance
WGo.abstractMethod = function() {
	throw Error('unimplemented abstract method');
};

// helping function for deep cloning of simple objects,
WGo.clone = function(obj) {
	if(obj && typeof obj == "object") {
		var n_obj = obj.constructor == Array ? [] : {};

		for(var key in obj) {
			if(obj[key] == obj) n_obj[key] = obj;
			else n_obj[key] = WGo.clone(obj[key]);
		}

		return n_obj;
	}
	else return obj;
}

// filter html to avoid XSS
WGo.filterHTML = function(text) {
	if(!text || typeof text != "string") return text;
	return text.replace(/</g, "&lt;").replace(/>/g, "&gt;");
}

//---------------------- WGo.Board -----------------------------------------------------------------------------

/**
 * Board class constructor - it creates a canvas board
 *
 * @param elem DOM element to put in
 * @param config configuration object. It is object with "key: value" structure. Possible configurations are:
 *
 * - size: number - size of the board (default: 19)
 * - width: number - width of the board (default: 0)
 * - height: number - height of the board (default: 0)
 * - font: string - font of board writings (!deprecated)
 * - lineWidth: number - line width of board drawings (!deprecated)
 * - autoLineWidth: boolean - if set true, line width will be automatically computed accordingly to board size - this option rewrites 'lineWidth' /and it will keep markups sharp/ (!deprecated)
 * - starPoints: Object - star points coordinates, defined for various board sizes. Look at Board.default for more info.
 * - stoneHandler: Board.DrawHandler - stone drawing handler (default: Board.drawHandlers.SHELL)
 * - starSize: number - size of star points (default: 1). Radius of stars is dynamic, however you can modify it by given constant. (!deprecated)
 * - stoneSize: number - size of stone (default: 1). Radius of stone is dynamic, however you can modify it by given constant. (!deprecated)
 * - shadowSize: number - size of stone shadow (default: 1). Radius of shadow is dynamic, however you can modify it by given constant. (!deprecated)
 * - background: string - background of the board, it can be either color (#RRGGBB) or url. Empty string means no background. (default: WGo.DIR+"wood1.jpg")
 * - section: {
 *     top: number,
 *     right: number,
 *     bottom: number,
 *     left: number
 *   }
 *   It defines a section of board to be displayed. You can set a number of rows(or cols) to be skipped on each side.
 *   Numbers can be negative, in that case there will be more empty space. In default all values are zeros.
 * - theme: Object - theme object, which defines all graphical attributes of the board. Default theme object is "WGo.Board.themes.default". For old look you may use "WGo.Board.themes.old".
 *
 * Note: properties lineWidth, autoLineWidth, starPoints, starSize, stoneSize and shadowSize will be considered only if you set property 'theme' to 'WGo.Board.themes.old'.
 */

var Board = function(elem, config) {
	var config = config || {};

	// set user configuration
	for(var key in config) this[key] = config[key];

	// add default configuration
	for(var key in WGo.Board.default) if(this[key] === undefined) this[key] = WGo.Board.default[key];

	// add default theme variables
	for(var key in Board.themes.default) if(this.theme[key] === undefined) this.theme[key] = Board.themes.default[key];

	// set section if set
	this.tx = this.section.left;
	this.ty = this.section.top;
	this.bx = this.size-1-this.section.right;
	this.by = this.size-1-this.section.bottom;

	// init board
	this.init();

	// append to element
	elem.appendChild(this.element);

	// set initial dimensions
	if(this.width && this.height) this.setDimensions(this.width, this.height);
	else if(this.width) this.setWidth(this.width);
	else if(this.height) this.setHeight(this.height);
}

// New experimental board theme system - it can be changed in future, if it will appear to be unsuitable.
Board.themes = {};

Board.themes.old = {
	shadowColor: "rgba(32,32,32,0.5)",
	shadowSize: function(board){
		return board.shadowSize;
	},
	markupBlackColor: "rgba(255,255,255,0.8)",
	markupWhiteColor: "rgba(0,0,0,0.8)",
	markupNoneColor: "rgba(0,0,0,0.8)",
	markupLinesWidth: function(board) {
		return board.autoLineWidth ? board.stoneRadius/7 : board.lineWidth;
	},
	gridLinesWidth: 1,
	gridLinesColor: function(board) {
		return "rgba(0,0,0,"+Math.min(1, board.stoneRadius/15)+")";
	},
	starColor: "#000",
	starSize: function(board) {
		return board.starSize*((board.width/300)+1);
	},
	stoneSize: function(board) {
		return board.stoneSize*Math.min(board.fieldWidth, board.fieldHeight)/2;
	},
	coordinatesColor: "rgba(0,0,0,0.7)",
	font: function(board) {
		return board.font;
	},
	linesShift: 0.5
}

/**
 * Object containing default graphical properties of a board.
 * A value of all properties can be even static value or function, returning final value.
 * Theme object doesn't set board and stone textures - they are set separately.
 */

Board.themes.default = {
	shadowColor: "rgba(62,32,32,0.5)",
	shadowSize: 1,
	markupBlackColor: "rgba(255,255,255,0.9)",
	markupWhiteColor: "rgba(0,0,0,0.7)",
	markupNoneColor: "rgba(0,0,0,0.7)",
	markupLinesWidth: function(board) {
		return board.stoneRadius/8;
	},
	gridLinesWidth: function(board) {
		return board.stoneRadius/15;
	},
	gridLinesColor: "#654525",
	starColor: "#531",
	starSize: function(board) {
		return (board.stoneRadius/8)+1;
	},
	stoneSize: function(board) {
		return Math.min(board.fieldWidth, board.fieldHeight)/2;
	},
	coordinatesColor: "#531",
	font: "calibri",
	linesShift: 0.25
}

var theme_variable = function(key, board) {
	return typeof board.theme[key] == "function" ? board.theme[key](board) : board.theme[key];
}

var shadow_handler = {
	draw: function(args, board) {
		var xr = board.getX(args.x),
			yr = board.getY(args.y),
			sr = board.stoneRadius;

		this.beginPath();
		this.fillStyle = theme_variable("shadowColor", board);
		this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
		this.fill();
	}
}

var get_markup_color = function(board, x, y) {
	if(board.obj_arr[x][y][0].c == WGo.B) return theme_variable("markupBlackColor", board);
	else if(board.obj_arr[x][y][0].c == WGo.W) return theme_variable("markupWhiteColor", board);
	return theme_variable("markupNoneColor", board);
}

var is_here_stone = function(board, x, y) {
	return (board.obj_arr[x][y][0] && board.obj_arr[x][y][0].c == WGo.W || board.obj_arr[x][y][0].c == WGo.B);
}

var redraw_layer = function(board, layer) {
	var handler;

	board[layer].clear();
	board[layer].draw(board);

	for(var x = 0; x < board.size; x++) {
		for(var y = 0; y < board.size; y++) {
			for(var key in board.obj_arr[x][y]) {
				if(!board.obj_arr[x][y][key].type) handler = board.stoneHandler;
				else if(typeof board.obj_arr[x][y][key].type == "string") handler = Board.drawHandlers[board.obj_arr[x][y][key].type];
				else handler = board.obj_arr[x][y][key].type;

				if(handler[layer]) handler[layer].draw.call(board[layer].context, board.obj_arr[x][y][key], board);
			}
		}
	}

	for(var key in board.obj_list) {
		var handler = board.obj_list[key].handler;

		if(handler[layer]) handler[layer].draw.call(board[layer].context, board.obj_list[key].args, board);
	}
}

// shell stone helping functions

var shell_seed;

var draw_shell_line = function(ctx, x, y, radius, start_angle, end_angle, factor, thickness) {
	ctx.strokeStyle = "rgba(64,64,64,0.2)";

	ctx.lineWidth = (radius/30)*thickness;
	ctx.beginPath();

	radius -= Math.max(1, ctx.lineWidth);

	var x1 = x + radius*Math.cos(start_angle*Math.PI);
	var y1 = y + radius*Math.sin(start_angle*Math.PI);
	var x2 = x + radius*Math.cos(end_angle*Math.PI);
	var y2 = y + radius*Math.sin(end_angle*Math.PI);

	var m, angle, x, diff_x, diff_y;
	if(x2 > x1) {
		m = (y2-y1)/(x2-x1);
		angle = Math.atan(m);
	}
	else if(x2 == x1) {
		angle = Math.PI/2;
	}
	else {
		m = (y2-y1)/(x2-x1);
		angle = Math.atan(m)-Math.PI;
	}

	var c = factor*radius;
	diff_x = Math.sin(angle) * c;
	diff_y = Math.cos(angle) * c;

	var bx1 = x1 + diff_x;
	var by1 = y1 - diff_y;

	var bx2 = x2 + diff_x;
	var by2 = y2 - diff_y;

	ctx.moveTo(x1,y1);
	ctx.bezierCurveTo(bx1, by1, bx2, by2, x2, y2);
	ctx.stroke();
}

var draw_shell = function(arg) {
	var from_angle = arg.angle;
	var to_angle = arg.angle;

	for(var i = 0; i < arg.lines.length; i++) {
		from_angle += arg.lines[i];
		to_angle -= arg.lines[i];
		draw_shell_line(arg.ctx, arg.x, arg.y, arg.radius, from_angle, to_angle, arg.factor, arg.thickness);
	}
}

// drawing handlers

Board.drawHandlers = {
	// handler for normal stones
	NORMAL: {
		// draw handler for stone layer
		stone: {
			// drawing function - args object contain info about drawing object, board is main board object
			// this function is called from canvas2D context
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius,
					radgrad;

				// set stone texture
				if(args.c == WGo.W) {
					radgrad = this.createRadialGradient(xr-2*sr/5,yr-2*sr/5,sr/3,xr-sr/5,yr-sr/5,5*sr/5);
					radgrad.addColorStop(0, '#fff');
					radgrad.addColorStop(1, '#aaa');
				}
				else {
					radgrad = this.createRadialGradient(xr-2*sr/5,yr-2*sr/5,1,xr-sr/5,yr-sr/5,4*sr/5);
					radgrad.addColorStop(0, '#666');
					radgrad.addColorStop(1, '#000');
				}

				// paint stone
				this.beginPath();
				this.fillStyle = radgrad;
				this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
				this.fill();
			}
		},
		// adding shadow handler
		shadow: shadow_handler,
	},

	PAINTED: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius,
					radgrad;

				if(args.c == WGo.W) {
					radgrad = this.createRadialGradient(xr-2*sr/5,yr-2*sr/5,2,xr-sr/5,yr-sr/5,4*sr/5);
					radgrad.addColorStop(0, '#fff');
					radgrad.addColorStop(1, '#ddd');
				}
				else {
					radgrad = this.createRadialGradient(xr-2*sr/5,yr-2*sr/5,1,xr-sr/5,yr-sr/5,4*sr/5);
					radgrad.addColorStop(0, '#111');
					radgrad.addColorStop(1, '#000');
				}

				this.beginPath();
				this.fillStyle = radgrad;
				this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
				this.fill();

				this.beginPath();
				this.lineWidth = sr/6;

				if(args.c == WGo.W) {
					this.strokeStyle = '#999';
					this.arc(xr+sr/8, yr+sr/8, sr/2, 0, Math.PI/2, false);
				}
				else {
					this.strokeStyle = '#ccc';
					this.arc(xr-sr/8, yr-sr/8, sr/2, Math.PI, 1.5*Math.PI);
				}

				this.stroke();
			}
		},
		shadow: shadow_handler,
	},

	GLOW: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius;

				var radgrad;
				if(args.c == WGo.W) {
					radgrad = this.createRadialGradient(xr-2*sr/5,yr-2*sr/5,sr/3,xr-sr/5,yr-sr/5,8*sr/5);
					radgrad.addColorStop(0, '#fff');
					radgrad.addColorStop(1, '#666');
				}
				else {
					radgrad = this.createRadialGradient(xr-2*sr/5,yr-2*sr/5,1,xr-sr/5,yr-sr/5,3*sr/5);
					radgrad.addColorStop(0, '#555');
					radgrad.addColorStop(1, '#000');
				}

				this.beginPath();
				this.fillStyle = radgrad;
				this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
				this.fill();
			},
		},
		shadow: shadow_handler,
	},

	SHELL: {
		stone: {
			draw: function(args, board) {
				var xr,
					yr,
					sr = board.stoneRadius;

				shell_seed = shell_seed || Math.ceil(Math.random()*9999999);

				xr = board.getX(args.x);
				yr = board.getY(args.y);

				var radgrad;

				if(args.c == WGo.W) {
					radgrad = "#aaa";
				}
				else {
					radgrad = "#000";
				}

				this.beginPath();
				this.fillStyle = radgrad;
				this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
				this.fill();

				// do shell magic here
				if(args.c == WGo.W) {
					// do shell magic here
					var type = shell_seed%(3+args.x*board.size+args.y)%3;
					var z = board.size*board.size+args.x*board.size+args.y;
					var angle = (2/z)*(shell_seed%z);

					if(type == 0) {
						draw_shell({
							ctx: this,
							x: xr,
							y: yr,
							radius: sr,
							angle: angle,
							lines: [0.10, 0.12, 0.11, 0.10, 0.09, 0.09, 0.09, 0.09],
							factor: 0.25,
							thickness: 1.75
						});
					}
					else if(type == 1) {
						draw_shell({
							ctx: this,
							x: xr,
							y: yr,
							radius: sr,
							angle: angle,
							lines: [0.10, 0.09, 0.08, 0.07, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06, 0.06],
							factor: 0.2,
							thickness: 1.5
						});
					}
					else {
						draw_shell({
							ctx: this,
							x: xr,
							y: yr,
							radius: sr,
							angle: angle,
							lines: [0.12, 0.14, 0.13, 0.12, 0.12, 0.12],
							factor: 0.3,
							thickness: 2
						});
					}
					radgrad = this.createRadialGradient(xr-2*sr/5,yr-2*sr/5,sr/3,xr-sr/5,yr-sr/5,5*sr/5);
					radgrad.addColorStop(0, 'rgba(255,255,255,0.9)');
					radgrad.addColorStop(1, 'rgba(255,255,255,0)');

					// add radial gradient //
					this.beginPath();
					this.fillStyle = radgrad;
					this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
					this.fill();
				}
				else {
					radgrad = this.createRadialGradient(xr+0.4*sr, yr+0.4*sr, 0, xr+0.5*sr, yr+0.5*sr, sr);
					radgrad.addColorStop(0, 'rgba(32,32,32,1)');
					radgrad.addColorStop(1, 'rgba(0,0,0,0)');

					this.beginPath();
					this.fillStyle = radgrad;
					this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
					this.fill();

					radgrad = this.createRadialGradient(xr-0.4*sr, yr-0.4*sr, 1, xr-0.5*sr, yr-0.5*sr, 1.5*sr);
					radgrad.addColorStop(0, 'rgba(64,64,64,1)');
					radgrad.addColorStop(1, 'rgba(0,0,0,0)');

					this.beginPath();
					this.fillStyle = radgrad;
					this.arc(xr-board.ls, yr-board.ls, Math.max(0, sr-board.ls), 0, 2*Math.PI, true);
					this.fill();
				}
			}
		},
		shadow: shadow_handler,
	},

	MONO: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius,
					lw = theme_variable("markupLinesWidth", board) || 1;

				if(args.c == WGo.W) this.fillStyle = "white";
				else this.fillStyle = "black";

				this.beginPath();
				this.arc(xr, yr, Math.max(0, sr-lw), 0, 2*Math.PI, true);
				this.fill();

				this.lineWidth = lw;
				this.strokeStyle = "black";
				this.stroke();
			}
		},
	},

	CR: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius;

				this.strokeStyle = args.c || get_markup_color(board, args.x, args.y);
				this.lineWidth = args.lineWidth || theme_variable("markupLinesWidth", board) || 1;
				this.beginPath();
				this.arc(xr-board.ls, yr-board.ls, sr/2, 0, 2*Math.PI, true);
				this.stroke();
			},
		},
	},

	// Label drawing handler
	LB: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius,
					font = args.font || theme_variable("font", board) || "";

				this.fillStyle = args.c || get_markup_color(board, args.x, args.y);

				if(args.text.length == 1) this.font = Math.round(sr*1.5)+"px "+font;
				else if(args.text.length == 2) this.font = Math.round(sr*1.2)+"px "+font;
				else this.font = Math.round(sr)+"px "+font;

				this.beginPath();
				this.textBaseline="middle";
				this.textAlign="center";
				this.fillText(args.text, xr, yr, 2*sr);

			},
		},

		// modifies grid layer too
		grid: {
			draw: function(args, board) {
				if(!is_here_stone(board, args.x, args.y) && !args._nodraw) {
					var xr = board.getX(args.x),
						yr = board.getY(args.y),
						sr = board.stoneRadius;
					this.clearRect(xr-sr,yr-sr,2*sr,2*sr);
				}
			},
			clear: function(args, board) {
				if(!is_here_stone(board, args.x, args.y))  {
					args._nodraw = true;
					redraw_layer(board, "grid");
					delete args._nodraw;
				}
			}
		},
	},

	SQ: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = Math.round(board.stoneRadius);

				this.strokeStyle = args.c || get_markup_color(board, args.x, args.y);
				this.lineWidth = args.lineWidth || theme_variable("markupLinesWidth", board) || 1;
				this.beginPath();
				this.rect(Math.round(xr-sr/2)-board.ls, Math.round(yr-sr/2)-board.ls, sr, sr);
				this.stroke();
			}
		}
	},

	TR: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius;

				this.strokeStyle = args.c || get_markup_color(board, args.x, args.y);
				this.lineWidth = args.lineWidth || theme_variable("markupLinesWidth", board) || 1;
				this.beginPath();
				this.moveTo(xr-board.ls, yr-board.ls-Math.round(sr/2));
				this.lineTo(Math.round(xr-sr/2)-board.ls, Math.round(yr+sr/3)+board.ls);
				this.lineTo(Math.round(xr+sr/2)+board.ls, Math.round(yr+sr/3)+board.ls);
				this.closePath();
				this.stroke();
			}
		}
	},

	MA: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius;

				this.strokeStyle = args.c || get_markup_color(board, args.x, args.y);
				this.lineCap="round";
				this.lineWidth = (args.lineWidth || theme_variable("markupLinesWidth", board) || 1) * 2 - 1;
				this.beginPath();
				this.moveTo(Math.round(xr-sr/2), Math.round(yr-sr/2));
				this.lineTo(Math.round(xr+sr/2), Math.round(yr+sr/2));
				this.moveTo(Math.round(xr+sr/2)-1, Math.round(yr-sr/2));
				this.lineTo(Math.round(xr-sr/2)-1, Math.round(yr+sr/2));
				this.stroke();
				this.lineCap="butt";
			}
		}
	},

	SL: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius;

				this.fillStyle = args.c || get_markup_color(board, args.x, args.y);
				this.beginPath();
				this.rect(xr-sr/2, yr-sr/2, sr, sr);
				this.fill();
			}
		}
	},

	SM: {
		stone: {
			draw: function(args, board) {
				var xr = board.getX(args.x),
					yr = board.getY(args.y),
					sr = board.stoneRadius;

				this.strokeStyle = args.c || get_markup_color(board, args.x, args.y);
				this.lineWidth = (args.lineWidth || theme_variable("markupLinesWidth", board) || 1)*2;
				this.beginPath();
				this.arc(xr-sr/3, yr-sr/3, sr/6, 0, 2*Math.PI, true);
				this.stroke();
				this.beginPath();
				this.arc(xr+sr/3, yr-sr/3, sr/6, 0, 2*Math.PI, true);
				this.stroke();
				this.beginPath();
				this.moveTo(xr-sr/1.5,yr);
				this.bezierCurveTo(xr-sr/1.5,yr+sr/2,xr+sr/1.5,yr+sr/2,xr+sr/1.5,yr);
				this.stroke();
			}
		}
	},

	outline: {
		stone: {
			draw: function(args, board) {
				if(args.alpha) this.globalAlpha = args.alpha;
				else this.globalAlpha = 0.3;
				if(args.stoneStyle) Board.drawHandlers[args.stoneStyle].stone.draw.call(this, args, board);
				else board.stoneHandler.stone.draw.call(this, args, board);
				this.globalAlpha = 1;
			}
		}
	},

	mini: {
		stone: {
			draw: function(args, board) {
				board.stoneRadius = board.stoneRadius/2;
				if(args.stoneStyle) Board.drawHandlers[args.stoneStyle].stone.draw.call(this, args, board);
				else board.stoneHandler.stone.draw.call(this, args, board);
				board.stoneRadius = board.stoneRadius*2;
			}
		}
	},
}

Board.coordinates = {
	grid: {
		draw: function(args, board) {
			var ch, t, xright, xleft, ytop, ybottom;

			this.fillStyle = theme_variable("coordinatesColor", board);
			this.textBaseline="middle";
			this.textAlign="center";
			this.font = board.stoneRadius+"px "+(board.font || "");

			xright = board.getX(-0.75);
			xleft = board.getX(board.size-0.25);
			ytop = board.getY(-0.75);
			ybottom = board.getY(board.size-0.25);

			for(var i = 0; i < board.size; i++) {
				ch = i+"A".charCodeAt(0);
				if(ch >= "I".charCodeAt(0)) ch++;

				t = board.getY(i);
				this.fillText(board.size-i, xright, t);
				this.fillText(board.size-i, xleft, t);

				t = board.getX(i);
				this.fillText(String.fromCharCode(ch), t, ytop);
				this.fillText(String.fromCharCode(ch), t, ybottom);
			}

			this.fillStyle = "black";
		}
	}
}

Board.CanvasLayer = function() {
	this.element = document.createElement('canvas');
	this.context = this.element.getContext('2d')
}

Board.CanvasLayer.prototype = {
	constructor: Board.CanvasLayer,

	setDimensions: function(width, height) {
		this.element.width = width;
		this.element.height = height;
	},

	draw: function(board) {	},

	clear: function() {
		this.context.clearRect(0,0,this.element.width,this.element.height);
	}
}

Board.GridLayer = WGo.extendClass(Board.CanvasLayer, function() {
	this.super.call(this);
});

Board.GridLayer.prototype.draw = function(board) {
	// draw grid
	var tmp;

	this.context.beginPath();
	this.context.lineWidth = theme_variable("gridLinesWidth", board);
	this.context.strokeStyle = theme_variable("gridLinesColor", board);

	var tx = Math.round(board.left),
		ty = Math.round(board.top),
		bw = Math.round(board.fieldWidth*(board.size-1)),
		bh = Math.round(board.fieldHeight*(board.size-1));

	this.context.strokeRect(tx-board.ls, ty-board.ls, bw, bh);

	for(var i = 1; i < board.size-1; i++) {
		tmp = Math.round(board.getX(i))-board.ls;
		this.context.moveTo(tmp, ty);
		this.context.lineTo(tmp, ty+bh);

		tmp = Math.round(board.getY(i))-board.ls;
		this.context.moveTo(tx, tmp);
		this.context.lineTo(tx+bw, tmp);
	}

	this.context.stroke();

	// draw stars
	this.context.fillStyle = theme_variable("starColor", board);

	if(board.starPoints[board.size]) {
		for(var key in board.starPoints[board.size]) {
			this.context.beginPath();
			this.context.arc(board.getX(board.starPoints[board.size][key].x)-board.ls, board.getY(board.starPoints[board.size][key].y)-board.ls, theme_variable("starSize", board), 0, 2*Math.PI,true);
			this.context.fill();
		}
	}
}

Board.ShadowLayer = WGo.extendClass(Board.CanvasLayer, function(shadowSize) {
	this.super.call(this);
	this.shadowSize = shadowSize === undefined ? 1 : shadowSize;
});

Board.ShadowLayer.prototype.setDimensions = function(width, height) {
	this.super.prototype.setDimensions.call(this, width, height);
	this.context.setTransform(1,0,0,1,Math.round(this.shadowSize*width/300),Math.round(this.shadowSize*height/300));
}

var default_field_clear = function(args, board) {
	var xr = board.getX(args.x),
		yr = board.getY(args.y),
		sr = board.stoneRadius;
	this.clearRect(xr-sr-board.ls,yr-sr-board.ls, 2*sr, 2*sr);
	this.clearRect(xr-sr-board.ls,yr-sr-board.ls, 2*sr, 2*sr);
}

// Private methods of WGo.Board

var calcLeftMargin = function() {
	return (3*this.width)/(4*(this.bx+1-this.tx)+2) - this.fieldWidth*this.tx;
}

var calcTopMargin = function() {
	return (3*this.height)/(4*(this.by+1-this.ty)+2) - this.fieldHeight*this.ty;
}

var calcFieldWidth = function() {
	return (4*this.width)/(4*(this.bx+1-this.tx)+2);
}

var calcFieldHeight = function() {
	return (4*this.height)/(4*(this.by+1-this.ty)+2);
}

var clearField = function(x,y) {
	var handler;
	for(var key in this.obj_arr[x][y]) {
		if(!this.obj_arr[x][y][key].type) handler = this.stoneHandler;
		else if(typeof this.obj_arr[x][y][key].type == "string") handler = Board.drawHandlers[this.obj_arr[x][y][key].type];
		else handler = this.obj_arr[x][y][key].type;

		for(var layer in handler) {
			if(handler[layer].clear) handler[layer].clear.call(this[layer].context, this.obj_arr[x][y][key], this);
			else default_field_clear.call(this[layer].context, this.obj_arr[x][y][key], this);
		}
	}
}

var drawField = function(x,y) {
	var handler;
	for(var key in this.obj_arr[x][y]) {
		if(!this.obj_arr[x][y][key].type) handler = this.stoneHandler;
		else if(typeof this.obj_arr[x][y][key].type == "string") handler = Board.drawHandlers[this.obj_arr[x][y][key].type];
		else handler = this.obj_arr[x][y][key].type;

		for(var layer in handler) {
			handler[layer].draw.call(this[layer].context, this.obj_arr[x][y][key], this);
		}
	}
}

var getMousePos = function(e) {
    var top = 0,
		left = 0,
		obj = this.grid.element,
		x, y;

	while (obj && obj.tagName != 'BODY') {
        top += obj.offsetTop;
        left += obj.offsetLeft;
        obj = obj.offsetParent;
    }

	x = Math.round((e.pageX-left-this.left)/this.fieldWidth);
	y = Math.round((e.pageY-top-this.top)/this.fieldHeight);

    return {
        x: x >= this.size ? -1 : x,
        y: y >= this.size ? -1 : y
    };
}

var updateDim = function() {
	this.element.style.width = this.width+"px";
	this.element.style.height = this.height+"px";

	this.stoneRadius = theme_variable("stoneSize", this);
	//if(this.autoLineWidth) this.lineWidth = this.stoneRadius/7; //< 15 ? 1 : 3;
	this.ls = theme_variable("linesShift", this);

	for(var key in this.layers) {
		this.layers[key].setDimensions(this.width, this.height);
	}
}

// Public methods are in the prototype:

Board.prototype = {
	constructor: Board,

	/**
     * Initialization method, it is called in constructor. You shouldn't call it, but you can alter it.
	 */

	init: function() {

		// placement of objects (in 3D array)
		this.obj_arr = [];
		for(var i = 0; i < this.size; i++) {
			this.obj_arr[i] = [];
			for(var j = 0; j < this.size; j++) this.obj_arr[i][j] = [];
		}

		// other objects, stored in list
		this.obj_list = [];

		// layers
		this.layers = [];

		// event listeners, binded to board
		this.listeners = [];

		this.element = document.createElement('div');
		this.element.className = 'wgo-board';

		if(this.background) {
			if(this.background[0] == "#") this.element.style.backgroundColor = this.background;
			else {
				this.element.style.backgroundImage = "url('"+this.background+"')";
				/*this.element.style.backgroundRepeat = "repeat";*/
			}
		}

		this.grid = new Board.GridLayer();
		this.shadow = new Board.ShadowLayer(theme_variable("shadowSize", this));
		this.stone = new Board.CanvasLayer();

		this.addLayer(this.grid, 100);
		this.addLayer(this.shadow, 200);
		this.addLayer(this.stone, 300);
	},

	/**
	 * Set new width of board, height is computed to keep aspect ratio.
	 *
	 * @param {number} width
	 */

	setWidth: function(width) {
		this.width = width;
		this.fieldHeight = this.fieldWidth = calcFieldWidth.call(this);
		this.left = calcLeftMargin.call(this);

		this.height = (this.by-this.ty+1.5)*this.fieldHeight;
		this.top = calcTopMargin.call(this);

		updateDim.call(this);
		this.redraw();
	},

	/**
	 * Set new height of board, width is computed to keep aspect ratio.
	 *
	 * @param {number} height
	 */

	setHeight: function(height) {
		this.height = height;
		this.fieldWidth = this.fieldHeight = calcFieldHeight.call(this);
		this.top = calcTopMargin.call(this);

		this.width = (this.bx-this.tx+1.5)*this.fieldWidth;
		this.left = calcLeftMargin.call(this);

		updateDim.call(this);
		this.redraw();
	},

	/**
	 * Set both dimensions.
	 *
	 * @param {number} width
	 * @param {number} height
	 */

	setDimensions: function(width, height) {
		this.width = width || this.width;
		this.height = height || this.height;

		this.fieldWidth = calcFieldWidth.call(this);
		this.fieldHeight = calcFieldHeight.call(this);
		this.left = calcLeftMargin.call(this);
		this.top = calcTopMargin.call(this);

		updateDim.call(this);
		this.redraw();
	},

	/**
	 * Get currently visible section of the board
	 */

	getSection: function() {
		return this.section;
	},

	/**
	 * Set section of the board to be displayed
	 */

	setSection: function(section_or_top, right, bottom, left) {
		if(typeof section_or_top == "object") {
			this.section = section_or_top;
		}
		else {
			this.section = {
				top: section_or_top,
				right: right,
				bottom: bottom,
				left: left,
			}
		}

		this.tx = this.section.left;
		this.ty = this.section.top;
		this.bx = this.size-1-this.section.right;
		this.by = this.size-1-this.section.bottom;

		this.setDimensions();
	},

	/**
	 * Set board size (eg: 9, 13 or 19), this will clear board's objects.
	 */

	setSize: function(size) {
		var size = size || 19;

		if(size != this.size) {
			this.size = size;

			this.obj_arr = [];
			for(var i = 0; i < this.size; i++) {
				this.obj_arr[i] = [];
				for(var j = 0; j < this.size; j++) this.obj_arr[i][j] = [];
			}

			this.bx = this.size-1-this.section.right;
			this.by = this.size-1-this.section.bottom;
			this.setDimensions();
		}
	},

	/**
	 * Redraw everything.
	 */

	redraw: function() {
		// redraw layers
		for(var i = 0; i < this.layers.length; i++) {
			this.layers[i].clear(this);
			this.layers[i].draw(this);
		}

		// redraw field objects
		for(var i = 0; i < this.size; i++) {
			for(var j = 0; j < this.size; j++) {
				drawField.call(this, i, j);
			}
		}

		// redraw custom objects
		for(var key in this.obj_list) {
			var handler = this.obj_list[key].handler;

			for(var layer in handler) {
				handler[layer].draw.call(this[layer].context, this.obj_list[key].args, this);
			}
		}
	},

	/**
	 * Get absolute X coordinate
	 *
	 * @param {number} x relative coordinate
	 */

	getX: function(x) {
		return this.left+x*this.fieldWidth;
	},

	/**
	 * Get absolute Y coordinate
	 *
	 * @param {number} y relative coordinate
	 */

	getY: function(y) {
		return this.top+y*this.fieldHeight;
	},

	/**
	 * Add layer to the board. It is meant to be only for canvas layers.
	 *
	 * @param {Board.CanvasLayer} layer to add
	 * @param {number} weight layer with biggest weight is on the top
	 */

	addLayer: function(layer, weight) {
		layer.element.style.position = 'absolute';
		layer.element.style.zIndex = weight;
		layer.setDimensions(this.width, this.height);
		this.element.appendChild(layer.element);
		this.layers.push(layer);
	},

	/**
	 * Remove layer from the board.
	 *
	 * @param {Board.CanvasLayer} layer to remove
	 */

	removeLayer: function(layer) {
		var i = this.layers.indexOf(layer);
		if(i >= 0) {
			this.layers.splice(i,1);
			this.element.removeChild(layer.element);
		}
	},

	update: function(changes) {
		if(changes.remove && changes.remove == "all") this.removeAllObjects();
		else if(changes.remove) {
			for(var key in changes.remove) this.removeObject(changes.remove[key]);
		}

		if(changes.add) {
			for(var key in changes.add) this.addObject(changes.add[key]);
		}
	},

	addObject: function(obj) {
		// handling multiple objects
		if(obj.constructor == Array) {
			for(var key in obj) this.addObject(obj[key]);
			return;
		}

		// clear all objects on object's coordinates
		clearField.call(this, obj.x, obj.y);

		// if object of this type is on the board, replace it
		for(var key in this.obj_arr[obj.x][obj.y]) {
			if(this.obj_arr[obj.x][obj.y][key].type == obj.type) {
				this.obj_arr[obj.x][obj.y][key] = obj;
				drawField.call(this, obj.x, obj.y);
				return;
			}
		}

		// if object is a stone, add it at the beginning, otherwise at the end
		if(!obj.type) this.obj_arr[obj.x][obj.y].unshift(obj);
		else this.obj_arr[obj.x][obj.y].push(obj);

		// draw all objects
		drawField.call(this, obj.x, obj.y);
	},

	removeObject: function(obj) {
		// handling multiple objects
		if(obj.constructor == Array) {
			for(var key in obj) this.removeObject(obj[key]);
			return;
		}

		var i;
		for(var j = 0; j < this.obj_arr[obj.x][obj.y].length; j++) {
			if(this.obj_arr[obj.x][obj.y][j].type == obj.type) {
				i = j;
				break;
			}
		}
		if(i === undefined) return;

		// clear all objects on object's coordinates
		clearField.call(this, obj.x, obj.y);

		this.obj_arr[obj.x][obj.y].splice(i,1);

		drawField.call(this, obj.x, obj.y);
	},

	removeObjectsAt: function(x, y) {
		if(!this.obj_arr[x][y].length) return;

		clearField.call(this, x, y);
		this.obj_arr[x][y] = [];
	},

	removeAllObjects: function() {
		this.obj_arr = [];
		for(var i = 0; i < this.size; i++) {
			this.obj_arr[i] = [];
			for(var j = 0; j < this.size; j++) this.obj_arr[i][j] = [];
		}
		this.redraw();
	},

	addCustomObject: function(handler, args) {
		this.obj_list.push({handler: handler, args: args});
		this.redraw();
	},

	removeCustomObject: function(handler, args) {
		for(var key in this.obj_list) {
			if(this.obj_list[key].handler == handler && this.obj_list[key].args == args) {
				delete this.obj_list[key];
				this.redraw();
				return true;
			}
		}
		return false;
	},

	addEventListener: function(type, callback) {
		var _this = this,
			evListener = {
				type: type,
				callback: callback,
				handleEvent: function(e) {
					var coo = getMousePos.call(_this, e);
					callback(coo.x, coo.y, e);
				}
			};

		this.element.addEventListener(type, evListener, true);
		this.listeners.push(evListener);
	},

	removeEventListener: function(type, callback) {
		for(var key in this.listeners) {
			if(this.listeners[key].type == type && this.listeners[key].callback == callback) {
				this.element.removeEventListener(this.listeners[key].type, this.listeners[key], true);
				delete this.listeners[key];
				return true;
			}
		}
		return false;
	},

	getState: function() {
		return {
			objects: WGo.clone(this.obj_arr),
			custom: WGo.clone(this.obj_list)
		};
	},

	restoreState: function(state) {
		this.obj_arr = state.objects || this.obj_arr;
		this.obj_list = state.custom || this.obj_list;

		this.redraw();
	}
}

Board.default = {
	size: 19,
	width: 0,
	height: 0,
	font: "Calibri", // deprecated
	lineWidth: 1, // deprecated
	autoLineWidth: false, // deprecated
	starPoints: {
		19:[{x:3, y:3 },
			{x:9, y:3 },
			{x:15,y:3 },
			{x:3, y:9 },
			{x:9, y:9 },
			{x:15,y:9 },
			{x:3, y:15},
			{x:9, y:15},
			{x:15,y:15}],
		13:[{x:3, y:3},
			{x:9, y:3},
			{x:3, y:9},
			{x:9, y:9}],
		9:[{x:4, y:4}],
	},
	stoneHandler: Board.drawHandlers.SHELL,
	starSize: 1, // deprecated
	shadowSize: 1, // deprecated
	stoneSize: 1, // deprecated
	section: {
		top: 0,
		right: 0,
		bottom: 0,
		left: 0,
	},
	background: WGo.DIR+"wood1.jpg",
	theme: {}
}

// save Board
WGo.Board = Board;

//-------- WGo.Game ---------------------------------------------------------------------------

/**
 * Creates instance of position object.
 *
 * @class
 * <p>WGo.Position is simple object storing position of go game. It is implemented as matrix <em>size</em> x <em>size</em> with values WGo.BLACK, WGo.WHITE or 0. It can be used by any extension.</p>
 *
 * @param {number} size of the board
 */

var Position = function(size) {
	this.size = size || 19;
	this.schema = [];
	for(var i = 0; i < this.size*this.size; i++) {
		this.schema[i] = 0;
	}
}

Position.prototype = {
	constructor: WGo.Position,

	/**
	 * Returns value of given coordinates.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @return {(WGo.BLACK|WGo.WHITE|0)} color
	 */

	get: function(x,y) {
		if(x < 0 || y < 0 || x >= this.size || y >= this.size) return undefined;
		return this.schema[x*this.size+y];
	},

	/**
	 * Sets value of given coordinates.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @param {(WGo.B|WGo.W|0)} c color
	 */

	set: function(x,y,c) {
		this.schema[x*this.size+y] = c;
		return this;
	},

	/**
	 * Clears the whole position (every value is set to 0).
	 */

	clear: function() {
		for(var i = 0; i < this.size*this.size; i++) this.schema[i] = 0;
		return this;
	},

	/**
	 * Clones the whole position.
	 *
	 * @return {WGo.Position} copy of position
	 */

	clone: function() {
		var clone = new Position(this.size);
		clone.schema = this.schema.slice(0);
		return clone;
	},

	/**
	 * Compares this position with another position and return change object
	 *
	 * @param {WGo.Position} position to compare to.
	 * @return {object} change object with structure: {add:[], remove:[]}
	 */

	compare: function(position) {
		var add = [], remove = [];

		for(var i = 0; i < this.size*this.size; i++) {
			if(this.schema[i] && !position.schema[i]) remove.push({
				x: Math.floor(i/this.size),
				y: i%this.size
			});
			else if(this.schema[i] != position.schema[i]) add.push({
				x: Math.floor(i/this.size),
				y: i%this.size,
				c: position.schema[i]
			});
		}

		return {
			add: add,
			remove: remove
		}
	}
}

WGo.Position = Position;

/**
 * Creates instance of game class.
 *
 * @class
 * This class implements game logic. It basically analyses given moves and returns capture stones.
 * WGo.Game also stores every position from beginning, so it has ability to check repeating positions
 * and it can effectively restore old positions.</p>
 *
 * @param {number} size of the board
 * @param {"KO"|"ALL"|"NONE"} checkRepeat (optional, default is "KO") - how to handle repeated position:
 * KO - ko is properly handled - position cannot be same like previous position
 * ALL - position cannot be same like any previous position - e.g. it forbids triple ko
 * NONE - position can be repeated
 *
 * @param {boolean} allowRewrite (optional, default is false) - allow to play moves, which were already played:
 * @param {boolean} allowSuicide (optional, default is false) - allow to play suicides, stones are immediately captured
 */

var Game = function(size, checkRepeat, allowRewrite, allowSuicide) {
	this.size = size || 19;
	this.repeating = checkRepeat === undefined ? "KO" : checkRepeat; // possible values: KO, ALL or nothing
	this.allow_rewrite = allowRewrite || false;
	this.allow_suicide = allowSuicide || false;

	this.stack = [];
	this.stack[0] = new Position(this.size);
	this.stack[0].capCount = {black:0, white:0};
	this.turn = WGo.B;

	Object.defineProperty(this, "position", {
		get : function(){ return this.stack[this.stack.length-1]; },
		set : function(pos){ this[this.stack.length-1] = pos; }
	});
}

// function for stone capturing
var do_capture = function(position, captured, x, y, c) {
	if(x >= 0 && x < position.size && y >= 0 && y < position.size && position.get(x,y) == c) {
		position.set(x,y,0);
		captured.push({x:x, y:y});

		do_capture(position, captured, x, y-1, c);
		do_capture(position, captured, x, y+1, c);
		do_capture(position, captured, x-1, y, c);
		do_capture(position, captured, x+1, y, c);
	}
}

// looking at liberties
var check_liberties = function(position, testing, x, y, c) {
	// out of the board there aren't liberties
	if(x < 0 || x >= position.size || y < 0 || y >= position.size) return true;
	// however empty field means liberty
	if(position.get(x,y) == 0) return false;
	// already tested field or stone of enemy isn't giving us a liberty.
	if(testing.get(x,y) == true || position.get(x,y) == -c) return true;

	// set this field as tested
	testing.set(x,y,true);

	// in this case we are checking our stone, if we get 4 trues, it has no liberty
	return 	check_liberties(position, testing, x, y-1, c) &&
			check_liberties(position, testing, x, y+1, c) &&
			check_liberties(position, testing, x-1, y, c) &&
			check_liberties(position, testing, x+1, y, c);
}

// analysing function - modifies original position, if there are some capturing, and returns array of captured stones
var check_capturing = function(position, x, y, c) {
	var captured = [];
	// is there a stone possible to capture?
	if(x >= 0 && x < position.size && y >= 0 && y < position.size && position.get(x,y) == c) {
		// create testing map
		var testing = new Position(position.size);
		// if it has zero liberties capture it
		if(check_liberties(position, testing, x, y, c)) {
			// capture stones from game
			do_capture(position, captured, x, y, c);
		}
	}
	return captured;
}

// analysing history
var checkHistory = function(position, x, y) {
	var flag, stop;

	if(this.repeating == "KO" && this.stack.length-2 >= 0) stop = this.stack.length-2;
	else if(this.repeating == "ALL") stop = 0;
	else return true;

	for(var i = this.stack.length-2; i >= stop; i--) {
		if(this.stack[i].get(x,y) == position.get(x,y)) {
			flag = true;
			for(var j = 0; j < this.size*this.size; j++) {
				if(this.stack[i].schema[j] != position.schema[j]) {
					flag = false;
					break;
				}
			}
			if(flag) return false;
		}
	}

	return true;
}

Game.prototype = {

	constructor: Game,

	/**
	 * Gets actual position.
	 *
	 * @return {WGo.Position} actual position
	 */

	getPosition: function() {
		return this.stack[this.stack.length-1];
	},

	/**
	 * Play move.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @param {(WGo.B|WGo.W)} c color
	 * @param {boolean} noplay - if true, move isn't played. Used by WGo.Game.isValid.
	 * @return {number} code of error, if move isn't valid. If it is valid, function returns array of captured stones.
	 *
	 * Error codes:
	 * 1 - given coordinates are not on board
	 * 2 - on given coordinates already is a stone
	 * 3 - suicide (currently they are forbbiden)
	 * 4 - repeated position
	 */

	play: function(x,y,c,noplay) {
		//check coordinates validity
		if(!this.isOnBoard(x,y)) return 1;
		if(!this.allow_rewrite && this.position.get(x,y) != 0) return 2;

		// clone position
		if(!c) c = this.turn;

		var new_pos = this.position.clone();
		new_pos.set(x,y,c);

		// check capturing
		var cap_color = c;
		var captured = check_capturing(new_pos, x-1, y, -c).concat(check_capturing(new_pos, x+1, y, -c), check_capturing(new_pos, x, y-1, -c), check_capturing(new_pos, x, y+1, -c));

		// check suicide
		if(!captured.length) {
			var testing = new Position(this.size);
			if(check_liberties(new_pos, testing, x, y, c)) {
				if(this.allow_suicide) {
					cap_color = -c;
					do_capture(new_pos, captured, x, y, c);
				}
				else return 3;
			}
		}

		// check history
		if(this.repeating && !checkHistory.call(this, new_pos, x, y)) {
			return 4;
		}

		if(noplay) return false;

		// update position info
		new_pos.color = c;
		new_pos.capCount = {
			black: this.position.capCount.black,
			white: this.position.capCount.white
		};
		if(cap_color == WGo.B) new_pos.capCount.black += captured.length;
		else new_pos.capCount.white += captured.length;

		// save position
		this.pushPosition(new_pos);

		// reverse turn
		this.turn = -c;

		return captured;

	},

	/**
	 * Play pass.
	 *
	 * @param {(WGo.B|WGo.W)} c color
	 */

	pass: function(c) {
		this.pushPosition();
		if(c) {
			this.position.color = c;
			this.turn = -c;
		}
		else {
			this.position.color = this.turn;
			this.turn = -this.turn;
		}
	},

	/**
	 * Finds out validity of the move.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @param {(WGo.B|WGo.W)} c color
	 * @return {boolean} true if move can be played.
	 */

	isValid: function(x,y,c) {
		return typeof this.play(x,y,c,true) != "number";
	},

	/**
	 * Controls position of the move.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @return {boolean} true if move is on board.
	 */

	isOnBoard: function(x,y) {
		return x >= 0 && y >= 0 && x < this.size && y < this.size;
	},

	/**
	 * Inserts move into current position. Use for setting position, for example in handicap game. Field must be empty.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @param {(WGo.B|WGo.W)} c color
	 * @return {boolean} true if operation is successfull.
	 */

	addStone: function(x,y,c) {
		if(this.isOnBoard(x,y) && this.position.get(x,y) == 0) {
			this.position.set(x,y,c || 0);
			return true;
		}
		return false;
	},

	/**
	 * Removes move from current position.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @return {boolean} true if operation is successfull.
	 */

	removeStone: function(x,y) {
		if(this.isOnBoard(x,y) && this.position.get(x,y) != 0) {
			this.position.set(x,y,0);
			return true;
		}
		return false;
	},

	/**
	 * Set or insert move of current position.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @param {(WGo.B|WGo.W)} c color
	 * @return {boolean} true if operation is successfull.
	 */

	setStone: function(x,y,c) {
		if(this.isOnBoard(x,y)) {
			this.position.set(x,y,c || 0);
			return true;
		}
		return false;
	},

	/**
	 * Get stone on given position.
	 *
	 * @param {number} x coordinate
	 * @param {number} y coordinate
	 * @return {(WGo.B|WGo.W|0)} color
	 */

	getStone: function(x,y) {
		if(this.isOnBoard(x,y)) {
			return this.position.get(x,y);
		}
		return 0;
	},

	/**
	 * Add position to stack. If position isn't specified current position is cloned and stacked.
	 * Pointer of actual position is moved to the new position.
	 *
	 * @param {WGo.Position} tmp position (optional)
	 */

	pushPosition: function(pos) {
		if(!pos) {
			var pos = this.position.clone();
			pos.capCount = {
				black: this.position.capCount.black,
				white: this.position.capCount.white
			};
			pos.color = this.position.color;
		}
		this.stack.push(pos);
		if(pos.color) this.turn = -pos.color;
		return this;
	},

	/**
	 * Remove current position from stack. Pointer of actual position is moved to the previous position.
	 */

	popPosition: function() {
		var old = null;
		if(this.stack.length > 0) {
			old = this.stack.pop();

			if(this.stack.length == 0) this.turn = WGo.B;
			else if(this.position.color) this.turn = -this.position.color;
			else this.turn = -this.turn;
		}
		return old;
	},

	/**
	 * Removes all positions.
	 */

	firstPosition: function() {
		this.stack = [];
		this.stack[0] = new Position(this.size);
		this.stack[0].capCount = {black:0, white:0};
		this.turn = WGo.B;
		return this;
	},

	/**
	 * Gets count of captured stones.
	 *
	 * @param {(WGo.BLACK|WGo.WHITE)} color
	 * @return {number} count
	 */

	getCaptureCount: function(color) {
		return color == WGo.B ? this.position.capCount.black : this.position.capCount.white;
	},

	/**
	 * Validate postion. Position is tested from 0:0 to size:size, if there are some moves, that should be captured, they will be removed.
	 * You can use this, after insertion of more stones.
	 *
	 * @return array removed stones
	 */

	validatePosition: function() {
		var c, p,
		    white = 0,
			black = 0,
		    captured = [],
		    new_pos = this.position.clone();

		for(var x = 0; x < this.size; x++) {
			for(var y = 0; y < this.size; y++) {
				c = this.position.get(x,y);
				if(c) {
					p = captured.length;
					captured = captured.concat(check_capturing(new_pos, x-1, y, -c),
											   check_capturing(new_pos, x+1, y, -c),
											   check_capturing(new_pos, x, y-1, -c),
											   check_capturing(new_pos, x, y+1, -c));

					if(c == WGo.B) black += captured-p;
					else white += captured-p;
				}
			}
		}
		this.position.capCount.black += black;
		this.position.capCount.white += white;
		this.position.schema = new_pos.schema;

		return captured;
	},
};

// save Game
WGo.Game = Game;

// register WGo
window.WGo = WGo;

})(window);
