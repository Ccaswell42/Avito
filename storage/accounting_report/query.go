package accounting_report

import (
	"avito/storage/reserve_account"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type ReportAcc struct {
	Service   int
	Cost      int
	OrderDate string
}

func ReportAccInsert(db *sql.DB, ra ReportAcc) error {
	_, err := db.Exec(
		"INSERT INTO accounting_report (service, cost, order_date) VALUES ($1, $2, $3)",
		ra.Service, ra.Cost, ra.OrderDate)
	if err != nil {
		log.Println("insert problem", err)
		return err
	}
	return nil
}

func ReportAccSelect(db *sql.DB) error {
	rows, err := db.Query("SELECT * from accounting_report")
	if err != nil {
		log.Println("zapros err", err)
	}
	var items []ReportAcc
	for rows.Next() {
		ra := ReportAcc{}
		err = rows.Scan(&ra.Service, &ra.Cost, &ra.OrderDate)
		if err != nil {
			log.Println("scan problem", err)
			return err
		}
		items = append(items, ra)
	}
	err = rows.Close()
	if err != nil {
		log.Println("close problem", err)
		return err
	}
	for _, val := range items {
		fmt.Println(val)
	}
	return nil
}

func Revenue(db *sql.DB, ra reserve_account.ReverseAcc) error {
	_, err := reserve_account.ReserveAccSelect(db, ra)
	if err != nil {
		return err
	}
	err = reserve_account.DeleteRow(db, ra)
	if err != nil {
		return err
	}
	date := time.Now().Format("01-02-2006")
	repAcc := ReportAcc{Service: ra.Service, Cost: ra.Cost, OrderDate: date}
	err = ReportAccInsert(db, repAcc)
	if err != nil {
		return err
	}
	return nil
}
