package HandlerCategory

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	/*Params := mux.Vars(r)
	db := dbconnect.ConnectToDB()
	ID := Params["id"]
	*/

	category := typedefs.Category_master{}
	getcategory_id := mux.Vars(r)["id"]

	rows, err := Helpers.QueryRun(queries.GetCategory, getcategory_id)
	if err != nil {
		fmt.Println("runQueryError", err)
	}
	if rows.Next() {
		err := rows.Scan(&category.Category_Id, &category.Category_Name)
		json.NewEncoder(w).Encode(category)
		if err != nil {
			fmt.Println("rowscanerror", err)
		}
	} else {
		Helpers.SendJResponse(response.CategoryidNotPresent, w)
	}

	/*

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
	*/

}
