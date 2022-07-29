/* tokeninsight.com - https://github.com/tokenInsight/coin-tools */

// coin-tools/contracts/src/BatchTransfer.sol
pragma solidity ^0.8.0; //SPDX-License-Identifier: MIT

// Simplified IERC20 interface containing only the required functions for this contract
interface IERC20 {
    function transfer(address recipient, uint256 amount) external;
    function transferFrom(address sender, address recipient, uint256 amount) external;
}

contract BatchTransfer {
    // Batch transfer Ether
    function batchTransferEther(
        address payable[] calldata recipients, uint256[] calldata amounts) external payable {
        require(recipients.length == amounts.length, 
                "Recipients and amounts arrays must have the same length");

        for (uint256 i = 0; i < recipients.length; i++) {
            recipients[i].transfer(amounts[i]);
        }

        uint256 remainingBalance = address(this).balance;
        if (remainingBalance > 0) {
            payable(msg.sender).transfer(remainingBalance);
        }
    }

    // Batch transfer ERC20 tokens
    function batchTransferToken(
        IERC20 token, address[] calldata recipients, uint256[] calldata amounts) external {
        require(recipients.length == amounts.length, 
                "Recipients and amounts arrays must have the same length");

        uint256 totalTokens = 0;
        for (uint256 i = 0; i < recipients.length; i++) {
            totalTokens += amounts[i];
        }

        token.transferFrom(msg.sender, address(this), totalTokens);

        for (uint256 i = 0; i < recipients.length; i++) {
            token.transfer(recipients[i], amounts[i]);
        }
    }

    function batchTransferTokenSimple(
        IERC20 token, address[] calldata recipients, uint256[] calldata amounts)
        external
    {
        require(recipients.length == amounts.length, 
                "Recipients and amounts arrays must have the same length");
        for (uint256 i = 0; i < recipients.length; i++) {
            token.transferFrom(msg.sender, recipients[i], amounts[i]);
        }
    }
}

// coin-tools/contracts/src/BatchTransferCall.sol
pragma solidity ^0.8.0; //SPDX-License-Identifier: MIT

// Simplified IERC20 interface containing only the required functions for this contract
interface IERC20 {
    function transfer(address recipient, uint256 amount) external;
    function transferFrom(address sender, address recipient, uint256 amount) external;
}

contract BatchTransferCall {
    uint256 private _inProgress = 1;
    // simple modifier to prevent reentrancy

    modifier preventReentrancy() {
        require(_inProgress != 2, "Batch transfer already in progress");
        _inProgress = 2;
        _;
        _inProgress = 1;
    }

    function batchTransferEther(address payable[] calldata recipients, uint256[] calldata amounts)
        external payable preventReentrancy
    {
        require(recipients.length == amounts.length, 
                "Recipients and amounts arrays must have the same length");

        for (uint256 i = 0; i < recipients.length; i++) {
            (bool succ,) = recipients[i].call{value: amounts[i]}("");
            require(succ, "Failed to send Ether");
        }

        uint256 remainingBalance = address(this).balance;
        require(remainingBalance == 0, 
                "Total transfer amount and individual transactions do not match");
    }

    // Batch transfer ERC20 tokens
    function batchTransferToken(
        IERC20 token, address[] calldata recipients, uint256[] calldata amounts) external {
        require(recipients.length == amounts.length, 
                "Recipients and amounts arrays must have the same length");

        uint256 totalTokens = 0;
        for (uint256 i = 0; i < recipients.length; i++) {
            totalTokens += amounts[i];
        }

        token.transferFrom(msg.sender, address(this), totalTokens);

        for (uint256 i = 0; i < recipients.length; i++) {
            token.transfer(recipients[i], amounts[i]);
        }
    }

    function batchTransferTokenSimple(
        IERC20 token, address[] calldata recipients, uint256[] calldata amounts)
        external
    {
        require(recipients.length == amounts.length, 
                "Recipients and amounts arrays must have the same length");
        for (uint256 i = 0; i < recipients.length; i++) {
            token.transferFrom(msg.sender, recipients[i], amounts[i]);
        }
    }
}

// coin-tools/contracts/src/BatchTransferEvm.sol
pragma solidity ^0.8.0; //SPDX-License-Identifier: MIT

contract BatchTransferEvm {
    function batchTransferEther(
        address payable[] calldata recipients, uint256[] calldata amounts) external payable {
        _batchTransferEther(recipients, amounts, 2300);
    }

    function batchTransferEtherCustomGas(
        address payable[] calldata recipients,
        uint256[] calldata amounts,
        uint256 transferGas
    ) external payable {
        _batchTransferEther(recipients, amounts, transferGas);
    }

    // Batch transfer Ether
    function _batchTransferEther(
        address payable[] calldata recipients, uint256[] calldata amounts, uint256 transferGas)
        internal
    {
        uint256 len = recipients.length;
        require(len == amounts.length, "Recipients and amounts arrays must have the same length");

        for (uint256 i = 0; i < len;) {
            (bool succ,) = recipients[i].call{value: amounts[i], gas: transferGas}("");
            require(succ, "Failed to send Ether");
            unchecked {
                ++i;
            }
        }

        uint256 remainingBalance = address(this).balance;
        require(remainingBalance == 0, 
                "Total transfer amount and individual transactions do not match");
    }

    // Batch transfer ERC20 tokens
    function batchTransferToken(
        address token, address[] calldata recipients, uint256[] calldata amounts) external {
        uint256 len = recipients.length;
        require(len == amounts.length, 
                "Recipients and amounts arrays must have the same length");
        uint256 totalTokens = 0;
        for (uint256 i = 0; i < len;) {
            totalTokens += amounts[i];
            unchecked {
                ++i;
            }
        }

        safeTransferFrom(token, msg.sender, address(this), totalTokens);
        for (uint256 i = 0; i < len;) {
            safeTransfer(token, recipients[i], amounts[i]);
            unchecked {
                ++i;
            }
        }
    }

    function batchTransferTokenSimple(
        address token, address[] calldata recipients, uint256[] calldata amounts)
        external
    {
        require(recipients.length == amounts.length, 
                "Recipients and amounts arrays must have the same length");
        uint256 len = recipients.length;
        for (uint256 i = 0; i < len;) {
            safeTransferFrom(token, msg.sender, recipients[i], amounts[i]);
            unchecked {
                ++i;
            }
        }
    }

    function safeTransfer(address token, address to, uint256 value) internal virtual {
        // bytes4(keccak256(bytes('transfer(address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0xa9059cbb, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 
                "TransferHelper: TRANSFER_FAILED");
    }

    function safeTransferFrom(address token, address from, address to, uint256 value) internal {
        // bytes4(keccak256(bytes('transferFrom(address,address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x23b872dd, from, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 
                "TransferHelper: TRANSFER_FROM_FAILED");
    }
}

// coin-tools/contracts/src/BatchTransferTron.sol
pragma solidity ^0.8.0; //SPDX-License-Identifier: MIT

import "./BatchTransferEvm.sol";

contract BatchTransferTron is BatchTransferEvm {
    function safeTransfer(address token, address to, uint256 value) internal override {
        // bytes4(keccak256(bytes('transfer(address,uint256)')));
        if (token == 0xa614f803B6FD780986A42c78Ec9c7f77e6DeD13C) {
            (bool success,) = token.call(abi.encodeWithSelector(0xa9059cbb, to, value));
            require(success, "TransferHelper: TRANSFER_FAILED, tron-usdt");
            return;
        }
        super.safeTransfer(token, to, value);
    }
}
