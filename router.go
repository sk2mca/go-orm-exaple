package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"local.com/rest/db/orm/dao"
)

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", dao.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{orderId}", dao.GetOrder).Methods("GET")
	router.HandleFunc("/orders", dao.GetOrders).Methods("GET")
	router.HandleFunc("/orders", dao.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", dao.DeleteOrder).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}
