package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Suryanshu185/pdfCraft/pkg/cache"
	"github.com/Suryanshu185/pdfCraft/pkg/services"
)

// Global cache instance
var fileCache *cache.FileCache

func init() {
	// Initialize cache with 1 hour TTL
	fileCache = cache.NewFileCache("cache", 1*time.Hour)
}

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{
		Status:  "healthy",
		Message: "PDF Toolbox API is running",
	}
	json.NewEncoder(w).Encode(response)
}

// MergePDFHandler handles PDF merge requests
func MergePDFHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form with 32MB max memory
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) < 2 {
		http.Error(w, "At least 2 PDF files are required for merging", http.StatusBadRequest)
		return
	}

	// Save uploaded files temporarily
	var filePaths []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error reading uploaded file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create temporary file
		tempFile, err := os.CreateTemp("temp", "upload_*.pdf")
		if err != nil {
			http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
			return
		}
		defer tempFile.Close()

		// Copy uploaded file to temporary file
		if _, err := io.Copy(tempFile, file); err != nil {
			http.Error(w, "Error saving uploaded file", http.StatusInternalServerError)
			return
		}

		filePaths = append(filePaths, tempFile.Name())
	}

	// Clean up temporary files
	defer func() {
		for _, path := range filePaths {
			os.Remove(path)
		}
	}()

	// Merge PDFs
	outputPath := fmt.Sprintf("temp/merged_%d.pdf", time.Now().Unix())
	if err := services.MergePDFs(filePaths, outputPath); err != nil {
		http.Error(w, fmt.Sprintf("Error merging PDFs: %v", err), http.StatusInternalServerError)
		return
	}

	// Serve the merged PDF
	serveFile(w, outputPath, "merged.pdf")
}

// SplitPDFHandler handles PDF split requests
func SplitPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get split range from form
	rangeStr := r.FormValue("range")
	if rangeStr == "" {
		http.Error(w, "Range parameter is required (e.g., '1-3' or '1,3,5')", http.StatusBadRequest)
		return
	}

	// Save uploaded file temporarily
	tempFile, err := os.CreateTemp("temp", "upload_*.pdf")
	if err != nil {
		http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	if _, err := io.Copy(tempFile, file); err != nil {
		http.Error(w, "Error saving uploaded file", http.StatusInternalServerError)
		return
	}

	// Split PDF
	outputPath := fmt.Sprintf("temp/split_%d.pdf", time.Now().Unix())
	if err := services.SplitPDF(tempFile.Name(), outputPath, rangeStr); err != nil {
		http.Error(w, fmt.Sprintf("Error splitting PDF: %v", err), http.StatusInternalServerError)
		return
	}

	// Serve the split PDF
	serveFile(w, outputPath, fmt.Sprintf("split_%s", fileHeader.Filename))
}

// WatermarkPDFHandler handles PDF watermark requests
func WatermarkPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get watermark text from form
	watermarkText := r.FormValue("text")
	if watermarkText == "" {
		http.Error(w, "Watermark text is required", http.StatusBadRequest)
		return
	}

	// Save uploaded file temporarily
	tempFile, err := os.CreateTemp("temp", "upload_*.pdf")
	if err != nil {
		http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	if _, err := io.Copy(tempFile, file); err != nil {
		http.Error(w, "Error saving uploaded file", http.StatusInternalServerError)
		return
	}

	// Add watermark
	outputPath := fmt.Sprintf("temp/watermarked_%d.pdf", time.Now().Unix())
	if err := services.AddWatermark(tempFile.Name(), outputPath, watermarkText); err != nil {
		http.Error(w, fmt.Sprintf("Error adding watermark: %v", err), http.StatusInternalServerError)
		return
	}

	// Serve the watermarked PDF
	serveFile(w, outputPath, fmt.Sprintf("watermarked_%s", fileHeader.Filename))
}

// CompressPDFHandler handles PDF compression requests
func CompressPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get compression level from form (optional)
	compressionLevel := r.FormValue("level")
	if compressionLevel == "" {
		compressionLevel = "medium"
	}

	// Save uploaded file temporarily
	tempFile, err := os.CreateTemp("temp", "upload_*.pdf")
	if err != nil {
		http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	if _, err := io.Copy(tempFile, file); err != nil {
		http.Error(w, "Error saving uploaded file", http.StatusInternalServerError)
		return
	}

	// Compress PDF
	outputPath := fmt.Sprintf("temp/compressed_%d.pdf", time.Now().Unix())
	if err := services.CompressPDF(tempFile.Name(), outputPath, compressionLevel); err != nil {
		http.Error(w, fmt.Sprintf("Error compressing PDF: %v", err), http.StatusInternalServerError)
		return
	}

	// Serve the compressed PDF
	serveFile(w, outputPath, fmt.Sprintf("compressed_%s", fileHeader.Filename))
}

// serveFile serves a file as a download
func serveFile(w http.ResponseWriter, filePath, fileName string) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Get file info
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Set headers
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	// Copy file to response
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Error serving file", http.StatusInternalServerError)
		return
	}

	// Clean up the temporary file
	go func() {
		time.Sleep(10 * time.Second) // Give time for download to complete
		os.Remove(filePath)
	}()
}