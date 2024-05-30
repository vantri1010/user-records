package contractCtl

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"user-records/env"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "user-records/api"
)

var (
	Env = env.NewEnv()
)

func LoadChainInfo() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	//env := env.NewEnv()
	client, err := ethclient.Dial(Env.RPC_URL)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := crypto.HexToECDSA(Env.PRIVATE_KEY)
	if err != nil {
		return nil, nil, err
	}

	return client, privateKey, nil
}

func TxOps(client *ethclient.Client, privateKey *ecdsa.PrivateKey, value int64) (*bind.TransactOpts, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type(contractDeploy.Write): publicKey is not of type *ecdsa.PublicKey")
	}
	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("msg.sender: 0x%x\n", addr)

	nonce, err := client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)  // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func Deploy(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*contracts.Api, common.Address, common.Hash, error) {

	auth, err := TxOps(client, privateKey, 0)
	if err != nil {
		log.Fatal(fmt.Sprintf("ERR (contractCtl.Deply): %v\n", err))
	}

	addr, tx, contract, err := contracts.DeployApi(auth, client)
	if err != nil {
		return nil, common.HexToAddress("0x0"), common.HexToHash("0x0"), err
	}
	fmt.Printf("Deployed CA: %x\nTx address: %v\n", addr, tx.Hash())
	return contract, addr, tx.Hash(), nil
}

func Load(client *ethclient.Client) (*contracts.Api, error) {
	ca := common.HexToAddress(Env.CONTRACT_ADDRESS)
	contractInstance, err := contracts.NewApi(ca, client)
	if err != nil {
		log.Fatal(err)
	}
	owner, err := contractInstance.GetUserCount(nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Contract owner: %x\n", owner)
	return contractInstance, nil
}
