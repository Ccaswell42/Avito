package handlers

import (
	"avito/storage/reserve_account"
	"database/sql"
	"log"
	"net/http"
)

func (d *Data) UnReserve(w http.ResponseWriter, r *http.Request) {

	errStr := ValidateRequest(r, w, http.MethodPost)
	if errStr != OK {
		return
	}
	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != OK {
		JsonResponse(ResponseError, w, errStr, http.StatusBadRequest)
		return
	}

	err := reserve_account.UnReserveMoney(d.DB, ra)
	if err != nil {
		if err == sql.ErrNoRows {
			JsonResponse(ResponseError, w, "can't unreserve money: no such note", http.StatusBadRequest)
		} else {
			JsonResponse(ResponseError, w, "can't unreserve money: "+err.Error(), http.StatusInternalServerError)
		}

		log.Println(err)
		return
	}
	JsonResponse(OK, w, "money successfully unreserved", http.StatusOK)

}
