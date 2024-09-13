package controllers

import (
	"encoding/json"
	"fetch-assignment/internal/models"
	"fetch-assignment/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProcessReceipts(w http.ResponseWriter, req *http.Request) {
	var receipt models.Receipt
	receiptJson := req.Body
	err := json.NewDecoder(receiptJson).Decode(receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	services.ProcessReceipts(receipt)
	w.Write([]byte("Receipt Processed"))

}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	receiptId := chi.URLParam(req, "id")
}
