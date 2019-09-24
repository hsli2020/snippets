var dog = {
    color: "Brown",

    eat: function() {
        // 'this' is must
        console.log(this.color + " Dog like eat bone");
    },
    sleep: function() {
        // 'this' is must
        console.log(this.color + " Dog sleep " + this.getHours() + " hours each day");
    },
    getHours: function() {
        return 4;
    }
}

dog.color = "Grey" // this works
dog.eat()
dog.sleep()

var cat = new function() {  // 'new' is must
    var color = "White";

    this.eat = function () {
        // 'this' is not allowed
        console.log(color + " Cat like eat fish");
    }

    this.sleep = function () {
        // 'this' is not allowed
        console.log(color + " Cat sleep " + getHours() + " hours");
    }

    // private function
    function getHours() {
        return 8;
    }
}

cat.color = "Black" // this doesn't work
cat.eat()
cat.sleep()

class Wolf {
    constructor(color) {
        this.color = color;
    }
    eat() {
        console.log(this.color + " Wolf eats any kind of meat");
    }
    sleep() {
        console.log(this.color + " Wolf sleeps " + this.getHours() + " hours");
    }
    getHours() {
        return 2
    }
}

var wolf = new Wolf('Gray');
wolf.color = "Yellow"; // this works
wolf.eat();
wolf.sleep();
