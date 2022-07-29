// https://codesandbox.io/s/xpmnzc

// src/App.js

import React, { Component } from "react";
import Web3 from "web3";
// import SimpleContract from "./Contract.json";
import getWeb3 from "./getWeb3";
import "./App.css";

const contractAddress = "0x2a980B680Aed6BeC1f02Dbe14ddC28Dc9680D73e";
const web3 = new Web3(Web3.givenProvider || "http://localhost/8545");

class App extends Component {
  state = { storageValue: 0, web3: null, accounts: null, contract: null };

  connectMetaMask = async () => {
    let accounts;
    try {
      await web3.givenProvider.request({ method: "eth_requestAccounts" });
      //setMetaMaskObject({ metaMaskConnected: true, metaMaskPresent });
      accounts = await web3.eth.getAccounts();
      const networkId = await web3.eth.net.getId();
      if (networkId !== 4) {
        alert("必須使用 Rinkeby 網絡！");
        await web3.currentProvider.request({
          method: "wallet_switchEthereumChain",
          params: [{ chainId: "0x4" }]
        });
      }
      //setPublicKey(() => accounts[0]);
      this.setState({ web3, accounts }, this.initData);
    } catch (error) {
      console.error("metmask error", error);
    }
  };

  componentDidMount = async () => {
    try {
      //this.connectMetaMask();
    } catch (error) {
      alert(`Failed to load required libraries.`);
      console.error(error);
    }
  };

  signEIP712Data = async () => {
    const { web3, accounts, contract } = this.state;
    var signer = accounts[0];

    await web3.currentProvider.sendAsync(
      {
        method: "net_version",
        params: [],
        jsonrpc: "2.0"
      },
      function (err, result) {
        const netId = result.result;
        console.log("netId", netId);
        const chainId = netId;

        const eip712Obj = {
          types: {
            EIP712Domain: [
              { name: "name", type: "string" },
              { name: "version", type: "string" },
              { name: "chainId", type: "uint256" },
              { name: "verifyingContract", type: "address" }
            ],
            Test: [
              { name: "owner", type: "address" },
              { name: "amount", type: "uint256" },
              { name: "nonce", type: "uint256" }
            ]
          },

          domain: (chainId: number) => ({
            name: "Frank 0x0016",
            version: "1",
            chainId: chainId,
            verifyingContract: contractAddress
          })
        };

        const data = JSON.stringify({
          types: eip712Obj.types,
          domain: eip712Obj.domain(chainId),
          primaryType: "Test",
          message: {
            owner: signer,
            amount: 120,
            nonce: 2374
          }
        });

        console.log("data", data);

        web3.currentProvider
          .request({
            method: "eth_signTypedData_v4",
            params: [signer, data]
          })
          .then((result) => {
            const signature = result.substring(2);
            const r = "0x" + signature.substring(0, 64);
            const s = "0x" + signature.substring(64, 128);
            const v = parseInt(signature.substring(128, 130), 16);
            console.log("Result:", { r, s, v });
          })
          .catch(function (err) {
            console.log(err);
          });
      }
    );
  };

  signPersonalSign = async () => {
    const { web3, accounts } = this.state;
    var signer = accounts[0];

    web3.currentProvider
      .request({
        method: "personal_sign",
        params: [
          signer,
          "Hello, this is some human readable message for you to sign. \n\nThis will not cost you any gas and should not be mark as dangerous by MataMask."
        ]
      })
      .then((result) => {
        console.log(result);
      })
      .catch(function (err) {
        console.log(err);
      });
  };

  signETHSignData = async () => {
    const { web3, accounts } = this.state;
    var signer = accounts[0];

    web3.currentProvider
      .request({
        method: "eth_sign",
        params: [signer, web3.utils.sha3("Hello world")]
      })
      .then((result) => {
        console.log(result);
      })
      .catch(function (err) {
        console.log(err);
      });
  };

