package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Error of invalid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, 201)

}
