package reserve_account

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type ReverseAcc struct {
	Id      int  `json:"Id"`
	Service int  `json:"Service"`
	OrderId int  `json:"OrderId"`
	Cost    uint `json:"Cost"`
}

func ReverseAccInsert(db *sql.DB, ra ReverseAcc) error {
	_, err := db.Exec(
		"INSERT INTO reserve_account (id, service, order_id, cost) VALUES ($1, $2, $3, $4 )",
		ra.Id, ra.Service, ra.OrderId, ra.Cost)
	if err != nil {
		log.Println("insert error: ", err)
		return err
	}
	return nil
}

func ReserveAccSelect(db *sql.DB, ra ReverseAcc) (ReverseAcc, error) {

	var newRa ReverseAcc

	row := db.QueryRow("SELECT id, service, order_id, cost FROM reserve_account WHERE id = $1 AND service = $2 "+
		"AND order_id = $3 AND cost = $4",
		ra.Id, ra.Service, ra.OrderId, ra.Cost)

	err := row.Scan(&newRa.Id, &newRa.Service, &newRa.OrderId, &newRa.Cost)
	fmt.Println(newRa)
	if err != nil {
		log.Println("scan error: ", err)
		return newRa, err
	}
	return newRa, nil
}

func DeleteRow(db *sql.DB, ra ReverseAcc) error {
	_, err := db.Exec("DELETE from reserve_account WHERE id = $1 AND service = $2 AND order_id = $3 AND cost = $4",
		ra.Id, ra.Service, ra.OrderId, ra.Cost)

	if err != nil {
		log.Println("delete error :", err)
		return err
	}
	return nil
}
