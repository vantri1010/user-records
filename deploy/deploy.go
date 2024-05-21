package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"user-records/api"
)

func main() {
	// connect to local blockchain network
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

	//hexKey := "cd0d310e996dc59f90ebdd53bd44bb185e7eacef2dafd1746c82aae0e45d945e"
	//privateKeyBin, err := hexutil.Decode(hexKey)
	//if err != nil {
	//	log.Fatal(err)
	//}

	privateKey, err := crypto.HexToECDSA("cef51cb8cb8bce652445f0958daaf1834981dceeba1fee295c4c4da0a160469f")
	//privateKey, err := crypto.ToECDSA(privateKeyBin)
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

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	panic(err)
	//}

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
	auth.GasPrice = big.NewInt(10000000000)

	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Api contract deployed to %s\n", address.Hex())
	fmt.Printf("Tx: %s\n", tx.Hash().Hex())

	_, _ = instance, tx

}
