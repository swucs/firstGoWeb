package handlers

import (
	"banking/service"
	"banking/utils"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return CustomerHandler{service}
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer("")

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, appErr := ch.service.GetCustomer(id)
	if appErr != nil {
		utils.WriteResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		utils.WriteResponse(w, http.StatusOK, customer)
	}
}
