package data

import (
	"encoding/json"
	"github.com/pandulaDW/go-ds-algo/hashing"
	"io/ioutil"
	"log"
)

type SingleAccount struct {
	ID        string   `json:"_id"`
	AccountID int      `json:"account_id"`
	Limit     int      `json:"limit"`
	Products  []string `json:"products"`
}

type Accounts struct {
	Data []SingleAccount `json:"data"`
}

func ReadAccountData() *Accounts {
	content, err := ioutil.ReadFile("data/account.json")
	if err != nil {
		log.Fatal(err)
	}

	accounts := new(Accounts)
	err = json.Unmarshal(content, accounts)
	if err != nil {
		log.Fatal(err)
	}

	return accounts
}

func ConvertToHashTableType(accounts *Accounts) []*hashing.DataWithStringID {
	data := make([]*hashing.DataWithStringID, 0, len(accounts.Data))

	for _, account := range accounts.Data {
		hashItem := &hashing.DataWithStringID{ID: account.ID, Content: account}
		data = append(data, hashItem)
	}

	return data
}
