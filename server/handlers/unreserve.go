package handlers

import (
	"avito/storage/reserve_account"
	"encoding/json"
	"log"
	"net/http"
)

func (d *Data) UnReserve(w http.ResponseWriter, r *http.Request) {
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
	err = reserve_account.UnReserveMoney(d.DB, ra)
	if err != nil {
		JsonResponse(false, w, "can't unreserve money", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	JsonResponse(true, w, "money successfully unreserved", http.StatusOK)

}
