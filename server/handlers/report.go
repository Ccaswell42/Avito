package handlers

import (
	"avito/service"
	"avito/storage/accounting_report"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (d *Data) Report(w http.ResponseWriter, r *http.Request) {

	errStr := service.ValidateRequest(r, w, http.MethodGet)
	if errStr != service.OK {
		return
	}

	dateRep, errStr := ValidateBodyReport(r.Body)
	if errStr != service.OK {
		service.JsonResponse(service.ResponseError, w, errStr, http.StatusBadRequest)
		return
	}
	fileName, err := service.GetReportCsv(d.DB, dateRep)
	if err != nil {
		if err.Error() == "no entries in the report" {
			service.JsonResponse(service.ResponseError, w, err.Error(), http.StatusBadRequest)
		} else {
			service.JsonResponse(service.ResponseError, w, err.Error(), http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}
	service.JsonResponse(service.OK, w, "./"+fileName, http.StatusOK)

}

func ValidateBodyReport(r io.Reader) (accounting_report.DateReport, string) {
	var dateRep accounting_report.DateReport
	err := json.NewDecoder(r).Decode(&dateRep)

	if err != nil {
		log.Println(err)
		return dateRep, err.Error()
	}
	if dateRep.Year == "" || dateRep.Month == "" {
		return dateRep, service.ReportZeroValue
	}

	return dateRep, service.OK
}
