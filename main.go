package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Ad3bay0c/eth-contract-go/go-contract"
	kstore "github.com/Ad3bay0c/eth-contract-go/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const address = "0x5fD630A0470cd7E21624a494c309d40DFe82633D"

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/3abadf556f434f56838a0b5d4448ba49")
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	key, err := kstore.Decrypt()
	if err != nil {
		log.Fatal(err)
	}

	// creates a transaction signer from a single private key
	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// gets the next nonce to use
	nonce, err := client.PendingNonceAt(context.Background(), key.Address)
	tx.GasPrice = gasPrice
	tx.GasLimit = 3000000
	tx.Nonce = big.NewInt(int64(nonce))

	// deploys a contract
	addr, txx, _, err := contract.DeployContract(tx, client, "1.0")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v ---> %+v ---> %+v", addr.String(), addr.Hex(), txx.Hash().Hex())

	//kstore.GenerateKey()
	//log.Println(wallet.GetBlockNumber(client))
}
