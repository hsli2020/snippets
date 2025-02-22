function Position(x, y) {
    this.x = x;
    this.y = y;
};
Position.prototype.toString = function() {
    return this.x + ":" + this.y;
};

function GraphGame() {
    this.canvas = document.getElementById("graph_canvas");
    this.canvas.width = this.canvas.height = this.canvas.offsetWidth;
    this.ctx = this.canvas.getContext("2d");
    this.gameStatus = document.getElementById("graph_status");
    this.input = document.getElementById("graph_eq");
    this.errorOut = document.getElementById("graph_error");
    this.butGraph = document.getElementById("graph_run");
    this.butReset = document.getElementById("graph_reset");
    this.displayEquation = document.getElementById("graph_output");
    var AudioContext = window.AudioContext || window.webkitAudioContext;
    this.audio = new AudioContext();
    this.butGraph.addEventListener("click", function(e) {
        this.parseEquation(this.input.value);
    }.bind(this), false);
    this.input.addEventListener("keypress", function(e) {
        if (e.key === "Enter") {
            this.parseEquation(this.input.value);
        }
        if (!e.key.match(/[\\d.xy=\\/=\\s()+-]/)) {
            e.preventDefault();
        }
    }.bind(this), false);
    this.butReset.addEventListener("click", function(e) {
        this.reset();
    }.bind(this), false);
    this.minX = -5;
    this.maxX = 5;
    this.minY = -5;
    this.maxY = 5;
    this.stepSize = 0.1;
    this.rangeX = this.maxX - this.minX;
    this.rangeY = this.maxY - this.minY;
    this.gridX = this.canvas.width / (1 + this.rangeX);
    this.gridY = this.canvas.height / (1 + this.rangeY);
    this.center = new Position(Math.round(Math.abs((this.minX - 0.5) / (1 + this.rangeX)) * this.canvas.width), Math.round(Math.abs((this.maxY + 0.5) / (1 + this.rangeY)) * this.canvas.height));
    this.reset();
};
GraphGame.prototype.reset = function() {
    this.targets = new Map();
    this.plots = [];
    this.equations = [];
    this.scores = [];
    this.displayEquation.innerHTML = "plotted equations appear here";
    this.input.disabled = false;
    this.butGraph.style.display = "inline-block";
    this.butReset.style.display = "none";
    this.ctx.fillStyle = "#0c3c00";
    this.ctx.fillRect(0, 0, this.canvas.width, this.canvas.height);
    this.drawAxes();
    this.generateTargets(5);
    this.tries = 5;
    this.displayScore();
    this.input.focus();
};
GraphGame.prototype.drawLine = function(x1, y1, x2, y2) {
    this.ctx.moveTo(x1, y1);
    this.ctx.lineTo(x2, y2);
};
GraphGame.prototype.drawAxes = function() {
    var x, y, i, xPos, yPos;
    var tickSize = 8;
    this.ctx.beginPath();
    for (i = 1; i <= Math.sqrt(Math.pow(this.maxX, 2) + Math.pow(this.maxY, 2)); i++) {
        this.ctx.arc(this.center.x, this.center.y, i * this.gridX, 0, 2 * Math.PI, false);
    }
    for (x = this.minX; x <= this.maxX; x++) {
        var xPos = Math.round(this.center.x + this.gridX * x);
        this.drawLine(xPos, 0, xPos, this.canvas.height);
    }
    for (y = this.minY; y <= this.maxY; y++) {
        yPos = Math.round(this.center.y - this.gridY * y);
        this.drawLine(0, yPos, this.canvas.width, yPos);
    }
    this.ctx.strokeStyle = "rgba(255,255,255,0.15)";
    this.ctx.lineWidth = 1;
    this.ctx.stroke();
    this.ctx.beginPath();
    this.drawLine(0, this.center.y, this.canvas.width, this.center.y);
    this.drawLine(this.center.x, 0, this.center.x, this.canvas.height);
    this.ctx.font = "8pt Times";
    this.ctx.fillStyle = "white";
    this.ctx.textAlign = "center";
    this.ctx.textBaseline = "top";
    for (x = this.minX; x <= this.maxX; x++) {
        if (x == 0) continue;
        xPos = Math.round(this.center.x + this.gridX * x);
        this.drawLine(xPos, this.center.y - tickSize, xPos, this.center.y + tickSize);
        this.ctx.fillText(x, xPos, this.center.y + tickSize + 3);
    }
    this.ctx.textAlign = "right";
    this.ctx.textBaseline = "middle";
    for (y = this.minY; y <= this.maxY; y++) {
        if (y == 0) continue;
        yPos = Math.round(this.center.y - this.gridY * y);
        this.drawLine(this.center.x - tickSize, yPos, this.center.x + tickSize, yPos);
        this.ctx.fillText(y, this.center.x - tickSize - 3, yPos);
    }
    this.ctx.strokeStyle = "white";
    this.ctx.lineWidth = 2;
    this.ctx.stroke();
};
GraphGame.prototype.gridToCanvas = function(pos) {
    return new Position(this.center.x + this.gridX * pos.x, this.center.y - this.gridY * pos.y);
};
GraphGame.prototype.generateTargets = function(count) {
    var x, y, i;
    var targetOptions = [];
    for (x = this.minX; x <= this.maxX; x++) {
        for (y = this.minY; y <= this.maxY; y++) {
            targetOptions.push(new Position(x, y));
        }
    }
    this.targets.clear();
    for (i = 0; i < count; i++) {
        var target = targetOptions.splice(Math.floor(Math.random() * targetOptions.length), 1)[0];
        this.targets.set(target.toString(), target);
    }
    this.targets.forEach(function(pos, key) {
        this.ctx.beginPath();
        pos = this.gridToCanvas(pos);
        this.ctx.arc(pos.x, pos.y, 8, 0, 2 * Math.PI, false);
        this.ctx.arc(pos.x, pos.y, 5, 0, 2 * Math.PI, true);
        this.ctx.fillStyle = "#22ff05";
        this.ctx.fill();
    }, this);
};
GraphGame.prototype.gameOver = function(winner) {
    this.input.disabled = true;
    this.butGraph.style.display = "none";
    this.butReset.style.display = "inline-block";
    this.hum.setVolume(0.1);
    var displayText = "";
    if (winner) {
        displayText = "YOU WIN!";
        (new SoundPlayer(this.audio)).play(587.3, 0.5, "sine").stop(0.25);
        (new SoundPlayer(this.audio)).play(587.3, 0.5, "sine", 0.3).stop(0.35);
        (new SoundPlayer(this.audio)).play(659.3, 0.5, "sine", 0.4).stop(0.55);
        (new SoundPlayer(this.audio)).play(587.3, 0.5, "sine", 0.6).stop(0.75);
        (new SoundPlayer(this.audio)).play(784.0, 0.5, "sine", 0.8).stop(0.95);
        (new SoundPlayer(this.audio)).play(740.0, 0.5, "sine", 1.0).stop(1.40);
    } else {
        displayText = "GAME OVER";
        (new SoundPlayer(this.audio)).play(329.6, 0.5, "sine").stop(0.15);
        (new SoundPlayer(this.audio)).play(329.6, 0.5, "sine", 0.20).stop(0.35);
        (new SoundPlayer(this.audio)).play(329.6, 0.5, "sine", 0.40).stop(0.55);
        (new SoundPlayer(this.audio)).play(261.6, 0.5, "sine", 0.60).stop(1.20);
    }
    this.ctx.save();
    var fontSize = 24;
    do {
        this.ctx.font = fontSize + "px Impact";
        var width = this.ctx.measureText(displayText).width;
        fontSize++;
    } while (width < this.canvas.width * 0.8);
    this.ctx.textBaseline = "middle";
    this.ctx.textAlign = "center";
    this.ctx.fillStyle = "white";
    this.ctx.shadowColor = "#22ff05";
    this.ctx.shadowOffsetX = this.ctx.shadowOffsetX = 0;
    this.ctx.shadowBlur = 4;
    this.ctx.fillText(displayText, this.canvas.width / 2, this.canvas.height / 2);
    this.ctx.restore();
};
GraphGame.prototype.drawHit = function(pos) {
    (new SoundPlayer(this.audio)).play(440, 0.5, "square").setFrequency(880, 0.1).stop(0.2);
    this.ctx.beginPath();
    pos = this.gridToCanvas(pos);
    this.ctx.lineWidth = 0;
    this.ctx.arc(pos.x, pos.y, 8, 0, 2 * Math.PI, false);
    this.ctx.fillStyle = "#22ff05";
    this.ctx.fill();
    this.ctx.moveTo(pos.x, pos.y);
};
GraphGame.prototype.plot = function(plotNo, stepNo, animate) {
    var coords = this.plots[plotNo][stepNo];
    var freq = Math.max(0, 220 + 10 * coords.y);
    var newCoords = this.gridToCanvas(coords);
    if (stepNo == 0) {
        this.hum.play(freq, 0.2, "sine");
        this.ctx.moveTo(newCoords.x, newCoords.y);
        this.input.disabled = true;
    } else {
        this.hum.setFrequency(freq);
        this.ctx.lineTo(newCoords.x, newCoords.y);
    }
    this.ctx.stroke();
    if (this.targets.has(coords.toString())) {
        this.drawHit(coords);
        this.scores[plotNo] = 10 + 2 * this.scores[plotNo];
        this.targets.delete(coords.toString());
        this.displayScore();
        if (this.targets.size == 0) {
            this.gameOver(true);
        }
    }
    if (++stepNo < this.plots[plotNo].length) {
        if (animate) {
            requestAnimationFrame(this.plot.bind(this, plotNo, stepNo, animate));
        } else {
            this.plot(plotNo, stepNo, animate);
        }
    } else {
        this.hum.stop();
        this.tries--;
        this.displayScore();
        if (this.tries == 0) {
            if (this.targets.size > 0) {
                this.gameOver(false);
            }
        } else {
            this.input.disabled = false;
        }
    }
};
GraphGame.prototype.displayScore = function() {
    var total = 0;
    this.scores.forEach(function(s) {
        total += s;
        return total;
    }, this);
    this.gameStatus.children[0].innerHTML = "Score: " + total;
    this.gameStatus.children[1].innerHTML = "&#x29BF;".repeat(this.tries);
    return total;
};
GraphGame.prototype.drawEquation = function(equation, sideways) {
    var x, y;
    var waypoints = [];
    if (sideways) {
        for (y = this.minY - 0.5; y <= this.maxY + 0.5; y += this.stepSize) {
            y = Math.round(y * 10) / 10;
            x = Math.round(equation(y) * 100) / 100;
            waypoints.push(new Position(x, y));
        }
    } else {
        for (x = this.minX - 0.5; x <= this.maxX + 0.5; x += this.stepSize) {
            x = Math.round(x * 10) / 10;
            y = Math.round(equation(x) * 100) / 100;
            waypoints.push(new Position(x, y));
        }
    }
    this.plots.push(waypoints);
    this.ctx.beginPath();
    this.ctx.lineJoin = "round";
    this.ctx.lineWidth = 2;
    this.ctx.strokeStyle = "#22ff05";
    this.scores[this.plots.length - 1] = 0;
    this.hum = new SoundPlayer(this.audio);
    this.plot(this.plots.length - 1, 0, true);
};
GraphGame.prototype.prettify = function(str, variable) {
    if (variable == "y") {
        str = str.replace(/x/g, "y");
    }
    str = str.replace(/\\s*\\+?\\s*\\-\\s*/g, "-", str);
    str = str.replace(/\\s*\\/\\
        s * /g," &divide; ");str=str.replace(/Math.pow\\(([xy]), (\\d + )\\) / g, "$1<sup>$2</sup>");
    str = str.replace(/\\s*([+-])\\s*/g, " $1 ");
    str = str.replace(/(^|\\s)1\\s*\\*\\s*/g, " ");
    str = str.replace(/\\*/g, "•");
    str = str.replace(/(\\d)\\1\\1+\\b/g, function(str, match) {
        return match[0] + match[0] + "<s>" + match[0] + "</s>"
    });
    return str;
};
GraphGame.prototype.parseEquation = function(str) {
    if (matches = str.trim().match(/^(x|y)\\s*=\\s*(.+)$/)) {
        var variable = (matches[1] == "y") ? "x" : "y";
        var equation = matches[2];
        equation = equation.replace(/\\s*/g, "");
        equation = equation.replace(/\\)(x|y|\\d+)/g, ")($1)");
        var re = new RegExp("(-?[\\\\d.]*)(" + variable + "(\\\\d)?)?(\\/\\\\d+)?", "g");
        var replacer = function(match, times, vari, pow, div, offset, string) {
            retval = "";
            if (times) {
                if (times == "-") {
                    times = "-1";
                }
                retval += times;
                if (vari) {
                    retval += " * ";
                }
            }
            if (vari && pow) {
                retval += "Math.pow(x, " + pow + ")";
            } else if (vari) {
                retval += "x";
            }
            if (div) {
                retval += div;
            }
            return retval;
        };
        var eq = equation.replace(re, replacer);
        eq = eq.replace(/([)xy\\d])\\s*([\\(xy])/g, "$1 * $2");
        this.errorOut.innerHTML = "";
        try {
            var x = 0;
            eval(eq);
        } catch (error) {
            this.errorOut.innerHTML = "Error: " + error.message;
            this.input.focus();
            return;
        }
        this.drawEquation(function(x) {
            return eval(eq);
        }, (variable == "y"));
        this.equations.push(matches[1] + " = " + this.prettify(eq, variable));
        this.displayEquations();
        this.input.value = "";
    }
    this.input.focus();
}
GraphGame.prototype.displayEquations = function() {
    this.displayEquation.innerHTML = this.equations.join("<br>");
};