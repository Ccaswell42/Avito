package user_balance

import (
	"database/sql"
	"fmt"
	"log"
)

type userBalance struct {
	id      int
	balance int
}

func userBalanceInsert(db *sql.DB, ub userBalance) error {
	_, err := db.Exec(
		"INSERT INTO user_balance (id, balance) VALUES ($1, $2 )", ub.id, ub.balance)
	if err != nil {
		log.Println("insert problem", err)
		return err
	}
	return nil
}

func userBalanceSelect(db *sql.DB) error {
	rows, err := db.Query("SELECT * from user_balance")
	if err != nil {
		log.Println("zapros err", err)
	}
	var items []userBalance
	for rows.Next() {
		ub := userBalance{}
		err = rows.Scan(&ub.id, &ub.balance)
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
