package HandlerCategory

import (
	"encoding/json"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/gorilla/mux"
)

func Deletecategory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	db := DbConnect.ConnectToDB()

	res, err := db.Exec(queries.DeleteCategory, x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				Helpers.SendJResponse(response.CategoryIdDoesNotExist, w)
				//result := fmt.Sprint("The value category_id does not exist,enter a Valid ID")
				//json.NewEncoder(w).Encode(result)
				return
			}
			json.NewEncoder(w).Encode(response.CategoryDetailDeleted)
		}

	}

	return

}
