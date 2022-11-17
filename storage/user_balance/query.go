package user_balance

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type UserBalance struct {
	Id      int  `json:"Id,omitempty"`
	Balance uint `json:"Balance,omitempty"`
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
		log.Println("update problem: ", err)
		return err
	}
	return nil
}

func GetBalance(db *sql.DB, ub UserBalance) (UserBalance, error) {
	var resp UserBalance
	row := db.QueryRow("SELECT balance FROM user_balance WHERE id = $1", ub.Id)

	err := row.Scan(&resp.Balance)
	if err != nil {
		log.Println("scan error: ", err)
		return resp, err
	}
	return resp, nil
}
