package get_test

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)


func GetCart_testAPI(ref string, t *testing.T) any {
	response, err := http.Get("http://localhost:8089" + "/cart/get?ref=" + ref)
	if err != nil {
		fmt.Println("ALERT:httpget Error produced", err)
	}

	var response_json any
	err = json.NewDecoder(response.Body).Decode(&response_json)
	if err != nil {
		fmt.Println("ALERT:DECODING ERROR", err)
	}

	return response_json
}

func TestGetCart(t *testing.T) {
	test_cases := []map[string]any{
		//Valid Cart reference
		{"ref": " 9b5f2120-293f-45d7-8f03-b680a01c2f35", "expected_response": `{"created_at":"2022-12-26 13:46:13.759578",[{"price":2000,"quantity":2,"product_name":"Nokia"}]
		"The total price of this cart is  4000\n"`},

		//Invalid cart Reference
		{"ref": "2325f", "expected_response": `{"message":"Cart Reference is Invalid"}`},
	}

	for _, v := range test_cases {
		response := GetCart_testAPI(v["ref"].(string), t)

		json_response, err := json.Marshal(response)
		if err != nil {
			fmt.Println("marshal error produced", err)
		}

		if string(json_response) != v["expected_response"].(string) {
			t.Errorf("Expected Response: %v, Got Response: %v", v["expected_response"].(string), string(json_response))
		}
	}
}
*/
