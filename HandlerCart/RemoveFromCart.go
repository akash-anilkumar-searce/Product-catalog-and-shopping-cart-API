package HandlerCart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
)

func RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	ref := urlQuery.Get("ref")
	product_id := urlQuery.Get("product_id")
	if ref == "" || product_id == "" {
		Helpers.SendJResponse(response.ReferenceOrProductidMissing, w)
		return
	}

	rows, err := Helpers.QueryRun(queries.GetFromCartReference, ref)
	if err != nil {
		fmt.Println("query run error found", err)
	}

	if !rows.Next() {
		Helpers.SendJResponse(response.InvalidCartRef, w)
		return
	}

	result, err := DbConnect.ConnectToDB().Exec(queries.DeleteFromCart, ref, product_id)
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
			Helpers.SendJResponse(response.CartItemDeleted, w)
		}
	} else {
		Helpers.SendJResponse(response.ProductNotInCart, w)
	}
}