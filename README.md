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

## Web Hosting for Public Use

### Quick Deploy Options

Deploy pdfCraft API to the web with these one-click solutions:

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/Suryanshu185/pdfCraft)

[![Deploy to Railway](https://railway.app/button.svg)](https://railway.app/template/pdfcraft-api)

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/Suryanshu185/pdfCraft)

### Cloud Platform Deployment

#### 1. **Heroku** (Recommended for beginners)

```bash
# Install Heroku CLI and login
heroku login

# Create a new Heroku app
heroku create your-pdfcraft-api

# Set environment variables
heroku config:set PORT=8080
heroku config:set GO_ENV=production

# Deploy
git push heroku main
```

**Custom Domain Setup:**
```bash
# Add custom domain
heroku domains:add api.yourdomain.com

# Enable SSL (automatic with custom domains)
heroku certs:auto:enable
```

#### 2. **Google Cloud Platform (Cloud Run)**

```bash
# Build and push to Google Container Registry
gcloud builds submit --tag gcr.io/[PROJECT-ID]/pdfcraft-api

# Deploy to Cloud Run
gcloud run deploy pdfcraft-api \
  --image gcr.io/[PROJECT-ID]/pdfcraft-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080 \
  --memory 1Gi \
  --cpu 1 \
  --max-instances 100
```

#### 3. **AWS (Elastic Container Service)**

```bash
# Create ECR repository
aws ecr create-repository --repository-name pdfcraft-api

# Build and push Docker image
docker build -t pdfcraft-api .
docker tag pdfcraft-api:latest [AWS_ACCOUNT_ID].dkr.ecr.[REGION].amazonaws.com/pdfcraft-api:latest
docker push [AWS_ACCOUNT_ID].dkr.ecr.[REGION].amazonaws.com/pdfcraft-api:latest

# Deploy using ECS (via AWS Console or CLI)
```

#### 4. **Railway** (Simple deployment)

```bash
# Install Railway CLI
npm install -g @railway/cli

# Login and deploy
railway login
railway link
railway up
```

#### 5. **Render** (Free tier available)

1. Connect your GitHub repository to Render
2. Select "Web Service"
3. Use Docker build
4. Set environment variables in Render dashboard

### Environment Variables for Production

Set these environment variables in your hosting platform:

```bash
# Server Configuration
PORT=8080
GO_ENV=production
HOST=0.0.0.0

# File Storage (for persistent storage)
UPLOAD_DIR=/app/uploads
CACHE_DIR=/app/cache
TEMP_DIR=/app/temp

# Cache Configuration
CACHE_TTL=3600  # 1 hour in seconds
MAX_CACHE_SIZE=1GB

# Security
MAX_FILE_SIZE=50MB
ALLOWED_ORIGINS=https://yourdomain.com,https://www.yourdomain.com
ENABLE_CORS=true

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_WINDOW=3600  # 1 hour in seconds

# Monitoring
LOG_LEVEL=info
ENABLE_METRICS=true
```

### Security Considerations

#### 1. **File Upload Security**
- Maximum file size limits (50MB recommended)
- File type validation (PDF only)
- Virus scanning for uploaded files
- Temporary file cleanup

#### 2. **API Security**
```bash
# Add these headers in production
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block
Content-Security-Policy: default-src 'self'
```

#### 3. **Rate Limiting**
Implement rate limiting to prevent abuse:
- 100 requests per hour per IP
- 10 MB total upload per hour per IP
- Concurrent request limiting

#### 4. **HTTPS/TLS**
- Always use HTTPS in production
- Redirect HTTP to HTTPS
- Use strong TLS configuration

### Domain Setup

#### 1. **Custom Domain Configuration**
```bash
# For Heroku
heroku domains:add api.yourdomain.com

# For Google Cloud Run
gcloud run domain-mappings create \
  --service pdfcraft-api \
  --domain api.yourdomain.com
```

#### 2. **DNS Configuration**
Add these DNS records:
```
Type: CNAME
Name: api
Value: your-app-name.herokuapp.com  # Or platform-specific URL
```

### Monitoring and Logging

#### 1. **Health Monitoring**
- Use the `/health` endpoint for uptime monitoring
- Set up alerts for API downtime
- Monitor response times and error rates

#### 2. **Logging Setup**
```bash
# Production logging configuration
LOG_LEVEL=info
LOG_FORMAT=json
LOG_OUTPUT=stdout

# Optional: Send logs to external services
LOGZIO_TOKEN=your_token  # For Logz.io
PAPERTRAIL_URL=your_url  # For Papertrail
```

#### 3. **Performance Monitoring**
- Monitor CPU and memory usage
- Track file processing times
- Set up alerts for high resource usage

### Scaling and Performance

#### 1. **Horizontal Scaling**
```bash
# Heroku
heroku ps:scale web=3

# Google Cloud Run (auto-scaling)
gcloud run services update pdfcraft-api \
  --min-instances=1 \
  --max-instances=100

# AWS ECS (auto-scaling group)
aws ecs update-service \
  --cluster pdfcraft-cluster \
  --service pdfcraft-api \
  --desired-count 3
```

#### 2. **Load Balancing**
- Use platform-native load balancers
- Implement health checks
- Configure SSL termination

#### 3. **Caching Strategy**
- Enable Redis for distributed caching
- Use CDN for static assets
- Implement response caching

### API Usage Examples

#### Production API Base URL
```bash
# Replace with your actual domain
API_BASE_URL="https://api.yourdomain.com"

# Health check
curl "$API_BASE_URL/health"

# Merge PDFs
curl -X POST \
  -F "files=@document1.pdf" \
  -F "files=@document2.pdf" \
  "$API_BASE_URL/api/v1/pdf/merge" \
  --output merged.pdf
```

### Cost Optimization

#### 1. **Resource Optimization**
- Use appropriate instance sizes
- Implement request batching
- Optimize Docker image size
- Use multi-stage builds

#### 2. **Storage Optimization**
- Implement file cleanup schedules
- Use object storage for large files
- Compress temporary files

### Backup and Disaster Recovery

#### 1. **Data Backup**
- Regular database backups (if using one)
- Configuration backup
- Code repository backup

#### 2. **Disaster Recovery**
- Multi-region deployment
- Automated failover
- Regular recovery testing

### Production Deployment

#### Docker Deployment

1. Build the Docker image:
```bash
docker build -t pdfcraft-api .
```

2. Run with Docker Compose:
```bash
docker-compose up -d
```

#### Kubernetes Deployment

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
        - name: GO_ENV
          value: "production"
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
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
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: pdfcraft-api-ingress
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  tls:
  - hosts:
    - api.yourdomain.com
    secretName: pdfcraft-tls
  rules:
  - host: api.yourdomain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: pdfcraft-api-service
            port:
              number: 80
```

### Troubleshooting

#### Common Issues

1. **File Upload Errors**
   - Check file size limits (50MB default)
   - Ensure PDF file format is valid
   - Verify disk space availability

2. **Memory Issues**
   - Increase container memory limits
   - Implement request queuing for large files
   - Monitor memory usage patterns

3. **Performance Issues**
   - Enable file caching
   - Implement request batching
   - Use CDN for static assets

4. **SSL/TLS Issues**
   - Verify domain DNS configuration
   - Check certificate validity
   - Ensure proper port binding

#### Debug Mode

Enable debug mode for troubleshooting:
```bash
# Set environment variable
DEBUG=true
LOG_LEVEL=debug

# Or run with debug flag
./pdfcraft-api --debug
```

### Support

- **Issues**: [GitHub Issues](https://github.com/Suryanshu185/pdfCraft/issues)
- **Documentation**: [API Documentation](https://api.yourdomain.com/docs)
- **Status Page**: [System Status](https://status.yourdomain.com)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## Quick Hosting Guide

Ready to host pdfCraft API online for public use? Here are the fastest ways to get started:

### 🚀 One-Click Deployment

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/Suryanshu185/pdfCraft)
[![Deploy to Railway](https://railway.app/button.svg)](https://railway.app/template/pdfcraft-api)
[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/Suryanshu185/pdfCraft)

### 📋 Quick Setup Steps

1. **Choose a platform**: Heroku (easiest), Railway (fast), or Render (free tier)
2. **Click deploy button** above or manually deploy using platform CLI
3. **Set environment variables**: `PORT=8080`, `GO_ENV=production`
4. **Configure custom domain** (optional) in platform settings
5. **Enable SSL/TLS** for secure API access

### 🔗 Your API will be available at:
- `https://your-app.herokuapp.com` (Heroku)
- `https://your-app.railway.app` (Railway)
- `https://your-app.onrender.com` (Render)

### 📖 Need detailed instructions?
See the complete [Web Hosting for Public Use](#web-hosting-for-public-use) section above for:
- Production configurations
- Security best practices
- Scaling and monitoring
- Custom domain setup
- SSL certificate installation

## License

This project is licensed under the MIT License.

## Technology Stack

- **Language**: Go 1.21
- **PDF Processing**: pdfcpu
- **Web Framework**: Gorilla Mux
- **Containerization**: Docker
- **CI/CD**: GitHub Actions
- **Testing**: Go built-in testing