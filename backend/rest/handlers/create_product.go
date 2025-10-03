package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	header := r.Header.Get("Authorization")

	if header == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	headerArr := strings.Split(header, " ")
	fmt.Println("------", headerArr)

	if len(headerArr) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accessToken := headerArr[1]

	tokenParts := strings.Split(accessToken, ".")

	if len(tokenParts) != 3 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	jwtHeader := tokenParts[0]
	jwtPayload := tokenParts[1]
	signature := tokenParts[2]

	message := jwtHeader + "." + jwtPayload

	cnf := config.GetConfig()

	byteArraySecret := []byte(cnf.JwtSecretKey)
	byteArrayMessage := []byte(message)

	h := hmac.New(sha256.New, byteArraySecret)
	h.Write(byteArrayMessage)

	hash := h.Sum(nil)
	newSignature := base64UrlEncode(hash)

	if newSignature != signature{
		http.Error(w, "Unauthorized, tui hacker", http.StatusUnauthorized)
	}

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
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
