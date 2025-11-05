package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	pageAsStr := r.URL.Query().Get("page")
	limitAsStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageAsStr)
	limit, _ := strconv.Atoi(limitAsStr)

	if page == 0 || limit == 0{
		page = 1
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil{
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusOK, productList)

}
