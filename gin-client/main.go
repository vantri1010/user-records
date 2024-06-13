package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"user-records/contractCtl"
)

func main() {
	// connect to local blockchain network
	//env := env.NewEnv()
	client, privateKey, err := contractCtl.LoadChainInfo()
	if err != nil {
		log.Fatal(err)
	}

	contractInstance, err := contractCtl.Load(client)
	if err != nil {
		log.Fatal(err)
	}

	txOpts, err := contractCtl.TxOps(client, privateKey, 0)
	if err != nil {
		log.Fatal(err)
	}

	exAdd := common.HexToAddress("0x71C7656EC7ab88b098defB751B7401B5f6d8976F")
	var exMail [32]byte
	var stringMail = "trinv@gmail.com"
	copy(exMail[:], stringMail)

	var exTime, ok = new(big.Int).SetString("1717025328", 10)
	if !ok {
		log.Fatal("failed to convert big int")
	}

	log.Println("email: ", exMail)
	log.Println("address: ", exAdd)
	log.Println("timestamp: ", exTime)

	//txHash, err := contractInstance.InsertUser(txOpts, exAdd, exMail, exTime)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Tx Hash: %x\n", txHash.Hash())

	var index, _ = new(big.Int).SetString("0", 10)

	txHash, err := contractInstance.GetUserAtIndex(nil, index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tx Hash: %x\n", txHash)

	_ = txOpts

	//router := gin.Default()
	//router.Use(gin.Recovery())
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//router.POST("/create")
	//
	//router.Run()

}
