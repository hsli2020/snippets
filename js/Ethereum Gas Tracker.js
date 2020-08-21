// https://medium.com/scrappy-squirrels/building-an-ethereum-gas-tracker-e7cd6fd5b691
// Tutorial: Building an Ethereum Gas Tracker

import './App.css';
import { useEffect, useState } from 'react';
import { createAlchemyWeb3 } from '@alch/alchemy-web3';

const NUM_BLOCKS = 20;

function App() {

  const [blockHistory, setBlockHistory] = useState(null);
  const [avgGas, setAvgGas] = useState(null);
  const [avgBlockVolume, setAvgBlockVolume] = useState(null);

  const formatOutput = (data) => {

    let avgGasFee = 0;
    let avgFill = 0;
    let blocks = [];

    for (let i = 0; i < NUM_BLOCKS; i++) {
      avgGasFee = avgGasFee + Number(data.reward[i][1]) + Number(data.baseFeePerGas[i])
      avgFill = avgFill + Math.round(data.gasUsedRatio[i] * 100);

      blocks.push({
        blockNumber: Number(data.oldestBlock) + i,
        reward: data.reward[i].map(r => Math.round(Number(r) / 10 ** 9)),
        baseFeePerGas: Math.round(Number(data.baseFeePerGas[i]) / 10 ** 9),
        gasUsedRatio: Math.round(data.gasUsedRatio[i] * 100),
      })
    }

    avgGasFee = avgGasFee / NUM_BLOCKS;
    avgGasFee = Math.round(avgGasFee / 10 ** 9)

    avgFill = avgFill / NUM_BLOCKS;

    return [blocks, avgGasFee, avgFill];
  }

  useEffect(() => {

    const web3 = createAlchemyWeb3(
      "wss://eth-mainnet.alchemyapi.io/v2/<--API KEY-->",
    );

    let subscription = web3.eth.subscribe('newBlockHeaders');

    subscription.on('data', () => {
      web3.eth.getFeeHistory(NUM_BLOCKS, "latest", [25, 50, 75]).then((feeHistory) => {
        const [blocks, avgGasFee, avgFill] = formatOutput(feeHistory, NUM_BLOCKS);
        setBlockHistory(blocks);
        setAvgGas(avgGasFee);
        setAvgBlockVolume(avgFill);
      });
    });

    return () => {
      web3.eth.clearSubscriptions();
    }
  }, [])

  return (
    <div className='main-container'>
      <h1>EIP-1559 Gas Tracker</h1>
      {!blockHistory && <p>Data is loading...</p>}
      {avgGas && avgBlockVolume && <h3>
        <span className='gas'>{avgGas} Gwei</span> |
        <span className='vol'>{avgBlockVolume}% Volume</span>
      </h3>}
      {blockHistory && <table>
        <thead>
          <tr>
            <th>Block Number</th>
            <th>Base Fee</th>
            <th>Reward (25%)</th>
            <th>Reward (50%)</th>
            <th>Reward (75%)</th>
            <th>Gas Used</th>
          </tr>
        </thead>
        <tbody>
          {blockHistory.map(block => {
            return (
              <tr key={block.blockNumber}>
                <td>{block.blockNumber}</td>
                <td>{block.baseFeePerGas}</td>
                <td>{block.reward[0]}</td>
                <td>{block.reward[1]}</td>
                <td>{block.reward[2]}</td>
                <td>{block.gasUsedRatio}%</td>
              </tr>
            )
          })}
        </tbody>
      </table>}
    </div>
  );
}

export default App;
