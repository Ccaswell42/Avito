package service

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	OK                   string = "OK"
	ResponseError        string = "errResponse"
	MethodNotAllowed     string = "invalid method"
	UnsupportedMediaType string = "content-type is not application/json"
	UserBalanceZeroValue string = "no user_id or user_id is nul"
	UB                   string = "userBalance"
	ReserveZeroValue     string = "request does not contain all fields or null values"
	ReportZeroValue      string = "no month or year specified"
)

type ErrResponse struct {
	Error interface{} `json:"Error"`
}

type ResultResponse struct {
	Result interface{} `json:"Result"`
}

func ValidateRequest(r *http.Request, w http.ResponseWriter, method string) string {
	if r.Method != method {
		JsonResponse(ResponseError, w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return MethodNotAllowed
	}
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		JsonResponse(ResponseError, w, UnsupportedMediaType, http.StatusUnsupportedMediaType)
		log.Println(contentType)
		return UnsupportedMediaType
	}
	return OK
}

func JsonResponse(ok string, w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	switch ok {
	case OK:
		resultResp := ResultResponse{data}
		err := json.NewEncoder(w).Encode(resultResp)
		if err != nil {
			log.Println("encode error", err)
			http.Error(w, "encode error", 500)
		}
	case ResponseError:
		errResp := ErrResponse{data}
		err := json.NewEncoder(w).Encode(errResp)
		if err != nil {
			log.Println("encode error", err)
			http.Error(w, "encode error", 500)
		}
	case UB:
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println("encode error", err)
			http.Error(w, "encode error", 500)
		}
	}
}
