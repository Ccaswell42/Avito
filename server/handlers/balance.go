package handlers

import (
	"avito/storage/user_balance"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	DB *sql.DB
}

type ErrResponse struct {
	Error interface{} `json:"Error"`
}

type ResultResponse struct {
	Result interface{} `json:"Result"`
}

func JsonResponse(ok bool, w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if !ok {
		errResp := ErrResponse{data}
		err := json.NewEncoder(w).Encode(errResp)
		if err != nil {
			log.Println("encode error", err)
			http.Error(w, "encode error", 500)
		}
	} else {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println("encode error", err)
			http.Error(w, "encode error", 500)
		}
	}
}

func (d *Data) Balance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		JsonResponse(false, w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var bal user_balance.UserBalance
	err := json.NewDecoder(r.Body).Decode(&bal)

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
	ub, err := user_balance.GetBalance(d.DB, bal)
	if err != nil {
		JsonResponse(false, w, "no such user_id", http.StatusOK)
		log.Println(err)
		return
	}
	JsonResponse(true, w, ub, http.StatusOK)
}
