package handlers

import (
	"avito/logic"
	"avito/storage/accounting_report"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (d *Data) Report(w http.ResponseWriter, r *http.Request) {

	errStr := logic.ValidateRequest(r, w, http.MethodGet)
	if errStr != logic.OK {
		return
	}

	dateRep, errStr := ValidateBodyReport(r.Body)
	if errStr != logic.OK {
		logic.JsonResponse(logic.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	fileName, err := logic.GetReportCsv(d.DB, dateRep)
	if err != nil {
		logic.JsonResponse(logic.ResponseError, w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	logic.JsonResponse(logic.OK, w, "./"+fileName, http.StatusOK)

}

func ValidateBodyReport(r io.Reader) (accounting_report.DateReport, string) {
	var dateRep accounting_report.DateReport
	err := json.NewDecoder(r).Decode(&dateRep)

	if err != nil {
		log.Println(err)
		return dateRep, err.Error()
	}
	if dateRep.Year == "" || dateRep.Month == "" {
		return dateRep, logic.ReportZeroValue
	}

	return dateRep, logic.OK
}
