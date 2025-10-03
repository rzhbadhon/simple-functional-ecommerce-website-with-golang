package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	productID := r.PathValue("id")
	pID, err := strconv.Atoi(productID)
	if err != nil{
		fmt.Println("Couldn't delete the product")
		http.Error(w, "Cant delete the product", 400)
		return
	}

	database.Delete(pID)

	util.SendData(w, "Successfully deleted the product", 201)

}