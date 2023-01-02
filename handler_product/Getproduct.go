package handler_product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	Params := mux.Vars(r)
	var product_spec string
	product := typedefs.Product_master{}
	db := dbconnect.ConnectToDB()
	ID := Params["id"]

	query := "SELECT p.product_id,p.name,p.specification,p.sku,c.category_id, p.price FROM product_master p JOIN category_master c ON p.category_id =c.category_id where product_id=$1"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(ID)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		rows.Scan(&product.Product_Id, &product.Name, &product_spec, &product.SKU, &product.Category_Id, &product.Price)
		json.Unmarshal([]byte(product_spec), &product.Specification)
		json.NewEncoder(w).Encode(product)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "Product does not exist"})
	}

}
