package handler_inventory

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func Getinventory(w http.ResponseWriter, r *http.Request) {
	Params := mux.Vars(r)
	inventory := typedefs.Inventory{}
	db := dbconnect.ConnectToDB()
	ID := Params["id"]

	stmt, err := db.Prepare("SELECT * from inventory where product_id=$1")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(ID)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		rows.Scan(&inventory.Product_Id, &inventory.Quantity)
		json.NewEncoder(w).Encode(inventory)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "Product does not exist"})
	}

}
