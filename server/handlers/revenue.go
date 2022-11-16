package handlers

import (
	"avito/storage/accounting_report"
	"log"
	"net/http"
)

func (d *Data) Revenue(w http.ResponseWriter, r *http.Request) {

	errStr := ValidateRequest(r, w, http.MethodPost)
	if errStr != OK {
		return
	}

	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != OK {
		JsonResponse(ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	err := accounting_report.Revenue(d.DB, ra)
	if err != nil {
		JsonResponse(ResponseError, w, "can't recognize avenue", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	JsonResponse(OK, w, "revenue successfully recognised", http.StatusOK)
}
