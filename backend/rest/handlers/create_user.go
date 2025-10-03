package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return 
	}

	createdUser := newUser.Store()

	util.SendData(w, createdUser, http.StatusCreated)


}