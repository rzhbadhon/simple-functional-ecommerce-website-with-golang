package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)


//    01:31:00


type ReqCreatProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
} 

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct repo.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Error of invalid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(newProduct) //database.Store(newProduct)
	if err != nil{
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	util.SendData(w, createdProduct, http.StatusOK)

}
