package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pID, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	err = h.productRepo.Delete(pID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal sever error")
		return
	}

	util.SendData(w, http.StatusOK, "Successfully deleted the product")

}
