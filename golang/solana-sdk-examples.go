How to create a Solana wallet with Go

go get -u github.com/portto/solana-go-sdk

// Connect to the Solana network

package main

import (
    "context"
    "fmt"
    "github.com/portto/solana-go-sdk/client"
    "github.com/portto/solana-go-sdk/client/rpc"
)

func main() {
    // create a RPC client
    c := client.NewClient(rpc.MainnetRPCEndpoint)

    // get the current running Solana version
    response, err := c.GetVersion(context.TODO())
    if err != nil {
        panic(err)
    }

    fmt.Println("version", response.SolanaCore)
}

// Creating a new Solana wallet

package main

import (
    "fmt"
    "github.com/portto/solana-go-sdk/types"
)

func main() {
    // create a new wallet using types.NewAccount()
    wallet := types.NewAccount()

    // display the wallet public and private keys
    fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())
    fmt.Println("Private Key:", wallet.PrivateKey)
}

// Importing a Solana wallet

types.AccountFromBase58("")      // import a wallet with base58 private key
types.AccountFromBytes([]byte{}) // import a wallet with bytes slice private key
types.AccountFromHex("")         // import a wallet with hex private key

package main

import (
    "fmt"
    "github.com/portto/solana-go-sdk/types"
)

func main() {
    // create a new wallet
    wallet := types.NewAccount()
    fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())

    // import the wallet using its private key
    importedWallet, err := types.AccountFromBytes(wallet.PrivateKey)

    // check for errors
    if err != nil {
        panic(err)
    }

    // display the imported wallet public and private keys
    fmt.Println("Imported Wallet Address:", importedWallet.PublicKey.ToBase58())
}

// Fetching the balance of a Solana wallet

package main

import (
    "context"
    "fmt"
    "github.com/portto/solana-go-sdk/client"
    "github.com/portto/solana-go-sdk/client/rpc"
    "github.com/portto/solana-go-sdk/types"
)

func main() {
    // create a RPC client
    c := client.NewClient(rpc.DevnetRPCEndpoint)

    // create a new wallet
    wallet := types.NewAccount()

    // request for 1 SOL airdrop using RequestAirdrop()
    txhash, err := c.RequestAirdrop(
        context.TODO(), // request context
        wallet.PublicKey.ToBase58(), // wallet address requesting airdrop
        1e9, // amount of SOL in lamport
    )

    // check for errors
    if err != nil {
        panic(err)
    }

    fmt.Println("Transaction Hash:", txhash)
}

package main

import (
    "context"
    "fmt"
    "github.com/portto/solana-go-sdk/client"
    "github.com/portto/solana-go-sdk/client/rpc"
)

func main() {
    // create a RPC client
    c := client.NewClient(rpc.DevnetRPCEndpoint)

    // fetch the balance using GetBalance()
    balance, err := c.GetBalance(
        context.TODO(), // request context
        "8LdDAFdGuvZdhhnheUv9jVtiv9wQT3eTk2E46FodZP38", // wallet to fetch balance for
    )

    // check for errors
    if err != nil {
        panic(err)
    }

    fmt.Println("Wallet Balance in Lamport:", balance)
    fmt.Println("Wallet Balance in SOL:", balance/1e9)
}

// Transferring Solana to another wallet

package main

import (
    "context"
    "fmt"
    "github.com/portto/solana-go-sdk/client"
    "github.com/portto/solana-go-sdk/client/rpc"
    "github.com/portto/solana-go-sdk/common"
    "github.com/portto/solana-go-sdk/program/sysprog"
    "github.com/portto/solana-go-sdk/types"
)

func main() {
    // create a RPC client
    c := client.NewClient(rpc.DevnetRPCEndpoint)

    // import a wallet with Devnet balance
    wallet, _ := types.AccountFromBytes([]byte{})

    // fetch the most recent blockhash
    response, err := c.GetRecentBlockhash(context.TODO())
    if err != nil {
        panic(err)
    }

    // make a transfer message with the latest block hash
    message := types.NewMessage(
        wallet.PublicKey, // public key of the transaction signer
        []types.Instruction{
            sysprog.Transfer(
                wallet.PublicKey, // public key of the transaction sender
                // wallet address of the transaction receiver
                common.PublicKeyFromString("8t88TuqUxDMVpYGHcVoXnBCAH7TPrdZ7ydr4xqcNu2Ym"), 
                1e9, // transaction amount in lamport
            ),
        },
        response.Blockhash, // recent block hash
    )

    // create a transaction with the message and TX signer
    tx, err := types.NewTransaction(message, []types.Account{wallet, wallet})
    if err != nil {
        panic(err)
    }

    // send the transaction to the blockchain
    txhash, err := c.SendTransaction2(context.TODO(), tx)
    if err != nil {
        panic(err)
    }

    fmt.Println("Transaction Hash:", txhash)
}
