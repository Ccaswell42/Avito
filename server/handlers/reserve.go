package handlers

import (
	"avito/storage/reserve_account"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func (d *Data) Reserve(w http.ResponseWriter, r *http.Request) {

	errStr := ValidateRequest(r, w, http.MethodPost)
	if errStr != OK {
		return
	}

	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != OK {
		JsonResponse2(ResponseError, w, errStr, http.StatusBadRequest)
		return
	}

	err := reserve_account.ReserveMoney(d.DB, ra)
	if err != nil {
		if err == sql.ErrNoRows {
			JsonResponse2(ResponseError, w, "can't reserve money: no such user", http.StatusBadRequest)
		} else if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			JsonResponse2(ResponseError, w, "can't reserve money: order_id must be unique", http.StatusBadRequest)
		} else if strings.HasPrefix(err.Error(), "pq: null value in column") {
			JsonResponse2(ResponseError, w, "can't reserve money: not all request fields are specified", http.StatusBadRequest)
		} else {
			str := err.Error()
			JsonResponse2(ResponseError, w, "can't reserve money: "+str, http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	JsonResponse2(OK, w, "reserve OK", http.StatusOK)
}

func ValidateBodyReserve(r io.Reader) (reserve_account.ReverseAcc, string) {
	var ra reserve_account.ReverseAcc
	err := json.NewDecoder(r).Decode(&ra)

	if err != nil {
		log.Println(err)
		return ra, err.Error()
	}
	if ra.Id == 0 || ra.Service == 0 || ra.Cost == 0 || ra.OrderId == 0 {
		return ra, ReserveZeroValue
	}
	return ra, OK
}
