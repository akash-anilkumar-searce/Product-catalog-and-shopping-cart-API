package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func Getcategory(w http.ResponseWriter, r *http.Request) {
	Params := mux.Vars(r)
	category := typedefs.Category_master{}
	db := dbconnect.ConnectToDB()
	ID := Params["id"]

	stmt, err := db.Prepare("SELECT category_id,category_name from category_master where category_id=$1")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(ID)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		rows.Scan(&category.Category_Id, &category.Category_Name)
		json.NewEncoder(w).Encode(category)
	} else {
		response := "category id not present"
		json.NewEncoder(w).Encode(response)
	}

}
