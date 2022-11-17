package handlers

import (
	"avito/logic"
	"database/sql"
	"log"
	"net/http"
)

func (d *Data) UnReserve(w http.ResponseWriter, r *http.Request) {

	errStr := logic.ValidateRequest(r, w, http.MethodPost)
	if errStr != logic.OK {
		return
	}
	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != logic.OK {
		logic.JsonResponse(logic.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}

	err := logic.UnReserveMoney(d.DB, ra)
	if err != nil {
		if err == sql.ErrNoRows {
			logic.JsonResponse(logic.ResponseError, w, "can't unreserve money: no such note", http.StatusBadRequest)
		} else {
			logic.JsonResponse(logic.ResponseError, w, "can't unreserve money: "+err.Error(), http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	logic.JsonResponse(logic.OK, w, "money successfully unreserved", http.StatusOK)
}
