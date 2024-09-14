package controllers

import (
	"encoding/json"
	"fetch-assignment/internal/models"
	"fetch-assignment/internal/services"
	"net/http"
	"strconv"

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
	services.ProcessReceipts(&receipt)
	w.Write([]byte("Receipt Processed"))

}

func GetPoints(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	receiptId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	points, err := services.GetPoints(receiptId)
	w.Write([]byte(strconv.Itoa(points)))
}
