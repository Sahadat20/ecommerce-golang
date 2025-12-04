package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// handlePreflightReq(w, r)

	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)
	util.SendData(w, newProduct, 201)

}
