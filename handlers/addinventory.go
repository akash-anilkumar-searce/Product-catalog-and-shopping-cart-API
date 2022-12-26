package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
)

func Add_into_inventory(w http.ResponseWriter, r *http.Request) {
	var inventory typedefs.Inventory
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &inventory)
	db := dbconnect.ConnectToDB()
	stmt, err := db.Prepare("INSERT INTO inventory (product_id, quantity) VALUES($1,$2);")
	_, err = stmt.Exec(inventory.Product_Id, inventory.Quantity)
	fmt.Println(inventory)

	if err != nil {
		fmt.Println(err) //check here
	}

}
