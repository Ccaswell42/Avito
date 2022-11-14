package user_balance

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type UserBalance struct {
	Id      int
	Balance int
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
