package handler_category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
)

func Updatecategory(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	category := typedefs.Category_master{}

	err := json.Unmarshal(reqBody, &category)
	if err != nil {
		fmt.Println(err)
	}

	db := dbconnect.ConnectToDB()

	result, err := db.Exec("UPDATE category_master SET category_name=$1 WHERE category_id=$2;", category.Category_Name, category.Category_Id)
	if err != nil {
		fmt.Println("ERROR PRODUCED", err)
	}
	// check errors

	rows, err := result.RowsAffected()

	if rows != 1 {
		response := "category id does not exist"
		json.NewEncoder(w).Encode(response)
	} else {
		fmt.Println("Updating DB")
		fmt.Println("Updating category id:", category.Category_Id)
		response := "category details have been updated"
		json.NewEncoder(w).Encode(response)
	}

}

/*
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
		db.Exec("UPDATE category_master SET category_name=$1 WHERE category_id =$2;", new_category.Category_Name, new_category.Category_Id)

		if err != nil {
			fmt.Println("error")
		} else {
			json.NewEncoder(w).Encode(map[string]string{"message": "category detail has been updated"})
		}
	}

*/
