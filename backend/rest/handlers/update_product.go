package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if(err != nil){
		http.Error(w, "Cant update the product", 400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil{
		fmt.Println(err)
		http.Error(w, "Cant update", 400)
	}

	newProduct.ID = pID

	database.Update(newProduct)

	util.SendData(w, "Successfully updated the product", 201)
	
}