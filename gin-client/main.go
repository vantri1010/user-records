package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"user-records/contractCtl"
	"user-records/gin-client/routers"
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

	options, err := contractCtl.CallOpts(txOpts.From)
	_ = options

	// Get record inserted below at index 0
	var index, _ = new(big.Int).SetString("1", 10)

	txHash, err := contractInstance.GetUserAtIndex(nil, index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tx Hash: %x\n", txHash)

	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/create", routers.CreateRecord(contractInstance, txOpts))

	router.Run()

}
