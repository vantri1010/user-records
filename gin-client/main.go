package main

import (
	"github.com/gin-gonic/gin"
	"log"
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

	router := gin.Default()
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/create", routers.CreateRecord(contractInstance, txOpts))
	router.GET("/read/:address", routers.ReadRecord(contractInstance, options))
	router.GET("/getaddress/:index", routers.GetAddress(contractInstance, options))

	router.Run()

}
