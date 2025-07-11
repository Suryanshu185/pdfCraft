package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

// MergePDFs merges multiple PDF files into a single PDF
func MergePDFs(inputPaths []string, outputPath string) error {
	if len(inputPaths) < 2 {
		return fmt.Errorf("at least 2 PDF files are required for merging")
	}

	// Use pdfcpu to merge PDFs
	return api.MergeCreateFile(inputPaths, outputPath, false, nil)
}

// SplitPDF splits a PDF file based on the specified range
func SplitPDF(inputPath, outputPath, rangeStr string) error {
	// Parse range string (e.g., "1-3" or "1,3,5")
	pageSelection, err := parsePageRange(rangeStr)
	if err != nil {
		return fmt.Errorf("invalid page range: %v", err)
	}

	// Convert string pages to int pages
	var pageNumbers []int
	for _, pageStr := range pageSelection {
		pageNum, err := strconv.Atoi(pageStr)
		if err != nil {
			return fmt.Errorf("invalid page number: %s", pageStr)
		}
		pageNumbers = append(pageNumbers, pageNum)
	}

	// Extract pages using TrimFile (which extracts specified pages)
	return api.TrimFile(inputPath, outputPath, pageSelection, nil)
}

// AddWatermark adds a text watermark to a PDF
func AddWatermark(inputPath, outputPath, watermarkText string) error {
	// Create a simple text watermark
	wm, err := api.TextWatermark(watermarkText, "Helvetica 48", false, false, types.POINTS)
	if err != nil {
		return fmt.Errorf("error creating watermark: %v", err)
	}

	// Apply watermark
	return api.AddWatermarksFile(inputPath, outputPath, nil, wm, nil)
}

// CompressPDF compresses a PDF file
func CompressPDF(inputPath, outputPath, level string) error {
	// Create optimization configuration based on level
	conf := model.NewDefaultConfiguration()
	
	switch strings.ToLower(level) {
	case "low":
		conf.Cmd = model.OPTIMIZE
	case "medium":
		conf.Cmd = model.OPTIMIZE
		conf.StatsFileName = ""
	case "high":
		conf.Cmd = model.OPTIMIZE
		conf.StatsFileName = ""
	default:
		conf.Cmd = model.OPTIMIZE
	}

	// Optimize/compress the PDF
	return api.OptimizeFile(inputPath, outputPath, conf)
}

// parsePageRange parses a page range string into a pdfcpu IntSet
func parsePageRange(rangeStr string) ([]string, error) {
	rangeStr = strings.TrimSpace(rangeStr)
	
	// Handle different range formats
	if strings.Contains(rangeStr, "-") {
		// Range format: "1-3"
		parts := strings.Split(rangeStr, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format")
		}
		
		start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start page: %v", err)
		}
		
		end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end page: %v", err)
		}
		
		if start > end {
			return nil, fmt.Errorf("start page must be less than or equal to end page")
		}
		
		var pages []string
		for i := start; i <= end; i++ {
			pages = append(pages, strconv.Itoa(i))
		}
		return pages, nil
	} else if strings.Contains(rangeStr, ",") {
		// Individual pages: "1,3,5"
		parts := strings.Split(rangeStr, ",")
		var pages []string
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if _, err := strconv.Atoi(part); err != nil {
				return nil, fmt.Errorf("invalid page number: %s", part)
			}
			pages = append(pages, part)
		}
		return pages, nil
	} else {
		// Single page: "1"
		if _, err := strconv.Atoi(rangeStr); err != nil {
			return nil, fmt.Errorf("invalid page number: %s", rangeStr)
		}
		return []string{rangeStr}, nil
	}
}