package service

import (
	"avito/storage/accounting_report"
	"avito/storage/reserve_account"
	"avito/storage/user_balance"
	"database/sql"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

func ReplenishBalance(db *sql.DB, ub user_balance.UserBalance) error {
	newUB, err := user_balance.GetBalance(db, ub)
	if err != nil && err != sql.ErrNoRows {
		log.Println("getBalance err: ", err)
		return err
	}
	if err == sql.ErrNoRows {
		err = user_balance.UserBalanceInsert(db, ub)
		if err != nil {
			return err
		}
		return nil
	}
	newUB.Balance += ub.Balance
	err = user_balance.UserBalanceUpdate(db, newUB)
	if err != nil {
		return err
	}
	return nil
}

func ReserveMoney(db *sql.DB, ra reserve_account.ReverseAcc) error {
	ub := user_balance.UserBalance{Id: ra.Id, Balance: ra.Cost}

	newUb, err := user_balance.GetBalance(db, ub)
	if err != nil {
		return err
	}
	if newUb.Balance < ub.Balance {
		return errors.New("not enough money")
	}
	newUb.Balance -= ub.Balance

	err = reserve_account.ReverseAccInsert(db, ra)
	if err != nil {
		return err
	}
	err = user_balance.UserBalanceUpdate(db, newUb)
	if err != nil {
		return err
	}
	return nil
}

func UnReserveMoney(db *sql.DB, ra reserve_account.ReverseAcc) error {
	_, err := reserve_account.ReserveAccSelect(db, ra)
	if err != nil {
		return err
	}
	err = reserve_account.DeleteRow(db, ra)
	if err != nil {
		return err
	}
	ub := user_balance.UserBalance{Id: ra.Id, Balance: ra.Cost}
	err = ReplenishBalance(db, ub)
	if err != nil {
		return err
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
	repAcc := accounting_report.ReportAcc{Service: ra.Service, Cost: ra.Cost, OrderDate: date}
	err = accounting_report.ReportAccInsert(db, repAcc)
	if err != nil {
		return err
	}
	return nil
}

func GetReportCsv(db *sql.DB, dateRep accounting_report.DateReport) (string, error) {

	var fileName string
	items, err := accounting_report.SelectReport(db, dateRep)
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
