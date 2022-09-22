package wallet

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(key string, client *ethclient.Client) (*big.Int, error) {
	account := common.HexToAddress(key)

	return client.BalanceAt(context.Background(), account, nil)
}

func GetBlockNumber(client *ethclient.Client) (uint64, error) {

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}
