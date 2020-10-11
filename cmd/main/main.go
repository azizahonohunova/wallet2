package main

import (
	"log"
	"os"

	"github.com/azizahonohunova/wallet/pkg/types"
	"github.com/azizahonohunova/wallet/pkg/wallet"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()
	accounts := []*types.Account{}
	w := wallet.Service{}
	account, _ := w.RegisterAccount("12")
	accounts = append(accounts, account)
	account, _ = w.RegisterAccount("15")
	accounts = append(accounts, account)
	account, _ = w.RegisterAccount("13")
	accounts = append(accounts, account)
	account, _ = w.RegisterAccount("14")
	accounts = append(accounts, account)

	err = w.ExportToFile("text.txt")
	if err != nil {
		log.Println(err)
		return
	}
}
