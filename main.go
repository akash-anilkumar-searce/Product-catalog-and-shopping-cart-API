package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akash-searce/product-catalog/handler_cart"
	"github.com/akash-searce/product-catalog/handler_category"
	"github.com/akash-searce/product-catalog/handler_inventory"
	"github.com/akash-searce/product-catalog/handler_product"
	"github.com/akash-searce/product-catalog/t_console_interface"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// to check if customer wants to use console interface or not
	f := func() {
		fmt.Println("If you want to start console interface? (yes or no)")
		var reply string
		_, err := fmt.Scanln(&reply)
		if err != nil {
			fmt.Println("error in reading input!!")
		}
		if reply == "yes" {
			t_console_interface.Console()
		} else if reply == "no" {
			fmt.Println("Console Interface cancelled")
		}
	}

	go f()

	r := mux.NewRouter()
	// handlers for product_master table
	r.HandleFunc("/addproduct", handler_product.Add_product).Methods("POST")
	r.HandleFunc("/product/{id:[0-9]+}", handler_product.GetProduct).Methods("GET")
	r.HandleFunc("/getproducts/{id:[0-9]+}", handler_product.GetProducts).Methods("GET")
	r.HandleFunc("/updateproduct", handler_product.Updateproduct).Methods("PUT")
	r.HandleFunc("/deleteproduct/{id:[0-9]+}", handler_product.Deleteproduct).Methods("DELETE")
	// handlers for category_master table
	r.HandleFunc("/addcategory", handler_category.Add_category).Methods("POST")
	r.HandleFunc("/getcategory/{id:[0-9]+}", handler_category.Getcategory).Methods("GET")
	r.HandleFunc("/updatecategory", handler_category.Updatecategory).Methods("PUT")
	r.HandleFunc("/deletecategory/{id:[0-9]+}", handler_category.Deletecategory).Methods("DELETE")
	//handlers for inventory table
	r.HandleFunc("/addinventory", handler_inventory.Add_into_inventory).Methods("POST")
	r.HandleFunc("/inventorydetail/{id:[0-9]+}", handler_inventory.Getinventory).Methods("GET")
	r.HandleFunc("/updateinventory", handler_inventory.Updateinventory).Methods("PUT")
	r.HandleFunc("/deleteinventory/{id:[0-9]+}", handler_inventory.Deleteinventory).Methods("DELETE")
	// handlers for cart table
	r.HandleFunc("/cart/createreference", handler_cart.CreateCart).Methods("POST")
	r.HandleFunc("/addtocart", handler_cart.AddToCart).Methods("POST")
	r.HandleFunc("/cartitems/add", handler_cart.AddItemsToCart).Methods("POST")
	r.HandleFunc("/cart/get", handler_cart.GetCart1).Methods("GET")
	r.HandleFunc("/deletefromcart", handler_cart.RemoveItemFromCart).Methods("DELETE")

	fmt.Printf("Starting server at port 8089\n")
	log.Fatal(http.ListenAndServe(":8089", r))
}
