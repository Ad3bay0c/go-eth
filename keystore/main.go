package kstore

import (
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const password = "password"

func GenerateKey() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(account.Address.Hex())
}

func Decrypt() (*keystore.Key, error) {
	f, err := ioutil.ReadFile("./tmp/UTC--2022-09-25T14-38-17.732338000Z--1bf4c911744f3cb4595a1454cac61446f118b7e2")
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(f, password)
	if err != nil {
		return nil, err
	}
	return key, nil
}
