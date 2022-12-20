package handlers

import (
	"avito/service"
	"avito/storage/user_balance"
	"database/sql"
	"log"
	"net/http"
	"strconv"
)

type Data struct {
	DB   *sql.DB
	Port string
}

func (d *Data) Balance(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		service.JsonResponse(service.ResponseError, w, service.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	bal, errStr := ValidateBodyUserBalance(id)

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

func ValidateBodyUserBalance(id string) (user_balance.UserBalance, string) {
	var ub user_balance.UserBalance

	if id == "" {
		return ub, service.UserBalanceZeroValue
	}

	num, err := strconv.Atoi(id)
	ub.Id = num
	if ub.Id == 0 || err != nil {
		return ub, service.UserBalanceZeroValue
	}
	return ub, service.OK
}
