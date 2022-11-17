package handlers

import (
	"avito/logic"
	"avito/storage/reserve_account"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func (d *Data) Reserve(w http.ResponseWriter, r *http.Request) {

	errStr := logic.ValidateRequest(r, w, http.MethodPost)
	if errStr != logic.OK {
		return
	}

	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != logic.OK {
		logic.JsonResponse(logic.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}

	err := logic.ReserveMoney(d.DB, ra)
	if err != nil {
		if err == sql.ErrNoRows {
			logic.JsonResponse(logic.ResponseError, w, "can't reserve money: no such user", http.StatusBadRequest)
		} else if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			logic.JsonResponse(logic.ResponseError, w, "can't reserve money: order_id must be unique", http.StatusBadRequest)
		} else if strings.HasPrefix(err.Error(), "pq: null value in column") {
			logic.JsonResponse(logic.ResponseError, w, "can't reserve money: not all request fields are specified",
				http.StatusBadRequest)
		} else {
			str := err.Error()
			logic.JsonResponse(logic.ResponseError, w, "can't reserve money: "+str, http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	logic.JsonResponse(logic.OK, w, "reserve OK", http.StatusOK)
}

func ValidateBodyReserve(r io.Reader) (reserve_account.ReverseAcc, string) {
	var ra reserve_account.ReverseAcc
	err := json.NewDecoder(r).Decode(&ra)

	if err != nil {
		log.Println(err)
		return ra, err.Error()
	}
	if ra.Id == 0 || ra.Service == 0 || ra.Cost == 0 || ra.OrderId == 0 {
		return ra, logic.ReserveZeroValue
	}
	return ra, logic.OK
}
