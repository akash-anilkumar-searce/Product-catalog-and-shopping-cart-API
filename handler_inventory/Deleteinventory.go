package handler_inventory

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/gorilla/mux"
)

func Deleteinventory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	db := dbconnect.ConnectToDB()

	res, err := db.Exec("DELETE FROM inventory WHERE product_id=$1", x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				result := fmt.Sprint("The value product_id does not exist, please enter a Valid ID")
				json.NewEncoder(w).Encode(result)
				return
			}
			result := fmt.Sprint("The value is deleted successfully")
			json.NewEncoder(w).Encode(result)
		}

	}

	return

}
