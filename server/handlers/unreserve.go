package handlers

import (
	"avito/service"
	"database/sql"
	"log"
	"net/http"
)

func (d *Data) UnReserve(w http.ResponseWriter, r *http.Request) {

	errStr := service.ValidateRequest(r, w, http.MethodPost)
	if errStr != service.OK {
		return
	}
	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != service.OK {
		service.JsonResponse(service.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}

	err := service.UnReserveMoney(d.DB, ra)
	if err != nil {
		if err == sql.ErrNoRows {
			service.JsonResponse(service.ResponseError, w, "can't unreserve money: no such note", http.StatusBadRequest)
		} else {
			service.JsonResponse(service.ResponseError, w, "can't unreserve money: "+err.Error(), http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	service.JsonResponse(service.OK, w, "money successfully unreserved", http.StatusOK)
}
