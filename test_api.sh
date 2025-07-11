#!/bin/bash

# Test script for pdfCraft API

API_URL="http://localhost:8080"
TEST_DIR="test_files"

echo "🧪 Testing pdfCraft API..."

# Create test directory
mkdir -p $TEST_DIR

# Test 1: Health check
echo "1. Testing health endpoint..."
curl -s "$API_URL/health" | jq .
echo ""

# Test 2: Create sample PDF for testing (using a simple text file converted to PDF)
echo "2. Creating test PDFs..."
echo "This is a test PDF content for document 1" > "$TEST_DIR/test1.txt"
echo "This is a test PDF content for document 2" > "$TEST_DIR/test2.txt"
echo "This is a test PDF content for document 3" > "$TEST_DIR/test3.txt"

# For this demo, we'll assume we have sample PDF files
# In a real scenario, you would have actual PDF files

echo "✅ Test setup complete!"
echo ""
echo "To test the API endpoints, you can use the following curl commands:"
echo ""
echo "# Test PDF merge:"
echo "curl -X POST -F 'files=@doc1.pdf' -F 'files=@doc2.pdf' $API_URL/api/v1/pdf/merge --output merged.pdf"
echo ""
echo "# Test PDF split:"
echo "curl -X POST -F 'file=@document.pdf' -F 'range=1-3' $API_URL/api/v1/pdf/split --output split.pdf"
echo ""
echo "# Test PDF watermark:"
echo "curl -X POST -F 'file=@document.pdf' -F 'text=CONFIDENTIAL' $API_URL/api/v1/pdf/watermark --output watermarked.pdf"
echo ""
echo "# Test PDF compression:"
echo "curl -X POST -F 'file=@document.pdf' -F 'level=high' $API_URL/api/v1/pdf/compress --output compressed.pdf"
echo ""
echo "📝 Note: You need actual PDF files to test the PDF processing endpoints."