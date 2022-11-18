package handlers

import (
	"avito/service"
	"avito/storage/user_balance"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (d *Data) Replenish(w http.ResponseWriter, r *http.Request) {

	errStr := service.ValidateRequest(r, w, http.MethodPost)
	if errStr != service.OK {
		return
	}

	bal, errStr := ValidateReplenish(r.Body)
	if errStr != service.OK {
		service.JsonResponse(service.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}

	err := service.ReplenishBalance(d.DB, bal)
	if err != nil {
		service.JsonResponse(service.ResponseError, w, "can't replenish the balance",
			http.StatusInternalServerError)
		log.Println(err)
		return
	}
	service.JsonResponse(service.OK, w, "balance successfully replenished", http.StatusOK)
}

func ValidateReplenish(r io.Reader) (user_balance.UserBalance, string) {
	var ub user_balance.UserBalance
	err := json.NewDecoder(r).Decode(&ub)

	if err != nil {
		log.Println(err)
		return ub, err.Error()
	}
	if ub.Id == 0 || ub.Balance <= 0 {
		return ub, service.ReserveZeroValue
	}
	return ub, service.OK
}
