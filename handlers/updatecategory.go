package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
)

func Updatecategory(w http.ResponseWriter, r *http.Request) {
	db := dbconnect.ConnectToDB()
	var new_category typedefs.Category_master
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "error")
	}

	err = json.Unmarshal(reqBody, &new_category)
	fmt.Println(new_category)
	rows, err := db.Query("SELECT * from category_master WHERE category_id = $1", new_category.Category_Id)
	if err != nil {
		fmt.Println("error while selecting category")
	}
	defer rows.Close()
	var existing_category typedefs.Category_master
	for rows.Next() {
		fmt.Println("working")
		err := rows.Scan(&existing_category.Category_Id, &existing_category.Category_Name)
		if err != nil {
			fmt.Println("error while scanning")
		}

		if new_category.Category_Name == "" {
			new_category.Category_Name = existing_category.Category_Name
		}
		db.Query("UPDATE category_master SET category_name=$1 WHERE category_id =$2;", new_category.Category_Name, new_category.Category_Id)
		if err != nil {
			fmt.Println("error")
		}
	}
}
