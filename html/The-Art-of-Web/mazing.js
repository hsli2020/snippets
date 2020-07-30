function Position(x, y) {
    this.x = x;
    this.y = y;
}
Position.prototype.toString = function() {
    return this.x + ":" + this.y;
};

function Mazing() {
    this.mazeContainer = document.getElementById("maze");
    this.mazeOutputDiv = document.getElementById("maze_output");
    this.mazeMessage = document.getElementById("maze_message");
    this.mazeScore = document.getElementById("maze_score");
    this.heroScore = this.mazeContainer.getAttribute("data-steps") - 2;
    this.maze = [];
    this.heroPos = {};
    this.heroHasKey = false;
    this.childMode = false;
    for (i = 0; i < this.mazeContainer.children.length; i++) {
        for (j = 0; j < this.mazeContainer.children[i].children.length; j++) {
            var el = this.mazeContainer.children[i].children[j];
            this.maze[new Position(i, j)] = el;
            if (el.classList.contains("entrance")) {
                this.heroPos = new Position(i, j);
                this.maze[this.heroPos].classList.add("hero");
            }
        }
    }
    this.mazeOutputDiv.style.width = this.mazeContainer.scrollWidth + "px";
    this.setMessage("first find the key");
    this.keyPressHandler = this.mazeKeyPressHandler.bind(this);
    document.addEventListener("keydown", this.keyPressHandler, false);
}
Mazing.prototype.setMessage = function(text) {
    this.mazeMessage.innerHTML = text;
    this.mazeScore.innerHTML = this.heroScore;
};
Mazing.prototype.heroTakeTreasure = function() {
    this.maze[this.heroPos].classList.remove("nubbin");
    this.heroScore += 10;
    this.setMessage("yay, treasure!");
};
Mazing.prototype.heroTakeKey = function() {
    this.maze[this.heroPos].classList.remove("key");
    this.heroHasKey = true;
    this.heroScore += 20;
    this.mazeScore.classList.add("has-key");
    this.setMessage("you have the key!");
};
Mazing.prototype.gameOver = function(text) {
    document.removeEventListener("keydown", this.keyPressHandler, false);
    this.setMessage(text);
    this.mazeContainer.classList.add("finished");
};
Mazing.prototype.heroWins = function() {
    this.mazeScore.classList.remove("has-key");
    this.maze[this.heroPos].classList.remove("door");
    this.heroScore += 50;
    this.gameOver("you finished !!!");
};
Mazing.prototype.tryMoveHero = function(pos) {
    var nextStep = this.maze[pos].className;
    if (nextStep.match(/sentinel/)) {
        this.heroScore = Math.max(this.heroScore - 5, 0);
        if (!this.childMode && this.heroScore <= 0) {
            this.gameOver("sorry, you didn't make it");
        } else {
            this.setMessage("ow, that hurt!");
        }
        return;
    }
    if (nextStep.match(/wall/)) {
        return;
    }
    if (nextStep.match(/exit/)) {
        if (this.heroHasKey) {
            this.heroWins();
        } else {
            this.setMessage("you need a key to unlock the door");
            return;
        }
    }
    this.maze[this.heroPos].classList.remove("hero");
    this.maze[pos].classList.add("hero");
    this.heroPos = pos;
    if (nextStep.match(/nubbin/)) {
        this.heroTakeTreasure();
        return;
    }
    if (nextStep.match(/key/)) {
        this.heroTakeKey();
        return;
    }
    if (nextStep.match(/exit/)) {
        return;
    }
    if (this.heroScore >= 1) {
        if (!this.childMode) {
            this.heroScore--;
        }
        if (!this.childMode && this.heroScore <= 0) {
            this.gameOver("sorry, you didn't make it");
        } else {
            this.setMessage("...");
        }
    }
};
Mazing.prototype.mazeKeyPressHandler = function(e) {
    var tryPos = new Position(this.heroPos.x, this.heroPos.y);
    switch (e.keyCode) {
        case 37:
            this.mazeContainer.classList.remove("face-right");
            tryPos.y--;
            break;
        case 38:
            tryPos.x--;
            break;
        case 39:
            this.mazeContainer.classList.add("face-right");
            tryPos.y++;
            break;
        case 40:
            tryPos.x++;
            break;
        default:
            return;
    }
    this.tryMoveHero(tryPos);
    e.preventDefault();
};
Mazing.prototype.setChildMode = function() {
    this.childMode = true;
    this.heroScore = 0;
    this.setMessage("collect all the treasure");
};