package main

import (
	"github.com/Ad3bay0c/eth-contract-go/wallet"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const address = "0x5fD630A0470cd7E21624a494c309d40DFe82633D"

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	balance, err := wallet.GetBalance(address, client)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(balance, "ETH")
	log.Println(wallet.GetBlockNumber(client))
}
