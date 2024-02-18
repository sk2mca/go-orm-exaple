package dao

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	database "local.com/rest/db/orm/dbconnection"
	"local.com/rest/db/orm/modal"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var orders []modal.Order
	database.Db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]

	var order modal.Order
	database.Db.Preload("Items").First(&order, inputOrderID)
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order modal.Order
	json.NewDecoder(r.Body).Decode(&order)
	// Creates new order by inserting records in the `orders` and `items` table
	database.Db.Create(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	// Convert `orderId` string param to uint64
	id64, _ := strconv.ParseUint(inputOrderID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	database.Db.Where("order_id = ?", idToDelete).Delete(&modal.Item{})
	database.Db.Where("order_id = ?", idToDelete).Delete(&modal.Order{})
	w.WriteHeader(http.StatusNoContent)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder modal.Order
	var updateItems modal.Item

	json.NewDecoder(r.Body).Decode(&updatedOrder)

	database.Db.Updates(&updatedOrder)
	for _, item := range updatedOrder.Items {
		updateItems = item
		database.Db.Updates(&updateItems)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedOrder)
}
