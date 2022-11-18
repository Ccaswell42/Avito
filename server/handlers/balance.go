package handlers

import (
	"avito/service"
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

	errStr := service.ValidateRequest(r, w, http.MethodGet)
	if errStr != service.OK {
		return
	}

	bal, errStr := ValidateBodyUserBalance(r.Body)
	if errStr != service.OK {
		service.JsonResponse(service.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	ub, err := user_balance.GetBalance(d.DB, bal)
	if err != nil {
		service.JsonResponse(service.ResponseError, w, "no such user_id", http.StatusBadRequest)
		log.Println(err)
		return
	}
	service.JsonResponse(service.UB, w, ub, http.StatusOK)
}

func ValidateBodyUserBalance(r io.Reader) (user_balance.UserBalance, string) {
	var ub user_balance.UserBalance
	err := json.NewDecoder(r).Decode(&ub)

	if err != nil {
		log.Println(err)
		return ub, err.Error()
	}
	if ub.Id == 0 {
		return ub, service.UserBalanceZeroValue
	}
	return ub, service.OK
}
