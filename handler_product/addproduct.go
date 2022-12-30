package handler_product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
)

func Add_product(w http.ResponseWriter, r *http.Request) {
	var product typedefs.Product_master = typedefs.Product_master{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &product)
	fmt.Println(product.Product_Id)
	db := dbconnect.ConnectToDB()
	stmt, err := db.Prepare("INSERT INTO product_master (product_id, name, specification,sku,category_id, price) VALUES($1,$2,$3,$4,$5,$6);")
	spec_json_byte, err := json.Marshal(product.Specification)
	_, err = stmt.Exec(product.Product_Id, product.Name, spec_json_byte, product.SKU, product.Category_Id, product.Price)

	if err != nil {
		fmt.Println(err) //check here
	}

}
