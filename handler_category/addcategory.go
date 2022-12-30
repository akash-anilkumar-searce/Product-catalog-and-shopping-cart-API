package handler_category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
)

func Add_category(w http.ResponseWriter, r *http.Request) {
	var category typedefs.Category_master = typedefs.Category_master{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	fmt.Println("hello")
	json.Unmarshal(reqBody, &category)
	fmt.Println(category.Category_Id, category.Category_Name)
	//unmarshal the json values from postman to put into database
	db := dbconnect.ConnectToDB()
	stmt, err := db.Prepare("INSERT INTO category_master (category_id, category_name) VALUES($1,$2);")
	_, err = stmt.Exec(category.Category_Id, category.Category_Name)

	if err != nil {
		fmt.Println(err) //check here
	}

}
