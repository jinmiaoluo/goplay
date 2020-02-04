package main

import (
	"encoding/json"
	"io/ioutil"
)

type Account struct {
	Name string
	Age  int
}

func Save(account *Account) error {
	jsonBytes, err := json.MarshalIndent(account, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./account.json", jsonBytes, 0600)
}

func main() {
	var account *Account
	account = &Account{
		Name: "Jinmiao Luo",
		Age:  24,
	}

	Save(account)
}
