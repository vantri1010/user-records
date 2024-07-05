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
	router.GET("/getuserbyaddr/:address", routers.GetUser(contractInstance, options))
	router.GET("/getaddressbyid/:index", routers.GetAddress(contractInstance, options))
	router.GET("/getuserbyid/:index", routers.GetUserByID(contractInstance, options))
	router.GET("/listusers", routers.ListUsers(contractInstance, options))
	router.DELETE("/deleteuserbyaddr/:address", routers.DeleteRecord(contractInstance, txOpts))

	router.Run()
}
