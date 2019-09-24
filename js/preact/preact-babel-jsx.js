const { h, Component, render } = preact;
const Markup = preactMarkup;
/**@jsx h*/


// Demo: textarea on the left, <Markup> on the right. Renders on change.
class App extends Component {
	componentDidMount() {
		this.loadExample();
	}

	// loads HTML example from a Gist
	loadExample() {
		fetch('https://gist.githubusercontent.com/developit' +
            '/2c6aa8e40bd3ab251968/raw/c470e7e82fcc3eaeea1e3bc26d0829b7592ebb06/example.html')
			.then( r => r.text() )
			.then( markup => this.setState({ type:'html', markup }) );
	}
	
	// click on the error button to show the error
	showError = () => {
		alert(this.state.error);
	};
	
	// clear error on success (<Markup> fires no event for success)
	componentWillUpdate({ }, { markup, type }) {
		if (markup!==this.lastMarkup || type!==this.lastType) {
			this.state.error = null;
			this.lastMarkup = markup;
			this.lastType = type;
		}
	}

	render({ }, { markup, type, error }) {
		// components to use as custom elements:
		let components = {
			Sidebar,
			Toolbar,
			WidgetA,
			OtherWidget
		};

		return (
			<div id="app">
				<div class="editor">
					<div class="toolbar">
						<button onClick={::this.loadExample}>Load Example Gist</button>
						<select value={type} onChange={this.linkState('type')}>
							<option value="xml">XML (Strict)</option>
							<option value="html">HTML5</option>
						</select>
						{ error ? (
							<button onClick={::this.showError}>Error</button>
						) : null }
					</div>

					<textarea value={markup} onInput={this.linkState('markup')} />
				</div>

				<div class="preview">
					<Markup
						type={type}
						markup={markup}
						components={components}
						onError={this.linkState('error', 'error')} />
				</div>
			</div>
		);
	}
}



// creates a dummy component that just
// shows its name and some debug info:
const createPlaceholderComponent = name => ({ children, ...props }) => (
	<div class="placeholder-component" data-component={name} {...props}>
		<h4>{ name }</h4>
		<pre>{ JSON.stringify(props,null,'  ') }</pre>
		<button onClick={ () => console.log(props) }>Log Props</button>
		{ children.length ? (<div class="inner">{children}</div>) : null }
	</div>
);

const WidgetA = createPlaceholderComponent('WidgetA');
const OtherWidget = createPlaceholderComponent('OtherWidget');


// A <Toolbar /> component, to show that
// Stateless Functional Components work too
const Toolbar = ({ children, ...props }) => (
	<header class="toolbar" {...props}>{ children }</header>
);


// A silly <Sidebar /> component
class Sidebar extends Component {
	componentWillMount() {
		let mounts = (this.state.mounts || 0)+1;
		this.setState({ mounts });
	}
	render({ children, ...props }, { mounts }) {
		return (
			<aside class="sidebar" {...props}>
				<span>Mounts: {mounts}</span>
				{ children }
			</aside>
		);
	}
}


render(<App />, document.body);
