// Class Component
class App extends Component {
  state = {
    count: 0
  }

  componentDidMount() {
    console.log('Did mount!');
  }

  increaseCount = () => {
    this.setState({ count: this.state.count + 1 });
  }

  decreaseCount = () => {
    this.setState({ count: this.state.count - 1 });
  }

  render() {
    return (
      <>
        <h1>Counter</h1>
        <div>Current count: {this.state.count}</div>
        <p>
          <button onClick={this.increaseCount}>Increase</button>
          <button onClick={this.decreaseCount}>Decrease</button>
        </p>
      </>
    );
  }
}

// Function Component
function App() {
  const [ count, setCount ] = useState(0);
  const increaseCount = () => setCount(count + 1);
  const decreaseCount = () => setCount(count - 1);

  useEffect(() => {
    console.log('Did mount!');
  }, []);

  return (
    <>
      <h1>Counter</h1>
      <div>Current count: {count}</div>
      <p>
        <button onClick={increaseCount}>Increase</button>
        <button onClick={decreaseCount}>Decrease</button>
      </p>
    </>
  );
}
