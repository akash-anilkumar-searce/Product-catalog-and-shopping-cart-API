package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/gorilla/mux"
)

func Deleteproduct(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	db := dbconnect.ConnectToDB()

	res, err := db.Exec("DELETE FROM product_master WHERE product_id=$1", x)

	if err == nil {

		count, err := res.RowsAffected() // by default rows affected would return count and error.
		if err == nil {
			if count == 0 {
				result := fmt.Sprint("The value of Product_Id Does not exist,enter a Valid ID")
				json.NewEncoder(w).Encode(result)
				return // if code hits return then it wont run the code written after it
			}
			result := fmt.Sprint("The value is deleted sucessfully")
			json.NewEncoder(w).Encode(result)
		}

	}

	return

}

/*Params := mux.Vars(r)
	db := dbconnect.ConnectToDB()
	ID := Params["id"]
	stmt, err := db.Prepare("DELETE FROM product_master where product_id=$1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stmt)
	if err != nil {
		fmt.Fprintf(w, "Error")
	} else {
		fmt.Println("successfully deleted ")
	}
	_, err = stmt.Query(ID)
	defer stmt.Close()

	if err != nil {
		fmt.Fprintf(w, "Error")
	}

}
*/
