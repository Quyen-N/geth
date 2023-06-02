package main

import (
	"context"
	"log"
	"q-geth/keys"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	logger := log.Default()

	client, err := ethclient.Dial("https://matic-mumbai.chainstacklabs.com")
	if err != nil {
		logger.Fatal("lost connection to rpc", err)
	}
	account := common.HexToAddress(keys.ADDR)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		logger.Panic("get balance error", err)
	}
	logger.Println(balance)
}
