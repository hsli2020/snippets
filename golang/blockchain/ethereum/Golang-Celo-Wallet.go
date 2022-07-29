// Building a Celo Wallet using Golang
// https://celo.academy/t/building-a-celo-wallet-using-golang/199
// https://github.com/orgs/celo-examples/repositories
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
)

const nodeURL = "https://alfajores-forno.celo-testnet.org"

func main() {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}
	defer client.Close()

	// replace with your own private key
	privateKeyHex := ""

	// Load the private key
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	newWallet()

	existingWallet()

	walletProtect()

	send(client)

	checkWalletBalance(client, privateKey)
}

func send(client *ethclient.Client) {
	// (replace with your own private key)
	privateKeyHex := ""

	// Load the private key
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Send a transaction
	toAddress := common.HexToAddress("0x8BdDeC1b7841bF9eb680bE911bd22051f6a00815")
	value := big.NewInt(100000000000000000) // 0.1 CELO

	txHash, err := sendTransaction(client, privateKey, toAddress, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash: %s\n", txHash.Hex())
}

func newWallet() {
	// Generate a new Celo wallet
	privateKey, publicKey, celoAddress, err := generateNewWallet()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key: %x\n", privateKey.D)
	fmt.Printf("Public key: %x\n", publicKey)
	fmt.Printf("Celo address: %s\n", celoAddress.Hex())
}

func existingWallet() {
	// replace with your own private key
	privateKeyHex := ""

	// Import an existing Celo wallet
	privateKey, publicKey, celoAddress, err := importWallet(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Private key: %x\n", privateKey.D)
	fmt.Printf("Public key: %x\n", publicKey)
	fmt.Printf("Celo address: %s\n", celoAddress.Hex())
}

func walletProtect() {
	// replace with your own private key
	privateKeyHex := ""

	// Load the private key
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Example password for encrypting the wallet
	password := "your-strong-password"

	// Encrypt the wallet
	encryptedWallet, err := encryptWallet(privateKey, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted wallet:", encryptedWallet)

	// Decrypt the wallet
	decryptedPrivateKey, err := decryptWallet(encryptedWallet, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decrypted private key: %x\n", decryptedPrivateKey.D)
}

func importWallet(privateKeyHex string) (*ecdsa.PrivateKey, []byte, common.Address, error) {
	// Decode the private key from hex string
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, nil, common.Address{}, err
	}

	// Load the private key
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, nil, common.Address{}, err
	}

	// Derive the public key from the private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// Convert the public key to compressed format
	publicKeyBytes := secp256k1.CompressPubkey(publicKey.X, publicKey.Y)

	// Generate the Celo address from the public key
	celoAddress := crypto.PubkeyToAddress(*publicKey)

	return privateKey, publicKeyBytes, celoAddress, nil
}

func generateNewWallet() (*ecdsa.PrivateKey, []byte, common.Address, error) {
	// Generate a private key using the secp256k1 curve
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, common.Address{}, err
	}

	// Derive the public key from the private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	// Convert the public key to compressed format
	publicKeyBytes := secp256k1.CompressPubkey(publicKey.X, publicKey.Y)

	// Generate the Celo address from the public key
	celoAddress := crypto.PubkeyToAddress(*publicKey)

	return privateKey, publicKeyBytes, celoAddress, nil
}

func encryptWallet(privateKey *ecdsa.PrivateKey, password string) (string, error) {
	keyJSON, err := keystore.EncryptKey(&keystore.Key{
		PrivateKey: privateKey,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
	}, password, keystore.StandardScryptN, keystore.StandardScryptP)

	if err != nil {
		return "", err
	}

	return string(keyJSON), nil
}

func decryptWallet(encryptedWallet string, password string) (*ecdsa.PrivateKey, error) {
	key, err := keystore.DecryptKey([]byte(encryptedWallet), password)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}

func checkWalletBalance(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {
	// Check wallet balance
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wallet balance: %s\n", balance.String())
}

func sendTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, toAddress common.Address, value *big.Int) (common.Hash, error) {
	// Create a new transactor
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return common.Hash{}, err
	}

	// Assuming you have an Ethereum client instance named 'client'
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get the gas price: %v", err)
	}

	// Increase the gas price by 10% (or more, depending on how quickly you want the transaction to be mined)
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(11)).Div(gasPrice, big.NewInt(10))

	// Assuming you have an Ethereum client instance named 'client'
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get the chain ID: %v", err)
	}

	// Assuming you have a private key instance named 'privateKey'
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create the auth object: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(21000)
	auth.GasPrice = gasPrice

	tx := types.NewTransaction(auth.Nonce.Uint64(), toAddress, auth.Value, auth.GasLimit, auth.GasPrice, nil)
	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		return common.Hash{}, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, err
	}

	return signedTx.Hash(), nil
}