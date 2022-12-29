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
	// handlers for product_master table
	r.HandleFunc("/addproduct", handlers.Add_product).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/getproducts/{id:[0-9]+}", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/updateproduct", handlers.Updateproduct).Methods("PUT")
	r.HandleFunc("/deleteproduct/{id:[0-9]+}", handlers.Deleteproduct).Methods("DELETE")
	// handlers for category_master table
	r.HandleFunc("/addcategory", handlers.Add_category).Methods("POST")
	r.HandleFunc("/getcategory/{id:[0-9]+}", handlers.Getcategory).Methods("GET")
	r.HandleFunc("/updatecategory", handlers.Updatecategory).Methods("PUT")
	r.HandleFunc("/deletecategory/{id:[0-9]+}", handlers.Deletecategory).Methods("DELETE")
	//handlers for inventory table
	r.HandleFunc("/addinventory", handlers.Add_into_inventory).Methods("POST")
	r.HandleFunc("/inventorydetail/{id:[0-9]+}", handlers.Getinventory).Methods("GET")
	r.HandleFunc("/updateinventory", handlers.Updateinventory).Methods("PUT")
	r.HandleFunc("/deleteinventory/{id:[0-9]+}", handlers.Deleteinventory).Methods("DELETE")
	// handlers for cart table
	r.HandleFunc("/cart/createreference", cart.CreateCart).Methods("POST")
	r.HandleFunc("/addtocart", cart.AddToCart).Methods("POST")
	r.HandleFunc("/cartitems/add", cart.AddItemsToCart).Methods("POST")
	r.HandleFunc("/cart/get", cart.GetCart1).Methods("GET")
	r.HandleFunc("/deletefromcart", cart.RemoveItemFromCart).Methods("DELETE")

	fmt.Printf("Starting server at port 8089\n")
	log.Fatal(http.ListenAndServe(":8089", r))
}
