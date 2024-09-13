package services

import (
	"http"

	"github.com/go-chi/chi/v5"
)

func ProcessReceipts(w http.ResponseWriter, req *http.Request) {
	var receipt Receipt
	receiptJson := req.Body
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	receiptId := chi.URLParam(req, "id")
}
