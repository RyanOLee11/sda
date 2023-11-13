package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListProducts(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this is list product"))
}

func GetProduct(w http.ResponseWriter , r *http.Request){
	chi.URLParam(r, "id")
    w.Write([]byte("this is get product"))
}

func ListCustomers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this is the list customer"))
}

func GetCustomer(w http.ResponseWriter , r *http.Request){
	chi.URLParam(r, "id")
    w.Write([]byte("this is get customer"))
}
func ListOrders(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this is the list orders"))
}

func GetOrder(w http.ResponseWriter , r *http.Request){
	chi.URLParam(r, "id")
    w.Write([]byte("this is get order"))
}

