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
}
