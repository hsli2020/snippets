// ---------------------------------------------------------
const printNumbers = {
  phrase: 'The current value is:',
  numbers: [1, 2, 3, 4],

  loop() {
    this.numbers.forEach(function (number) {
      console.log(this.phrase, number)
    })
  },
}

printNumbers.loop()
 
This will give the following:

undefined 1
undefined 2
undefined 3
undefined 4
// ---------------------------------------------------------
// Use bind to fix the issue:

const printNumbers = {
  phrase: 'The current value is:',
  numbers: [1, 2, 3, 4],

  loop() {
    // Bind the `this` from printNumbers to the inner forEach function
    this.numbers.forEach(
      function (number) {
        console.log(this.phrase, number)
      }.bind(this),
    )
  },
}

printNumbers.loop()

This will give the expected result:

The current value is: 1
The current value is: 2
The current value is: 3
The current value is: 4
// ---------------------------------------------------------
// Use arrow function to fix the issue:

const printNumbers = {
  phrase: 'The current value is:',
  numbers: [1, 2, 3, 4],

  loop() {
    this.numbers.forEach((number) => {
      console.log(this.phrase, number)
    })
  },
}

printNumbers.loop()
 
This will give the expected result:

The current value is: 1
The current value is: 2
The current value is: 3
The current value is: 4
// ---------------------------------------------------------
// Arrow Functions as Object Methods

const printNumbers = {
  phrase: 'The current value is:',
  numbers: [1, 2, 3, 4],

  loop: () => {
    this.numbers.forEach((number) => {
      console.log(this.phrase, number)
    })
  },
}

printNumbers.loop()
 
This will give the following:

Uncaught TypeError: Cannot read property 'forEach' of undefined
// ---------------------------------------------------------
