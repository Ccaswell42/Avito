package handlers

import (
	"avito/storage/user_balance"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Data struct {
	DB   *sql.DB
	Port string
}

func (d *Data) Balance(w http.ResponseWriter, r *http.Request) {

	errStr := ValidateRequest(r, w, http.MethodGet)
	if errStr != OK {
		return
	}

	bal, errStr := ValidateBodyUserBalance(r.Body)
	if errStr != OK {
		JsonResponse(ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	ub, err := user_balance.GetBalance(d.DB, bal)
	if err != nil {
		JsonResponse(ResponseError, w, "no such user_id", http.StatusBadRequest)
		log.Println(err)
		return
	}
	JsonResponse(UB, w, ub, http.StatusOK)
}

func ValidateBodyUserBalance(r io.Reader) (user_balance.UserBalance, string) {
	var ub user_balance.UserBalance
	err := json.NewDecoder(r).Decode(&ub)

	if err != nil {
		log.Println(err)
		return ub, err.Error()
	}
	if ub.Id == 0 {
		return ub, UserBalanceZeroValue
	}
	return ub, OK
}
