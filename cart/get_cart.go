package cart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/handlers"
	"github.com/akash-searce/product-catalog/typedefs"
)

func GetCart(w http.ResponseWriter, r *http.Request) {

	//To retrieve the items present in the cart using the generated reference id
	urlQ := r.URL.Query()
	ref := urlQ.Get("ref")

	rows, err := handlers.QueryRun("SELECT * FROM cart_reference WHERE ref=$1", ref)
	if err != nil {
		fmt.Println("Query Run Error generated", err)
	}

	if !rows.Next() {
		json.NewEncoder(w).Encode((map[string]string{"ERROR GENERATED": "Cart Reference is Invalid"}))
		return
	}

	cart := typedefs.Cart{}
	err = rows.Scan(&cart.Ref, &cart.CreatedAt)
	if err != nil {
		fmt.Println("Error while scanning error", err)
	}

	rows, err = handlers.QueryRun("SELECT * FROM cart_item WHERE ref=$1", ref)
	if err != nil {
		fmt.Println("Query Run Error generated", err)
	}
	for rows.Next() {
		new_cart_item := typedefs.CartItem{}
		cart_ref_id := ""
		err = rows.Scan(&cart_ref_id, &new_cart_item.ProductID, &new_cart_item.Quantity)
		if err != nil {
			fmt.Println("Error while scanning row", err)
		}
		cart.Items = append(cart.Items, new_cart_item)
	}

	json.NewEncoder(w).Encode(cart)
}
