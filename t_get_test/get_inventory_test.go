package get_test

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func GetInventoryViaAPI(id int, t *testing.T) []map[string]int {
	res, err := http.Get("http://localhost:8089" + "/inventorydetail/" + fmt.Sprint(id))
	if err != nil {
		fmt.Println("httpGetError", err)
	}

	res_json := []map[string]int{}
	err = json.NewDecoder(res.Body).Decode(&res_json)
	if err != nil {
		fmt.Println("json decode error", err)
	}
	return res_json
}

func TestGetInventory(t *testing.T) {
	Product_id := 2
	expected_response := `[{"product_id":2,"quantity":4}]`
	res := GetInventoryViaAPI(Product_id, t)
	res_json, err := json.Marshal(res)
	if err != nil {
		fmt.Println("jsonmarshal error", err)
	}

	if expected_response != string(res_json) {
		t.Errorf("Expected: %v, Got: %v", expected_response, res_json)
	}
}
*/
