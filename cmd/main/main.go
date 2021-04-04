package main

import (
	"chainsafe/internal/common"
	contract "chainsafe/internal/contracts"
	"chainsafe/internal/service"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"math/big"
	"os"
	"strings"
)

const pathToEnvFile = "../env.yml"

func main() {
	args := os.Args

	if len(args) <= 1 {
		logrus.Panic("invalid arguments. please specify the file name")
	}

	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		logrus.Panic("can not able to open file. make sure file is in the root directory of project")
	}

	ipfs, err := service.NewIPFSService()
	if err != nil {
		logrus.Panic(err)
	}

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	schema, err := common.ReadConf(pathToEnvFile)
	if err != nil {
		logrus.Panicf("unable to read env file. Error: %v", err)
	}

	conn, err := contract.NewIPFS(ethCommon.HexToAddress(schema.ContractAddress), client)
	if err != nil {
		panic(err)
	}

	cid, err := ipfs.SaveFile(context.Background(), file)
	if err != nil {
		logrus.Panic(err)
	}

	//connect to ethereum
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = conn.SetIPFSIdentifier(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, cid)

	if err != nil {
		logrus.Panicf("error while storing CID into etheruem blockchain. %v", logrus.Fields{"Error": err.Error()})
	}

	cids, err := conn.PrintIPFSIdentifiers(nil)
	if err != nil {
		logrus.Errorf("error while fetching CIDs array %v", logrus.Fields{"Error": err.Error()})
	}

	logrus.Info("Array of CIDs: ", cids)

}
