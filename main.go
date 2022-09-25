package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Ad3bay0c/eth-contract-go/go-contract"
	kstore "github.com/Ad3bay0c/eth-contract-go/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
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
	//addr, txx, _, err := contract.DeployContract(tx, client, "1.0")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("%+v ---> %+v ---> %+v", addr.String(), addr.Hex(), txx.Hash().Hex())

	// deployed smart contract address
	address := common.HexToAddress("0x37310eB01871BfA8538626107700e092082d87a7")

	st, err := contract.NewContract(address, client)

	// Access the version of the deployed contract
	version, _ := st.Version(nil)
	log.Println(version)

	// Set Item
	auth, _ := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	auth.GasPrice = gasPrice
	auth.GasLimit = 300000
	auth.Value = big.NewInt(0)

	k := [32]byte{}
	copy(k[:], []byte("one"))
	v := [32]byte{}
	copy(v[:], []byte("1"))

	t, _ := st.SetItem(auth, k, v)

	log.Printf("%+v", t.Hash().Hex())

	// Get the item from the contract
	items, _ := st.Items(nil, k)
	log.Printf("%s", items)
	//kstore.GenerateKey()
	//log.Println(wallet.GetBlockNumber(client))
}
