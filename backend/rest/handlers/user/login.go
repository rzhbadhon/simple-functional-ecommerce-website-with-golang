package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		util.SendError(w, http.StatusBadRequest, "Invalid request data")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if usr == nil || err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, accessToken)

}
