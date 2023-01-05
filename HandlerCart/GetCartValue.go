package HandlerCart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func GetCart1(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()
	reference := urlQuery.Get("ref")
	db := DbConnect.ConnectToDB()

	rows, err := db.Query("SELECT product_master.price,product_master.name,cart_item.quantity FROM (cart_item JOIN product_master ON cart_item.product_id = product_master.product_id) WHERE cart_item.ref=$1", reference)
	if err != nil {
		fmt.Println("error in query", err)
	}
	if rows.Next() == false {
		fmt.Println("Reference_id not found")
		w.Header().Add("Content-Type", "application/json")
		Helpers.SendJResponse(response.ReferenceIdNotExist, w)
	} else {
		list_of_cart := []typedefs.Newcart{}

		if err != nil {
			fmt.Println("error in getting category", err)
		}
		var total float32

		for rows.Next() {
			new_cart := typedefs.Newcart{}
			err := rows.Scan(&new_cart.Price, &new_cart.Product_name, &new_cart.Quantity)
			if err != nil {
				fmt.Println("error in rows next", err)
			}
			total += (float32(new_cart.Price) * float32(new_cart.Quantity))
			list_of_cart = append(list_of_cart, new_cart)
		}

		if len(list_of_cart) == 0 {
			Helpers.SendJResponse(response.DataNotFound, w)
			return
		}

		err = json.NewEncoder(w).Encode(list_of_cart)
		res := fmt.Sprintln("The total price of this cart is ", total)
		err = json.NewEncoder(w).Encode(res)

		fmt.Println(list_of_cart)
		fmt.Println(res)

	}
}
