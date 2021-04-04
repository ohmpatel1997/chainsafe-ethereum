package main

import (
	contract2 "chainsafe/internal/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

// Test message gets updated correctly
func TestSetMessage(t *testing.T) {

	//Setup simulated blockchain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, uint64(3000000))

	//Deploy contract

	_, _, contract, _ := contract2.DeployIPFS(
		auth,
		blockchain,
	)

	// commit all pending transactions
	blockchain.Commit()

	_, err := contract.SetIPFSIdentifier(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, "QmTBK2YrdgUg8tpiovbVWr43QGirNGRW1PMpMtDfPp8Nu6")

	if err != nil {
		t.Errorf("Error setting CID. Error: %v", err.Error())
	}

	blockchain.Commit()

	cids, err := contract.PrintIPFSIdentifiers(nil)
	if err != nil {
		t.Errorf("Error while fetching CIDs. Error %v", err)
	}

	blockchain.Commit()

	fmt.Println("cids-->", cids)
	t.Log(cids)
}
