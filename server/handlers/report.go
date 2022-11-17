package handlers

import (
	"avito/storage/accounting_report"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (d *Data) Report(w http.ResponseWriter, r *http.Request) {

	errStr := ValidateRequest(r, w, http.MethodGet)
	if errStr != OK {
		return
	}

	dateRep, errStr := ValidateBodyReport(r.Body)
	if errStr != OK {
		JsonResponse(ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	err := accounting_report.GetReportCsv(d.DB, dateRep)
	if err != nil {
		JsonResponse(ResponseError, w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	JsonResponse(OK, w, "тестим", http.StatusOK)

}

func ValidateBodyReport(r io.Reader) (accounting_report.DateReport, string) {
	var dateRep accounting_report.DateReport
	err := json.NewDecoder(r).Decode(&dateRep)

	if err != nil {
		log.Println(err)
		return dateRep, err.Error()
	}
	if dateRep.Year == "" || dateRep.Month == "" {
		return dateRep, ReportZeroValue
	}

	return dateRep, OK
}
