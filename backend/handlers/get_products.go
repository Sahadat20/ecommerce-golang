package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"log"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	log.Println("from get produts")
	util.SendData(w, database.ProductList, 200)

}
