package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}
