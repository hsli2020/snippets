// Modal.js
import React from 'react';
import './Modal.css';

const modal = (props) => {
    return (
      <div>
        <div className="modal-wrapper"
          style={{
              transform: props.show ? 'translateY(0vh)' : 'translateY(-100vh)',
              opacity: props.show ? '1' : '0'
          }}>
          <div className="modal-header">
              <h3>Modal Header</h3>
              <span className="close-modal-btn" onClick={props.close}>×</span>
          </div>
          <div className="modal-body">
              <p>{props.children}</p>
          </div>
          <div className="modal-footer">
              <button className="btn-cancel" onClick={props.close}>CLOSE</button>
              <button className="btn-continue">CONTINUE</button>
          </div>
        </div>
      </div>
    )
}

export default modal;

// App.js
import React, { Component } from 'react';
import Modal from './components/Modal/Modal';

class App extends Component {
    constructor() {
        super();
        this.state = { isShowing: false }
    }

    openModalHandler = ()  => { this.setState({ isShowing: true }); }
    closeModalHandler = () => { this.setState({ isShowing: false }); }

    render () {
      return (
        <div>
          { this.state.isShowing ? <div onClick={this.closeModalHandler} className="back-drop"></div> : null }

          <button className="open-modal-btn" onClick={this.openModalHandler}>Open Modal</button>

          <Modal className="modal" show={this.state.isShowing} close={this.closeModalHandler}>
            Maybe aircrafts fly very high because they don't want to be seen in plane sight?
          </Modal>
        </div>
      );
    }
}

export default App;
