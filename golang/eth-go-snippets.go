go get -u github.com/ethereum/go-ethereum
------------------------------------------------------------
We will learn how to

- connect to an Ethereum node
- call basic functions to get blocks or transactions from the blockchain
- create go-bindings for your smart contracts using abigen
- deploy and connect to your contracts.
- filter for past events, and
- listen to new events from our contracts.

------------------------------------------------------------

// Use a local node:
cl, err := ethclient.Dial(“/tmp/geth.ipc”)
// Use infura:
infura := “wss://goerli.infura.io/ws/v3/xxxxxx”
cl, err := ethclient.Dial(infura)
// Use a non-local node:
cl, err := ethclient.Dial("http://192.168.1.2:8545")

------------------------------------------------------------

// Retrieve a block by number
ctx := context.Background()
block, err := cl.BlockByNumber(ctx, big.NewInt(123))
// Get Balance of an account (nil means at newest block)
addr := common.HexToAddress("0xb02A2EdA1b317FBd16760...")
balance, err := cl.BalanceAt(ctx, addr, nil)
// Send transaction
tx := new(types.Transaction)
err = cl.SendTransaction(ctx, tx)
// Get sync progress of the node
progress, err := cl.SyncProgress(ctx)

------------------------------------------------------------

// Retrieve the pending nonce for an account
nonce, err := cl.NonceAt(ctx, addr, nil)
to := common.HexToAddress("0xABCD")
amount := big.NewInt(10 * params.GWei)
gasLimit := uint64(21000)
gasPrice := big.NewInt(10 * params.GWei)
data := []byte{}
// Create a raw unsigned transaction
tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

------------------------------------------------------------

// Use secret key hex string to sign a raw transaction
SK := "0x0000"
sk := crypto.ToECDSAUnsafe(common.FromHex(SK))// Sign the transaction
signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), sk)
// You could also create a TransactOpts object
opts := bind.NewKeyedTransactor(sk)
// To get the address corresponding to your private key
addr := crypto.PubkeyToAddress(sk.PublicKey)

------------------------------------------------------------

// Open Keystore
ks := keystore.NewKeyStore("/home/matematik/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
// Create an account
acc, err := ks.NewAccount("password")
// List all accounts
accs := ks.Accounts()
// Unlock an account
ks.Unlock(accs[0], "password")
// Create a TransactOpts object
ksOpts, err := bind.NewKeyStoreTransactor(ks, accs[0])

------------------------------------------------------------

// If you have a bind.TransactOpts object you can sign a transaction
sigTx, err := opts.Signer(types.NewEIP155Signer(nil), senderAddr, tx)

------------------------------------------------------------

// Connect to a node
backend, err := ethclient.Dial("/tmp/geth.ipc")
// Get the contract address from hex string
addr := common.HexToAddress("0x000")
// Bind to an already deployed contract
ctr, err := contract.NewContract(addr, backend)

------------------------------------------------------------

// Deploy a new contract
addr, tx, ctr, err := coolcontract.DeployCoolContract(transactOpts, backend)
// Check if the contract was deployed successfully
_, err = bind.WaitDeployed(ctx, backend, tx)

------------------------------------------------------------

// Call a pure/view function
callOpts := &bind.CallOpts{Context: ctx, Pending: false}
bal, err := ctr.SeeBalance(callOpts)
// Call a normal function
tx, err := ctr.Deposit(transactOpts)
receipt, err := bind.WaitMined(ctx, backend, tx)
if receipt.Status != types.ReceiptStatusSuccessful {
    panic("Call failed")
}

------------------------------------------------------------

// Filter for a Deposited event
filterOpts := &bind.FilterOpts{Context: ctx, Start: 9000000, End: nil}
itr, err := ctr.FilterDeposited(filterOpts)
// Loop over all found events
for itr.Next() {
    event := itr.Event
    // Print out all caller addresses
    fmt.Printf(event.Addr.Hex())
}

------------------------------------------------------------

// Watch for a Deposited event
watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
// Setup a channel for results
channel := make(chan *coolcontract.CoolContractDeposited)
// Start a goroutine which watches new events
go func() {
    sub, err := ctr.WatchDeposited(watchOpts, channel)
    defer sub.Unsubscribe()
}()
// Receive events from the channel
event := <-channel

------------------------------------------------------------

// Parse event from types.Log
log := *new(types.Log)
event, err := ctr.ParseDeposited(log)

------------------------------------------------------------

// Create a new SimulatedBackend with a default allocation
backend := backends.NewSimulatedBackend(core.DefaultGenesisBlock().Alloc, 9000000)
bal, err := backend.BalanceAt(ctx, common.HexToAddress("0x000"), nil)
		
// Create a meaningful allocation with a faucet secret key
faucetSK, err := crypto.GenerateKey()
faucetAddr := crypto.PubkeyToAddress(faucetSK.PublicKey)
addr := map[common.Address]core.GenesisAccount{
	common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
	common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
	common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
	common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
	common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
	common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
	common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
	common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
	faucetAddr: {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
}
alloc := core.GenesisAlloc(addr)
backend := backends.NewSimulatedBackend(alloc, 9000000)

------------------------------------------------------------

type SimulatedBackend struct {
	backends.SimulatedBackend
}

func (s *SimulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if err := s.SimulatedBackend.SendTransaction(ctx, tx); err != nil {
		return err
	}
	s.Commit()
	return nil
}

------------------------------------------------------------

tx, tx2, tx3 := new(types.Transaction), new(types.Transaction), new(types.Transaction)
// Send and abort both transactions
backend.SendTransaction(ctx, tx)
backend.SendTransaction(ctx, tx2)
backend.Rollback()

// Commits transaction to the chain
backend.SendTransaction(ctx, tx3)
backend.Commit()

// Use adjust time to test time-dependent functions
backend.AdjustTime(24 * time.Hour)
// Adjust time can only be called on an empty block and you need to call commit afterwards
backend.Commit()

------------------------------------------------------------
