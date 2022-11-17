package accounting_report

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type ReportAcc struct {
	Service   int    `json:"Service"`
	Cost      uint   `json:"Cost"`
	OrderDate string `json:"OrderDate"`
}

type DateReport struct {
	Year  string `json:"Year"`
	Month string `json:"Month"`
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

func SelectReport(db *sql.DB, dateRep DateReport) ([]ReportAcc, error) {
	var items []ReportAcc
	rows, err := db.Query("SELECT service, cost from accounting_report WHERE EXTRACT(Month from order_date) = $1 "+
		"AND EXTRACT(YEAR from order_date) = $2",
		dateRep.Month, dateRep.Year)
	if err != nil {
		log.Println("query err", err)
		return items, err
	}

	for rows.Next() {
		ra := ReportAcc{}
		err = rows.Scan(&ra.Service, &ra.Cost)
		if err != nil {
			log.Println("scan problem", err)
			return items, err
		}
		items = append(items, ra)
	}
	err = rows.Close()
	if err != nil {
		log.Println("close problem", err)
		return items, err
	}
	return items, nil
}
