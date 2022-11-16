package user_balance

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type UserBalance struct {
	Id      int `json:"Id,omitempty"`
	Balance int `json:"Balance,omitempty"`
}

func UserBalanceInsert(db *sql.DB, ub UserBalance) error {
	_, err := db.Exec(
		"INSERT INTO user_balance (id, balance) VALUES ($1, $2 )", ub.Id, ub.Balance)
	if err != nil {
		log.Println("insert problem", err)
		return err
	}
	return nil
}

func UserBalanceUpdate(db *sql.DB, ub UserBalance) error {
	_, err := db.Exec(
		"UPDATE user_balance SET balance = $1", ub.Balance)
	if err != nil {
		log.Println("update problem", err)
		return err
	}
	return nil
}

func UserBalanceSelect(db *sql.DB) error {
	rows, err := db.Query("SELECT * from user_balance")
	if err != nil {
		log.Println("zapros err", err)
	}
	var items []UserBalance
	for rows.Next() {
		ub := UserBalance{}
		err = rows.Scan(&ub.Id, &ub.Balance)
		if err != nil {
			log.Println("scan problem", err)
			return err
		}
		items = append(items, ub)
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

func GetBalance(db *sql.DB, ub UserBalance) (UserBalance, error) {
	var resp UserBalance
	row := db.QueryRow("SELECT balance FROM user_balance WHERE id = $1", ub.Id)

	err := row.Scan(&resp.Balance)
	if err != nil {
		log.Println("scan////:", err)
		return resp, err
	}
	return resp, nil
}

func ReplenishBalance(db *sql.DB, ub UserBalance) error {
	newUB, err := GetBalance(db, ub)
	if err != nil && err != sql.ErrNoRows {
		log.Println("getbalance err///:", err)
		return err
	}
	if err == sql.ErrNoRows {
		err = UserBalanceInsert(db, ub)
		if err != nil {
			return err
		}
		return nil
	}
	newUB.Balance += ub.Balance
	err = UserBalanceUpdate(db, newUB)
	if err != nil {
		return err
	}
	return nil
}
