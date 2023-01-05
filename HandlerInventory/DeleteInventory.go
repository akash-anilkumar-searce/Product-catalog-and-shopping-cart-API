package HandlerInventory

import (
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/gorilla/mux"
)

func Deleteinventory(w http.ResponseWriter, r *http.Request) {
	x := mux.Vars(r)["id"]
	db := DbConnect.ConnectToDB()

	res, err := db.Exec(queries.DeleteInventory, x)

	if err == nil {

		count, err := res.RowsAffected()
		if err == nil {
			if count == 0 {
				Helpers.SendJResponse(response.ProductidDoesNotExist, w)
				return
			}
			Helpers.SendJResponse(response.InventoryDetailDeleted, w)
		}

	}

	return

}
