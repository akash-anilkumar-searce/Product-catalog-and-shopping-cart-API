package handler_cart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/akash-searce/product-catalog/helpers"
	"github.com/google/uuid"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	//create a cart and use specific cart reference id to access the cart.
	ref := uuid.New()

	_, err := helpers.QueryRun("INSERT INTO cart_reference VALUES($1, $2);", ref, time.Now())
	if err != nil {
		fmt.Println("Query Run Error", err)
	}
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
	} else {
		json.NewEncoder(w).Encode((map[string]uuid.UUID{"ref": ref}))
	}
}
