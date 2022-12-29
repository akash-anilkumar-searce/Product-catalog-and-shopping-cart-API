package get_test

/*
import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/subramanyam-searce/product-catalog-go/helpers"
)

func GetCategoriesViaAPI(id int,t *testing.T) string {

	res, err := http.Get("http://localhost:8089" + "/getcategory"+ fmt.Sprint(id))
	helpers.HandleTestError("httpGetError", err, t)

	var v any

	err = json.NewDecoder(response.Body).Decode(&v)
	helpers.HandleTestError("jsonDecodingError", err, t)

	return v
}

func TestGetCategories(t *testing.T) {
	category := GetProduct_testapi(1, t)
	res := GetProduct_testapi(1, t)
	_, ok := category.([]any)

	if !ok {
		t.Errorf("Expected a slice of categories but got: " + fmt.Sprint(categories))
	}

	categories = GetProductsViaAPI(2, t)

	_, ok = categories.(map[string]any)["message"]

	if !ok {
		t.Errorf("Expected an error of categories but got: " + fmt.Sprint(categories))
	}
}
*/
