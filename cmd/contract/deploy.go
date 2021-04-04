package main

import (
	"chainsafe/internal/common"
	contract "chainsafe/internal/contracts"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"strings"

	"math/big"
)

const pathToEnvFile = "../env.yml"

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	schema, err := common.ReadConf(pathToEnvFile)
	if err != nil {
		logrus.Panicf("unable to read env file. Error: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(schema.PrivateKey, "0x"))
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	address, tx, instance, err := contract.DeployIPFS(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
}

//0xBbC4b74f844214DB63D2a66552d7FB78abE8100c -> ganache CLI

//0x786D3Afdba9d314046d2468Eb80101eB6B783e9C -> ganache UI
