package handlers

import (
	"avito/logic"
	"log"
	"net/http"
)

func (d *Data) Revenue(w http.ResponseWriter, r *http.Request) {

	errStr := logic.ValidateRequest(r, w, http.MethodPost)
	if errStr != logic.OK {
		return
	}

	ra, errStr := ValidateBodyReserve(r.Body)
	if errStr != logic.OK {
		logic.JsonResponse(logic.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	err := logic.Revenue(d.DB, ra)
	if err != nil {
		logic.JsonResponse(logic.ResponseError, w, "can't recognize avenue", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	logic.JsonResponse(logic.OK, w, "revenue successfully recognised", http.StatusOK)
}
