package main

import (
	"encoding/json"
	"fetch-assignment/internal/controllers"
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	cors "github.com/rs/cors"
)

var ()

func main() {
	port := 8080
	log.Printf("Serving to http://localhost:%d/", port)

	r := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		Debug:            true,
	}).Handler)

	// r.Use(loggingMiddleware)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		finalResponse := map[string]any{}
		finalResponseJson, _ := json.Marshal(finalResponse)
		w.Write(finalResponseJson)
	})
	r.Route("/receipts", func(r chi.Router) {
		r.Route("/process", func(r chi.Router) {
			// r.Use(ReceiptValidator)
			r.Post("/", controllers.ProcessReceipts)
		})
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/points", controllers.GetPoints)
		})

	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	_ = srv.ListenAndServe()
}
