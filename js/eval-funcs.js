function test_a1(app) {
    console.log(app.name + ': test_a1()');
}

function test_b2(app) {
    console.log(app.name + ': test_b2()');
}

function test_c3(app) {
    console.log(app.name + ': test_c3()');
}

function exectests(names) {
	var arr = (typeof names === 'string' ? [names] : names);
	arr.forEach(function(name, idx){ eval('test_' + name + '(app)'); });
}

var app = { name: "TEST EVAL" };

exectests([ 'a1', 'b2', 'c3' ]);
