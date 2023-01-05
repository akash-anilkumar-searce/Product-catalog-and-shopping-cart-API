package HandlerInventory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/akash-searce/product-catalog/DbConnect"
	"github.com/akash-searce/product-catalog/Helpers"
	queries "github.com/akash-searce/product-catalog/Queries"
	response "github.com/akash-searce/product-catalog/Response"
	"github.com/akash-searce/product-catalog/typedefs"
)

func AddIntoInventory(w http.ResponseWriter, r *http.Request) {
	var inventory typedefs.Inventory
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
	json.Unmarshal(reqBody, &inventory)
	db := DbConnect.ConnectToDB()
	stmt, err := db.Prepare(queries.AddInventory)
	_, err = stmt.Exec(inventory.Product_Id, inventory.Quantity)
	fmt.Println(inventory)

	if err != nil {
		fmt.Println(err) //check here
	} else {
		Helpers.SendJResponse(response.InventoryDetailAdded, w)
	}

}
