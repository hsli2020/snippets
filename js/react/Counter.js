import React from 'react';

class Counter extends React.Component {
  state = { count: 0 }  // Not ES6 Syntax

  increment = () => {   // Not ES6 Syntax
    this.setState({
      count: this.state.count + 1
    });
  }

  decrement = () => {   // Not ES6 Syntax
    this.setState({
      count: this.state.count - 1
    });
  }

  render() {
    return (
      <div>
        <h2>Counter</h2>
        <div>
          <button onClick={this.decrement}>-</button>
          <span>{this.state.count}</span>
          <button onClick={this.increment}>+</button>
        </div>
      </div>
    )
  }
}

export default Counter;