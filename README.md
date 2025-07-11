# pdfCraft - PDF Toolbox API

🚀 A powerful REST API for PDF processing operations built with Go and pdfcpu. 

## Features

- **PDF Merge**: Combine multiple PDF files into one
- **PDF Split**: Extract specific pages from a PDF
- **PDF Watermark**: Add text watermarks to PDFs
- **PDF Compression**: Optimize and compress PDF files
- **File Caching**: Intelligent caching system for processed files
- **Containerized**: Ready for Docker deployment
- **CI/CD**: GitHub Actions pipeline included

## Quick Start

### Running Locally

```bash
# Clone the repository
git clone https://github.com/Suryanshu185/pdfCraft.git
cd pdfCraft

# Install dependencies
go mod download

# Run the server
go run cmd/api/main.go
```

The API server will start on port 8080. Visit `http://localhost:8080/health` to check if it's running.

### Using Docker

```bash
# Build and run with Docker Compose
docker-compose up --build

# Or build and run manually
docker build -t pdfcraft-api .
docker run -p 8080:8080 pdfcraft-api
```

## API Documentation

### Health Check

**GET** `/health`

Returns the health status of the API.

```json
{
  "status": "healthy",
  "message": "PDF Toolbox API is running"
}
```

### PDF Operations

All PDF operations require file uploads via multipart form data.

#### 1. Merge PDFs

**POST** `/api/v1/pdf/merge`

Merges multiple PDF files into a single PDF.

**Parameters:**
- `files` (required): Multiple PDF files to merge

**Example:**
```bash
curl -X POST \
  -F "files=@document1.pdf" \
  -F "files=@document2.pdf" \
  -F "files=@document3.pdf" \
  http://localhost:8080/api/v1/pdf/merge \
  --output merged.pdf
```

#### 2. Split PDF

**POST** `/api/v1/pdf/split`

Extracts specific pages from a PDF file.

**Parameters:**
- `file` (required): PDF file to split
- `range` (required): Page range to extract
  - Single page: `"1"`
  - Range: `"1-3"`
  - Individual pages: `"1,3,5"`

**Example:**
```bash
curl -X POST \
  -F "file=@document.pdf" \
  -F "range=1-3" \
  http://localhost:8080/api/v1/pdf/split \
  --output split.pdf
```

#### 3. Add Watermark

**POST** `/api/v1/pdf/watermark`

Adds a text watermark to a PDF file.

**Parameters:**
- `file` (required): PDF file to watermark
- `text` (required): Watermark text

**Example:**
```bash
curl -X POST \
  -F "file=@document.pdf" \
  -F "text=CONFIDENTIAL" \
  http://localhost:8080/api/v1/pdf/watermark \
  --output watermarked.pdf
```

#### 4. Compress PDF

**POST** `/api/v1/pdf/compress`

Compresses a PDF file to reduce its size.

**Parameters:**
- `file` (required): PDF file to compress
- `level` (optional): Compression level (`low`, `medium`, `high`). Default: `medium`

**Example:**
```bash
curl -X POST \
  -F "file=@document.pdf" \
  -F "level=high" \
  http://localhost:8080/api/v1/pdf/compress \
  --output compressed.pdf
```

## Project Structure

```
pdfCraft/
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── pkg/
│   ├── handlers/            # HTTP handlers
│   │   ├── handlers.go
│   │   └── handlers_test.go
│   ├── services/            # Business logic
│   │   └── pdf.go
│   └── cache/               # Caching functionality
│       └── cache.go
├── .github/
│   └── workflows/
│       └── ci.yml           # CI/CD pipeline
├── uploads/                 # Temporary upload directory
├── temp/                    # Temporary processing directory
├── cache/                   # File cache directory
├── Dockerfile              # Container configuration
├── docker-compose.yml      # Docker Compose configuration
└── README.md               # This file
```

## Development

### Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerization)

### Building

```bash
# Build the application
go build -o pdfcraft-api cmd/api/main.go

# Run tests
go test -v ./...

# Build Docker image
docker build -t pdfcraft-api .
```

### Configuration

The application uses the following environment variables:

- `PORT`: Server port (default: 8080)

### Caching

The API includes built-in file caching to improve performance:

- Cache TTL: 1 hour
- Cache directory: `cache/`
- Automatic cleanup of expired files

## Production Deployment

### Docker Deployment

1. Build the Docker image:
```bash
docker build -t pdfcraft-api .
```

2. Run with Docker Compose:
```bash
docker-compose up -d
```

### Kubernetes Deployment

Example Kubernetes configuration:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: pdfcraft-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: pdfcraft-api
  template:
    metadata:
      labels:
        app: pdfcraft-api
    spec:
      containers:
      - name: pdfcraft-api
        image: pdfcraft-api:latest
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
---
apiVersion: v1
kind: Service
metadata:
  name: pdfcraft-api-service
spec:
  selector:
    app: pdfcraft-api
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License.

## Technology Stack

- **Language**: Go 1.21
- **PDF Processing**: pdfcpu
- **Web Framework**: Gorilla Mux
- **Containerization**: Docker
- **CI/CD**: GitHub Actions
- **Testing**: Go built-in testing