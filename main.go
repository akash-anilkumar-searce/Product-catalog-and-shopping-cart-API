package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akash-searce/product-catalog/cart"
	"github.com/akash-searce/product-catalog/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/addproduct", handlers.Add_product).Methods("POST")
	r.HandleFunc("/addcategory", handlers.Add_category).Methods("POST")
	r.HandleFunc("/addinventory", handlers.Add_into_inventory).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/getproducts/{id:[0-9]+}", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/getcategory/{id:[0-9]+}", handlers.Getcategory).Methods("GET")
	r.HandleFunc("/inventorydetail/{id:[0-9]+}", handlers.Getinventory).Methods("GET")
	r.HandleFunc("/deleteproduct/{id:[0-9]+}", handlers.Deleteproduct).Methods("DELETE")
	r.HandleFunc("/deletecategory/{id:[0-9]+}", handlers.Deletecategory).Methods("DELETE")
	r.HandleFunc("/deleteinventory/{id:[0-9]+}", handlers.Deleteinventory).Methods("DELETE")
	r.HandleFunc("/updatecategory", handlers.Updatecategory).Methods("PUT")
	r.HandleFunc("/updateproduct", handlers.Updateproduct).Methods("PUT")
	r.HandleFunc("/updateinventory", handlers.Updateinventory).Methods("PUT")

	r.HandleFunc("/cart/get", cart.GetCart).Methods("GET")
	r.HandleFunc("/cart/createreference", cart.CreateCart).Methods("POST")
	r.HandleFunc("/addtocart", cart.AddToCart).Methods("POST")
	r.HandleFunc("/deleteitemfromcart", cart.RemoveItemFromCart).Methods("DELETE")
	fmt.Printf("Starting server at port 8089\n")
	log.Fatal(http.ListenAndServe(":8089", r))
}
