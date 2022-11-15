package main

import (
	"avito/config"
	"avito/server"
	"avito/storage"
	"avito/storage/user_balance"
	"log"
	"net/http"
)

func main() {

	//rep := accounting_report.ReportAcc{Service: 13, Cost: 1000, OrderDate: "Jan-08-1999"}
	//revers := reserve_account.ReverseAcc{Id: 1, Service: 12, OrderId: 22, Cost: 3500}
	ub := user_balance.UserBalance{Id: 31, Balance: 30000}
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := storage.ConnectToDb(conf)
	if err != nil {
		log.Fatal(err)
	}
	user_balance.UserBalanceInsert(db, ub)
	user_balance.UserBalanceSelect(db)
	srv := &http.Server{}
	server.LaunchServer(conf, srv, db)
}
