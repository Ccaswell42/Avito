package handlers

import (
	"encoding/json"
	"io"
	"log"
)

type DateReport struct {
	Year  int `json:"Year"`
	Month int `json:"Month"`
}

//func (d *Data) Report(w http.ResponseWriter, r *http.Request) {
//
//	errStr := ValidateRequest(r, w, http.MethodGet)
//	if errStr != OK {
//		return
//	}
//
//	dateRep, errStr := ValidateBodyReport(r.Body)
//	if errStr != OK {
//		JsonResponse(ResponseError, w, errStr, http.StatusBadRequest)
//		return
//	}
//
//}

func ValidateBodyReport(r io.Reader) (DateReport, string) {
	var dateRep DateReport
	err := json.NewDecoder(r).Decode(&dateRep)

	if err != nil {
		log.Println(err)
		return dateRep, err.Error()
	}
	if dateRep.Year == 0 || dateRep.Month == 0 {
		return dateRep, ReportZeroValue
	}

	return dateRep, OK
}
