package handlers

import (
	"avito/storage/reserve_account"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func (d *Data) Reserve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	var ra reserve_account.ReverseAcc
	err := json.NewDecoder(r.Body).Decode(&ra)

	if err != nil {
		JsonResponse(false, w, "no request body", http.StatusBadRequest)
		log.Println(err)
		return
	}
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		JsonResponse(false, w, "content-type is not application/json", http.StatusUnsupportedMediaType)
		log.Println(contentType)
		return
	}

	err = reserve_account.ReserveMoney(d.DB, ra)
	if err != nil {
		if err == sql.ErrNoRows {
			JsonResponse(false, w, "can't reserve money: no such user", http.StatusBadRequest)
		} else {
			str := err.Error()
			JsonResponse(false, w, "can't reserve money: "+str, http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	JsonResponse(true, w, "reserve OK", http.StatusOK)
}
