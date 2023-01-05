package HandlerProduct

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	Params := mux.Vars(r)
	var product_spec string
	product := typedefs.Product_master{}
	db := DbConnect.ConnectToDB()
	ID := Params["id"]

	stmt, err := db.Prepare(queries.GetProduct)
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
		Helpers.SendJResponse(response.ProductNotFound, w)
	}

}
