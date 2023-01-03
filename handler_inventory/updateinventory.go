package handler_inventory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
)

func Updateinventory(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	inventory := typedefs.Inventory{}

	err := json.Unmarshal(reqBody, &inventory)
	if err != nil {
		fmt.Println(err)
	}

	db := dbconnect.ConnectToDB()

	result, err := db.Exec("UPDATE inventory SET quantity=$1 WHERE product_Id=$2;", inventory.Quantity, inventory.Product_Id)
	if err != nil {
		fmt.Println("ERROR FOUND", err)
	}
	// check errors

	rows, err := result.RowsAffected()

	if rows != 1 {
		response := "Inventory id doesn't exist"
		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Println("Updating product id:", inventory.Product_Id)
		response := "Inventory detail has been  has been updated successfully!"
		json.NewEncoder(w).Encode(response)
	}

}

/*
	db := dbconnect.ConnectToDB()
	var inventory_update typedefs.Inventory
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error")
	}

	err = json.Unmarshal(reqBody, &inventory_update)
	fmt.Println(inventory_update)
	rows, err := db.Query("SELECT * from inventory WHERE product_id = $1", inventory_update.Product_Id)
	if err != nil {
		fmt.Println("error while selecting inventory details")
	}
	defer rows.Close()
	var existing_inventory typedefs.Inventory
	for rows.Next() {
		fmt.Println("working")
		err := rows.Scan(&existing_inventory.Product_Id, &existing_inventory.Quantity)
		if err != nil {
			fmt.Println("error while scanning")
		}

		if inventory_update.Quantity == 0 {
			inventory_update.Quantity = existing_inventory.Quantity
		}
		db.Query("UPDATE inventory SET quantity=$1 WHERE product_id =$2;", inventory_update.Quantity, inventory_update.Product_Id)
		if err != nil {
			fmt.Println("error")
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "product detail has been updated"})
*/
