package main

import (
	"fraud-detector-api/pkg"
	"log"
)

func main() {
	if err := pkg.GenerateTransactions(300, "transactions.json"); err != nil {
		log.Fatal(err)
	}
	log.Println("transactions.json berhasil dibuat!")
}
