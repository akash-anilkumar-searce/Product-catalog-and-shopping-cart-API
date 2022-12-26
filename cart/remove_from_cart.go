package cart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/dbconnect"
	"github.com/akash-searce/product-catalog/handlers"
)

func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")
	product_id := urlQuery.Get("product_id")

	if ref == "" || product_id == "" {
		json.NewEncoder(w).Encode(map[string]string{"message": "ref / product_id missing in the url"})
		return
	}

	rows, err := handlers.QueryRun("SELECT * FROM cart_reference WHERE ref=$1;", ref)
	if err != nil {
		fmt.Println("query run error found", err)
	}

	if !rows.Next() {
		json.NewEncoder(w).Encode(map[string]string{"message": "Invalid cart_reference"})
		return
	}

	result, err := dbconnect.ConnectToDB().Exec("DELETE FROM cart_item WHERE ref=$1 AND product_id=$2", ref, product_id)
	if err != nil {
		fmt.Println("query run error has occured", err)
	}

	rows_affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("row affected error has occured")
	}

	if rows_affected != 0 {
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"message": "Cart item deleted successfully"})
		}
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "Product is not found in your cart"})
	}
}
