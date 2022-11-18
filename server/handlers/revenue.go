package handlers

import (
	"avito/service"
	"log"
	"net/http"
)

func (d *Data) Revenue(w http.ResponseWriter, r *http.Request) {

	errStr := service.ValidateRequest(r, w, http.MethodPost)
	if errStr != service.OK {
		return
	}

	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != service.OK {
		service.JsonResponse(service.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	err := service.Revenue(d.DB, ra)
	if err != nil {
		service.JsonResponse(service.ResponseError, w, "can't recognize revenue", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	service.JsonResponse(service.OK, w, "revenue successfully recognised", http.StatusOK)
}
