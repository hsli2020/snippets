import React, { Component } from 'react'
import ReactDOM from 'react-dom'
import ContentEditable from 'react-contenteditable'
import { Table, Button } from 'semantic-ui-react'
import './styles.css'

class App extends Component {
  initialState = {
    store: [
      { id: 1, item: 'silver', price: 15.41 },
      { id: 2, item: 'gold', price: 1284.3 },
      { id: 3, item: 'platinum', price: 834.9 },
    ],
    row: {
      item: '',
      price: '',
    },
  }

  state = this.initialState

  addRow = () => {
    const { store, row } = this.state

    row.id = store.length + 1

    this.setState({
      store: [...store, row],
      row: this.initialState.row,
    })
  }

  deleteRow = id => {
    const { store } = this.state

    this.setState({
      store: store.filter(item => id !== item.id),
    })
  }

  handleContentEditable = event => {
    const { row } = this.state
    const {
      currentTarget: {
        dataset: { column },
      },
      target: { value },
    } = event

    this.setState({ row: { ...row, [column]: value } })
  }

  render() {
    const {
      store,
      row: { item, price },
    } = this.state

    return (
      <div className="App">
        <h1>React Contenteditable</h1>

        <Table celled>
          <Table.Header>
            <Table.Row>
              <Table.HeaderCell>Item</Table.HeaderCell>
              <Table.HeaderCell>Price</Table.HeaderCell>
              <Table.HeaderCell>Action</Table.HeaderCell>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {store.map(row => {
              return (
                <Table.Row key={row.id}>
                  <Table.Cell>{row.item}</Table.Cell>
                  <Table.Cell>{row.price}</Table.Cell>
                  <Table.Cell className="narrow">
                    <Button
                      onClick={() => {
                        this.deleteRow(row.id)
                      }}
                    >
                      Delete
                    </Button>
                  </Table.Cell>
                </Table.Row>
              )
            })}
            <Table.Row>
              <Table.Cell className="narrow">
                <ContentEditable
                  html={item}
                  data-column="item"
                  className="content-editable"
                  onChange={this.handleContentEditable}
                />
              </Table.Cell>
              <Table.Cell className="narrow">
                <ContentEditable
                  html={price}
                  data-column="price"
                  className="content-editable"
                  onChange={this.handleContentEditable}
                />
              </Table.Cell>
              <Table.Cell className="narrow">
                <Button onClick={this.addRow}>Add</Button>
              </Table.Cell>
            </Table.Row>
          </Table.Body>
        </Table>
      </div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('root'))
