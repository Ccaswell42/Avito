package reserve_account

import (
	"avito/storage/user_balance"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type ReverseAcc struct {
	Id      int `json:"Id"`
	Service int `json:"Service"`
	OrderId int `json:"OrderId"`
	Cost    int `json:"Cost"`
}

func ReverseAccInsert(db *sql.DB, ra ReverseAcc) error {
	_, err := db.Exec(
		"INSERT INTO reserve_account (id, service, order_id, cost) VALUES ($1, $2, $3, $4 )",
		ra.Id, ra.Service, ra.OrderId, ra.Cost)
	if err != nil {
		log.Println("insert problem", err)
		return err
	}
	return nil
}

func ReverseAccSelect(db *sql.DB) error {
	rows, err := db.Query("SELECT * from reserve_account")
	if err != nil {
		log.Println("zapros err", err)
	}
	var items []ReverseAcc
	for rows.Next() {
		ra := ReverseAcc{}
		err = rows.Scan(&ra.Id, &ra.Service, &ra.OrderId, &ra.Cost)
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

func ReserveMoney(db *sql.DB, ra ReverseAcc) error {
	ub := user_balance.UserBalance{Id: ra.Id, Balance: ra.Cost}

	newUb, err := user_balance.GetBalance(db, ub)
	if err != nil {
		return err
	}
	if newUb.Balance < ub.Balance {
		return errors.New("not enough money")
	}
	newUb.Balance -= ub.Balance

	err = user_balance.UserBalanceUpdate(db, newUb)
	if err != nil {
		return err
	}
	err = ReverseAccInsert(db, ra)
	if err != nil {
		return err
	}
	return nil
}
