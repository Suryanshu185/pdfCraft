package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Suryanshu185/pdfCraft/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	// PDF processing endpoints
	r.HandleFunc("/api/v1/pdf/merge", handlers.MergePDFHandler).Methods("POST")
	r.HandleFunc("/api/v1/pdf/split", handlers.SplitPDFHandler).Methods("POST")
	r.HandleFunc("/api/v1/pdf/watermark", handlers.WatermarkPDFHandler).Methods("POST")
	r.HandleFunc("/api/v1/pdf/compress", handlers.CompressPDFHandler).Methods("POST")
	
	// Health check endpoint
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	
	fmt.Println("PDF Toolbox API server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}