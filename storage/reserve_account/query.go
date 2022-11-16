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

//func ReverseAccSelect(db *sql.DB) error {
//	rows, err := db.Query("SELECT * from reserve_account")
//	if err != nil {
//		log.Println("zapros err", err)
//	}
//	var items []ReverseAcc
//	for rows.Next() {
//		ra := ReverseAcc{}
//		err = rows.Scan(&ra.Id, &ra.Service, &ra.OrderId, &ra.Cost)
//		if err != nil {
//			log.Println("scan problem", err)
//			return err
//		}
//		items = append(items, ra)
//	}
//	err = rows.Close()
//	if err != nil {
//		log.Println("close problem", err)
//		return err
//	}
//	for _, val := range items {
//		fmt.Println(val)
//	}
//	return nil
//}

func ReserveAccSelect(db *sql.DB, ra ReverseAcc) (ReverseAcc, error) {

	var newRa ReverseAcc

	row := db.QueryRow("SELECT id, service, order_id, cost FROM reserve_account WHERE id = $1 AND service = $2 AND order_id = $3 AND cost = $4",
		ra.Id, ra.Service, ra.OrderId, ra.Cost)

	err := row.Scan(&newRa.Id, &newRa.Service, &newRa.OrderId, &newRa.Cost)
	fmt.Println(newRa)
	if err != nil {
		log.Println("scan////:", err)
		return newRa, err
	}
	return newRa, nil
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

	err = ReverseAccInsert(db, ra)
	if err != nil {
		return err
	}
	err = user_balance.UserBalanceUpdate(db, newUb)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRow(db *sql.DB, ra ReverseAcc) error {
	_, err := db.Exec("DELETE from reserve_account WHERE id = $1 AND service = $2 AND order_id = $3 AND cost = $4",
		ra.Id, ra.Service, ra.OrderId, ra.Cost)

	if err != nil {
		log.Println("delete////:", err)
		return err
	}
	return nil
}

func UnReserveMoney(db *sql.DB, ra ReverseAcc) error {
	_, err := ReserveAccSelect(db, ra)
	if err != nil {
		return err
	}
	err = DeleteRow(db, ra)
	if err != nil {
		return err
	}
	ub := user_balance.UserBalance{Id: ra.Id, Balance: ra.Cost}
	err = user_balance.ReplenishBalance(db, ub)
	if err != nil {
		return err
	}
	return nil
}
