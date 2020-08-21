package main

import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/types"
	"time"
)

type Wallet struct {
	account types.Account
	c       *client.Client
}

func CreateNewWallet(RPCEndpoint string) Wallet {
	// create a new wallet using types.NewAccount()
	return Wallet{
		types.NewAccount(),
		client.NewClient(RPCEndpoint),
	}
}

func ImportOldWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {
	// import a wallet with bytes slice private key
	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

func (w Wallet) RequestAirdrop(amount uint64) (string, error) {
	// request for SOL using RequestAirdrop()
	txhash, err := w.c.RequestAirdrop(
		context.TODO(),                 // request context
		w.account.PublicKey.ToBase58(), // wallet address requesting airdrop
		amount,                         // amount of SOL in lamport
	)
	if err != nil {
		return "", err
	}

	return txhash, nil
}

func (w Wallet) GetBalance() (uint64, error) {
	// fetch the balance using GetBalance()
	balance, err := w.c.GetBalance(
		context.TODO(),                 // request context
		w.account.PublicKey.ToBase58(), // wallet to fetch balance for
	)
	if err != nil {
		return 0, nil
	}

	return balance, nil
}

func (w Wallet) Transfer(receiver string, amount uint64) (string, error) {
	// fetch the most recent blockhash
	response, err := w.c.GetRecentBlockhash(context.TODO())
	if err != nil {
		return "", err
	}

	// make a transfer message with the latest block hash
	message := types.NewMessage(
		w.account.PublicKey, // public key of the transaction signer
		[]types.Instruction{
			sysprog.Transfer(
				w.account.PublicKey,                  // public key of the transaction sender
				common.PublicKeyFromString(receiver), // wallet address of the transaction receiver
				amount,                               // transaction amount in lamport
			),
		},
		response.Blockhash, // recent block hash
	)

	// create a transaction with the message and TX signer
	tx, err := types.NewTransaction(message, []types.Account{w.account, w.account})
	if err != nil {
		return "", err
	}

	// send the transaction to the blockchain
	txhash, err := w.c.SendTransaction2(context.TODO(), tx)
	if err != nil {
		return "", err
	}

	return txhash, nil
}

func main() {
	// create a new wallet
	wallet := CreateNewWallet(rpc.DevnetRPCEndpoint)

	// request for an airdrop
	fmt.Println(wallet.RequestAirdrop(1e9))
	time.Sleep(time.Second * 20)

	// make transfer to another wallet
	fmt.Println(wallet.Transfer("8t88TuqUxDMVpYGHcVoXnBCAH7TPrdZ7ydr4xqcNu2Ym", 5e8))
	time.Sleep(time.Second * 20)

	// fetch wallet balance
	fmt.Println(wallet.GetBalance())
}