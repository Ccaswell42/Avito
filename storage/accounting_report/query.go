package accounting_report

import (
	"avito/storage/reserve_account"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
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

func RepAccSelect(db *sql.DB) error {

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

func GetReportCsv(db *sql.DB, dateRep DateReport) (string, error) {

	var fileName string
	items, err := SelectReport(db, dateRep)
	if len(items) == 0 {
		return fileName, errors.New("no entries in the report")
	}
	if err != nil {
		return fileName, err
	}
	csvMap := make(map[int]uint)
	for _, value := range items {
		csvMap[value.Service] += value.Cost
	}

	csvFile, err := os.Create("data.csv")
	if err != nil {
		return fileName, err
	}
	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = ';'
	err = csvWriter.Write([]string{"название услуги", "общая сумма выручки за отчетный период"})
	for key, value := range csvMap {
		keyStr := strconv.Itoa(key)
		valueStr := strconv.FormatUint(uint64(value), 10)
		fmt.Println(keyStr, valueStr)
		err = csvWriter.Write([]string{keyStr, valueStr})
		if err != nil {
			return fileName, errors.New("can't write data in csv-file")
		}
	}
	fileName = csvFile.Name()
	csvWriter.Flush()
	if csvWriter.Error() != nil {
		return fileName, err
	}
	return fileName, nil
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
	//for _, val := range items {
	//	fmt.Println(val)
	//}

	return items, nil
}
