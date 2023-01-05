package HandlerCart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	//add the required products into the cart using reference id generated
	urlQ := r.URL.Query()
	reference := urlQ.Get("ref")
	quantity_value := urlQ.Get("quantity")
	product_id := urlQ.Get("product_id")

	if reference == "" || quantity_value == "" || product_id == "" {
		Helpers.SendJResponse(response.ParameterMissing, w)
		return
	}

	quantity, err := strconv.Atoi(quantity_value)
	if err != nil {
		fmt.Println("string conv error", err)
	}

	rows, err := Helpers.QueryRun(queries.GetCartReference, reference)
	if err != nil {
		fmt.Println("query run error", err)
	}

	if !rows.Next() {
		response := "Invalid cart_reference"
		Helpers.SendJResponse(response, w)
		return
	}

	rows, err = Helpers.QueryRun(queries.JoinInventoryAndProductmaster, product_id)
	if err != nil {
		fmt.Println("query run error", err)
	}

	if !rows.Next() {
		Helpers.SendJResponse(response.ProductidInvalid, w)
		return
	}

	inventory_item := typedefs.Inventory{}
	err = rows.Scan(&inventory_item.Product_Id, &inventory_item.Quantity)
	if err != nil {
		fmt.Println("row scan error", err)
	}

	if inventory_item.Quantity-quantity < 0 {
		json.NewEncoder(w).Encode(map[string]string{"message": "Inventory Quantity is less than the required quantity: " + fmt.Sprint(inventory_item.Quantity)})
		return
	}

	_, err = Helpers.QueryRun(queries.UpdateInventory, inventory_item.Quantity-quantity, product_id)
	if err != nil {
		fmt.Println("run query error", err)
	}

	rows, err = Helpers.QueryRun(queries.GetQuantityFromCart, reference, product_id)
	if err != nil {
		fmt.Println("run query error", err)
	}

	if rows.Next() {
		var db_quantity int
		rows.Scan(&db_quantity)

		_, err = Helpers.QueryRun(queries.UpdateCartItem, db_quantity+quantity, reference, product_id)
		if err != nil {
			fmt.Println("run query error ", err)
		}

	} else {
		_, err = Helpers.QueryRun(queries.InsertIntoCart, reference, product_id, quantity)
		if err != nil {
			fmt.Println("run query error", err)
		}
	}

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error generated": err.Error()})
		return
	}
	Helpers.SendJResponse(response.ProductsAddedToCart, w)
}
