package main

import (
	"avito/config"
	"avito/storage"
	"avito/storage/reserve_account"
	"log"
)

func main() {

	//rep := accounting_report.ReportAcc{Service: 13, Cost: 1000, OrderDate: "Jan-08-1999"}
	revers := reserve_account.ReverseAcc{Id: 1, Service: 12, OrderId: 22, Cost: 3500}
	//ub := user_balance.UserBalance{Id: 1, Balance: 30000}
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := storage.ConnectToDb(conf)
	if err != nil {
		log.Fatal(err)
	}
	reserve_account.ReverseAccInsert(db, revers)
	reserve_account.ReverseAccSelect(db)
}
