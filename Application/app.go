package application

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"database/sql"

	_ "github.com/lib/pq"
)

var Router *chi.Mux
var DB *sql.DB

func StartApplication() {
	connectionStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	DB = conn
	defer conn.Close()
	Router := chi.NewRouter()
	AdminRouter := chi.NewRouter()
	Router.Use(middleware.Logger)
	Router.Get("/", home)
	Router.Get("/orders", ListOrders)
	Router.Get("/customers", ListCustomers)
	Router.Get("/products", ListProducts)

	AdminRouter.Post("/order", CreateOrder)
	AdminRouter.Put("/order/{id}", UpdateOrder)
	AdminRouter.Get("/order/{id}", GetOrder)
	AdminRouter.Delete("/order/{id}", DeleteOrder)

	AdminRouter.Post("/customer", CreateCustomer)
	AdminRouter.Put("/customer/{id}", UpdateCustomer)
	AdminRouter.Get("/customer/{id}", GetCustomer)
	AdminRouter.Delete("/customer/{id}", DeleteCustomer)

	AdminRouter.Post("/product", CreateProduct)
	AdminRouter.Put("/product/{id}", UpdateProduct)
	AdminRouter.Get("/product/{id}", GetProduct)
	AdminRouter.Delete("/product/{id}", DeleteProduct)

	Router.Mount("/admin", AdminRouter)
	fmt.Println("app start")
	wd, _ := os.Getwd()

	fmt.Println(wd)


	http.ListenAndServe(":8080", Router)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("Views/Pages/Home.html")

	tmpl.Execute(w,"")
}
