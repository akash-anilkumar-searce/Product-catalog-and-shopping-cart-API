package handler_cart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akash-searce/product-catalog/helpers"
	"github.com/akash-searce/product-catalog/typedefs"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	//add the required products into the cart using reference id generated
	urlQ := r.URL.Query()
	reference := urlQ.Get("ref")
	quantity_value := urlQ.Get("quantity")
	product_id := urlQ.Get("product_id")

	if reference == "" || quantity_value == "" || product_id == "" {
		json.NewEncoder(w).Encode(map[string]string{"ALERT": "one of the required parameter is missing"})
		return
	}

	quantity, err := strconv.Atoi(quantity_value)
	if err != nil {
		fmt.Println("string conv error", err)
	}

	rows, err := helpers.QueryRun("SELECT * FROM cart_reference WHERE ref=$1;", reference)
	if err != nil {
		fmt.Println("query run error", err)
	}

	if !rows.Next() {
		json.NewEncoder(w).Encode((map[string]string{"message": "Invalid cart_reference"}))
		return
	}

	rows, err = helpers.QueryRun("SELECT p.product_id, i.quantity FROM product_master p LEFT JOIN inventory i ON p.product_id=i.product_id WHERE p.product_id=$1", product_id)
	if err != nil {
		fmt.Println("query run error", err)
	}

	if !rows.Next() {
		json.NewEncoder(w).Encode(map[string]string{"message": "Product id is invalid"})
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

	_, err = helpers.QueryRun("UPDATE inventory SET quantity=$1 WHERE product_id=$2", inventory_item.Quantity-quantity, product_id)
	if err != nil {
		fmt.Println("run query error", err)
	}

	rows, err = helpers.QueryRun("SELECT quantity FROM cart_item WHERE ref=$1 AND product_id=$2", reference, product_id)
	if err != nil {
		fmt.Println("run query error", err)
	}

	if rows.Next() {
		var db_quantity int
		rows.Scan(&db_quantity)

		_, err = helpers.QueryRun("UPDATE cart_item SET quantity=$1 WHERE ref=$2 AND product_id=$3", db_quantity+quantity, reference, product_id)
		if err != nil {
			fmt.Println("run query error ", err)
		}

	} else {
		_, err = helpers.QueryRun("INSERT INTO cart_item VALUES($1, $2, $3);", reference, product_id, quantity)
		if err != nil {
			fmt.Println("run query error", err)
		}
	}

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error generated": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"SUCCESS": "Respective products have been successfully added to the cart"})
}