  render() {
    if (!this.state.web3) {
      return (
        <center>
          <p>&nbsp;</p>
          這個範例演示 MetaMask 簽名的三種方式。 <br />
          請使用 Rinkeby 測試網路，然後連接 MetaMask
          <p>&nbsp;</p>
          <div>
            <button onClick={() => this.connectMetaMask()}>
              {" "}
              連接 MetaMask{" "}
            </button>
          </div>
        </center>
      );
    }
    return (
      <div className="App">
        <h2>EIP712 簽名</h2>
        <p>請確認目前是在 Rinkeby 測試網路。</p>
        <p>可以在 console 中查看詳細 log。</p>

        <center>
          <div
            style={{
              width: "80%",
              textAlign: "left",
              backgroundColor: "#efefef",
              padding: "30px",
              borderRadius: "15px",
              margin: "20px"
            }}
          >
            <h3>personal_sign Sign</h3>
            <p>
              調用 personal_sign， 在簽名信息前加入 prefix 來避免未授權使用.
              可顯示 UTF-8
              編碼的文字，讓簽名者明確正在簽名什麼內容。這種方式多用在網站登入。
            </p>
            <h3>eth_sign Sign</h3>
            <p>
              eth_sign 必須傳入一個 32 byte 的 message hash 以供簽名，因此，僅憑
              message hash
              簽名者不會知道自己正在簽署什麼內容，可能是一個交易資料，又或者其它任何內容，
              <strong>
                因此有潛在的被釣魚的風險。 MetaMask 會彈出紅色警告。
              </strong>
            </p>
            <h3>EIP712 Sign</h3>
            <p>
              {" "}
              EIP712 Sign 是根據 EIP712 標準為 signature 加入了如
              domain，contract address 等資訊，從而提高簽名的安全性。
            </p>
          </div>
        </center>

        <button onClick={() => this.signPersonalSign()}>
          {" "}
          personal_sign Sign{" "}
        </button>
        <button onClick={() => this.signETHSignData()}> eth_sign Sign </button>
        <button onClick={() => this.signEIP712Data()}> EIP712 Sign </button>
        <p></p>
      </div>
    );
  }
}

export default App;

// src/getWeb3.js

import Web3 from "web3";

const getWeb3 = () =>
  new Promise((resolve, reject) => {

    console.log("getWeb3");
    // Wait for loading completion to avoid race conditions with web3 injection timing.
    window.addEventListener("load", async () => {
      // Modern dapp browsers...
      if (window.ethereum) {
        console.log("true");
        const web3 = new Web3(window.ethereum);
        //const accounts = await window.ethereum.request({ method: "eth_requestAccounts" });
       // console.log(accounts);

        
        try {
          // Request account access if needed
          await window.ethereum.enable();
          // Accounts now exposed
          resolve(web3);
        } catch (error) {
          console.log("error");
          reject(error);
        }
      }
      // Legacy dapp browsers...
      else if (window.web3) {
        // Use Mist/MetaMask's provider.
        const web3 = window.web3;
        console.log("Injected web3 detected.");
        resolve(web3);
      }
      // Fallback to localhost; use dev console port by default...
      else {
        const provider = new Web3.providers.HttpProvider(
          "http://127.0.0.1:8545"
        );
        const web3 = new Web3(provider);
        console.log("No web3 instance injected, using Local web3.");
        resolve(web3);
      }
    });
  });

export default getWeb3;

// src/index.js

import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';


ReactDOM.render(<App />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
//serviceWorker.unregister();

// package.json

{
  "name": "client",
  "version": "0.1.0",
  "private": true,
  "dependencies": {
    "@openzeppelin/contracts": "^4.3.1",
    "eth-sig-util": "^3.0.1",
    "ethereumjs-util": "^6.2.0",
    "react": "16.11.0",
    "react-dom": "16.11.0",
    "react-scripts": "3.2.0",
    "web3": "1.2.2"
  },
  "scripts": {
    "start": "react-scripts start",
    "build": "react-scripts build",
    "test": "react-scripts test",
    "eject": "react-scripts eject"
  },
  "eslintConfig": {
    "extends": "react-app"
  },
  "browserslist": {
    "production": [
      ">0.2%",
      "not dead",
      "not op_mini all"
    ],
    "development": [
      "last 1 chrome version",
      "last 1 firefox version",
      "last 1 safari version"
    ]
  }
}
