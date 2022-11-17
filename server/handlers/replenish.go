package handlers

import (
	"avito/logic"
	"log"
	"net/http"
)

func (d *Data) Replenish(w http.ResponseWriter, r *http.Request) {

	errStr := logic.ValidateRequest(r, w, http.MethodPost)
	if errStr != logic.OK {
		return
	}

	bal, errStr := ValidateBodyUserBalance(r.Body)
	if errStr != logic.OK {
		logic.JsonResponse(logic.ResponseError, w, errStr, http.StatusOK)
		return
	}

	err := logic.ReplenishBalance(d.DB, bal)
	if err != nil {
		logic.JsonResponse(logic.ResponseError, w, "can't replenish the balance",
			http.StatusInternalServerError)
		log.Println(err)
		return
	}
	logic.JsonResponse(logic.OK, w, "balance successfully replenished", http.StatusOK)
}
