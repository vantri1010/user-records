package main

import (
	"log"
	"user-records/contractCtl"
)

func main() {
	// connect to local blockchain network
	client, privateKey, err := contractCtl.LoadChainInfo()
	if err != nil {
		log.Fatal(err)
	}

	contractInstance, ca, txa, err := contractCtl.Deploy(client, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	_ = ca
	_ = contractInstance
	_ = txa

}
