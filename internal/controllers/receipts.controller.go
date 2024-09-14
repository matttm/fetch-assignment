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
	err := json.NewDecoder(receiptJson).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := services.ProcessReceipts(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := models.TransactionIdResponse{
		Id: id,
	}
	b, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(b)

}

func GetPoints(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	points, err := services.GetPoints(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := models.GetPointsResponse{
		Points: points,
	}
	b, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(b)
}
