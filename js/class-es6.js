class Polygon {
    constructor(height, width) {
        this.height = height;
        this.width = width;
    }

    get area() {
        return this.calcArea();
    }

    calcArea() {
        return this.height * this.width;
    }
}

let square = new Polygon(10, 10);

console.log(square.area);

class Point {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }

    static distance(a, b) {
        let dx = a.x - b.x;
        let dy = a.y - b.y;

        return Math.sqrt(dx*dx + dy*dy);
    }
}

let p1 = new Point(5, 5);
let p2 = new Point(10, 10);

console.log(Point.distance(p1, p2));

class Animal {
    constructor(name) {
        this.name = name;
    }

    speak() {
        console.log(this.name + ' makes a noise.');
    }
}

class Dog extends Animal {
    speak() {
        super.speak();
        console.log(this.name + ' barks.');
    }
}

let dog = new Dog("chocolate");
dog.speak();
