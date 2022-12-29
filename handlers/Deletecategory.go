package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/gorilla/mux"
)

func Deletecategory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	db := dbconnect.ConnectToDB()

	res, err := db.Exec("DELETE FROM category_master WHERE category_id=$1", x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				result := "The value category_id does not exist,enter a Valid ID"
				json.NewEncoder(w).Encode(result)
				//result := fmt.Sprint("The value category_id does not exist,enter a Valid ID")
				//json.NewEncoder(w).Encode(result)
				return
			}
			result := fmt.Sprint("The value is deleted successfully")
			json.NewEncoder(w).Encode(result)
		}

	}

	return

}
