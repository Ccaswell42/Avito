package handlers

import (
	"avito/service"
	"avito/storage/reserve_account"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func (d *Data) Reserve(w http.ResponseWriter, r *http.Request) {

	errStr := service.ValidateRequest(r, w, http.MethodPost)
	if errStr != service.OK {
		return
	}

	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != service.OK {
		service.JsonResponse(service.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}

	err := service.ReserveMoney(d.DB, ra)
	if err != nil {
		if err == sql.ErrNoRows {
			service.JsonResponse(service.ResponseError, w, "can't reserve money: no such user", http.StatusBadRequest)
		} else if strings.HasPrefix(err.Error(), "pq: duplicate key value") {
			service.JsonResponse(service.ResponseError, w, "can't reserve money: order_id must be unique", http.StatusBadRequest)
		} else if strings.HasPrefix(err.Error(), "pq: null value in column") {
			service.JsonResponse(service.ResponseError, w, "can't reserve money: not all request fields are specified",
				http.StatusBadRequest)
		} else {
			str := err.Error()
			service.JsonResponse(service.ResponseError, w, "can't reserve money: "+str, http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	service.JsonResponse(service.OK, w, "reserve OK", http.StatusOK)
}

func ValidateBodyReserve(r io.Reader) (reserve_account.ReverseAcc, string) {
	var ra reserve_account.ReverseAcc
	err := json.NewDecoder(r).Decode(&ra)

	if err != nil {
		log.Println(err)
		return ra, err.Error()
	}
	if ra.Id == 0 || ra.Service == 0 || ra.Cost == 0 || ra.OrderId == 0 {
		return ra, service.ReserveZeroValue
	}
	return ra, service.OK
}
