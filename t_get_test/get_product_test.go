package get_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func GetProduct_testapi(id int, t *testing.T) map[string]string {
	res, err := http.Get("http://localhost:8089" + "/product/" + fmt.Sprint(id))
	if err != nil {
		fmt.Println("ALERT:HTTP ERROR FOUND", err)
	}

	res_json := map[string]string{}
	json.NewDecoder(res.Body).Decode(&res_json)

	return res_json
}

func TestGetProduct(t *testing.T) {

	//CHECK IF THE PRODUCT ID IS PRESENT OR NOT
	product_id := 1
	res := GetProduct_testapi(product_id, t)
	_, ok := res["product_id"]
	if !ok {
		t.Errorf("Expected Response: %v, Got Response: %v", "A Valid Product Map", res)
	}

	product_id = 500
	res = GetProduct_testapi(product_id, t)
	message, ok := res["message"]
	if !ok || message != "Product does not exist" {
		t.Errorf("Expected Response: %v, Got Response: %v", "Product does not exist", res)
	}

}
