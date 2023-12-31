package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/go-chi/chi/v5"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newProd Product
	err := decoder.Decode(&newProd)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	var id int
	err = DB.QueryRow(
		"INSERT INTO products(name, price, description) VALUES($1, $2, $3) RETURNING id",
		newProd.Name, newProd.Price, newProd.Description).Scan(&id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write([]byte("this is create a product"))
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Prod Product
	err := decoder.Decode(&Prod)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	fmt.Println(id)
	_, err = DB.Exec("update products set name = $1, price = $2, description = $3 where id = $4",
		Prod.Name, Prod.Price, Prod.Description, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write([]byte("this is update a product"))
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := DB.Exec("delete from products where id = $1", id)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	w.Write([]byte("this is delete a product"))
}
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newOrder Order
	err := decoder.Decode(&newOrder)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	var id int
	err = DB.QueryRow("insert into orders (customer_id, order_date, order_total) values($1, $2, $3) returning ID", newOrder.CustomerID, newOrder.OrderDate, newOrder.OrderTotal).Scan(&id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write([]byte("this is create a order"))
}
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ord Order
	err := decoder.Decode(&ord)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err = DB.Exec("update orders set customer_id = $1 , order_date = $2, order_total = $3 where id = $4", ord.CustomerID, ord.OrderDate, ord.OrderTotal, id)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write([]byte("this is update a order"))
}
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := DB.Exec("delete from orders where id = $1", id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write([]byte("this is delete a order"))
}
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newCustomer Customer
	err := decoder.Decode(&newCustomer)
	if err != nil {
		// Log the error and return a 500 Internal Server Error response
		log.Printf("Error deleting order: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var id int
	err = DB.QueryRow(
		"INSERT INTO customers(first_name, last_name, state, address) VALUES($1, $2, $3, $4) RETURNING id",
		newCustomer.FirstName, newCustomer.LastName, newCustomer.State, newCustomer.Address).Scan(&id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write([]byte("this is create a customer"))
}
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var Cust Customer
	err := decoder.Decode(&Cust)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	fmt.Println(id)
	_, err = DB.Exec("update customers set first_name = $1, last_name = $2, state = $3, address = $4 where id = $5",
		Cust.FirstName, Cust.LastName, Cust.State, Cust.Address, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	w.Write([]byte("this is update a customer"))
}
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := DB.Exec("delete from customers where id = $1", id)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	w.Write([]byte("this is delete a customer"))
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
type Customer struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	State     string `json:"state"`
	Address   string `json:"address"`
}
type Order struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	OrderDate  string `json:"order_date"`
	OrderTotal int    `json:"order_total"`
	OrderItems []OrderItem
}

type OrderItem struct {
	ID        int    `json:"id"`
	OrderID   int    `json:"order_id"`
	ProductID int    `json:"product_id"`
	Price     string `json:"price"`
	Name      string `json:"name"`
	Qty       int    `json:"qty"`
}
