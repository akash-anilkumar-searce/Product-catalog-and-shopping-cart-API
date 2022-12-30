package handler_product

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/typedefs"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]

	page_no, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("page ni invalid")
	}
	db := dbconnect.ConnectToDB()
	endlimit := page_no * 20
	startlimit := endlimit - 20
	fmt.Println(startlimit)
	rows, err := db.Query("Select name,price from product_master")
	defer rows.Close()
	products := []typedefs.Product_master{} // emptyarray
	for rows.Next() {
		new_product := typedefs.Product_master{} //content of the array products example it will take product of id1
		err := rows.Scan(&new_product.Name, &new_product.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, new_product) //add value into the array products
	}
	// response:=map[string]string{}
	min_len := int(math.Min(float64(len(products)), float64(endlimit)))
	if (startlimit < min_len) && startlimit >= 0 {
		fmt.Println(products[startlimit:min_len])
	}
	mapslist := []map[string]any{}
	for _, v := range products {
		prodnew := map[string]any{
			"name":  v.Name,
			"price": v.Price,
		}
		mapslist = append(mapslist, prodnew)

	}
	json.NewEncoder(w).Encode(mapslist)

}
