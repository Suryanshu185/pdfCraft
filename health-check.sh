#!/bin/bash

# Health check script for pdfCraft API
# This script can be used by cloud platforms for health monitoring

# Configuration
API_URL="${API_URL:-http://localhost:8080}"
HEALTH_ENDPOINT="/health"
TIMEOUT=10

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to log messages
log() {
    echo -e "$(date '+%Y-%m-%d %H:%M:%S') - $1"
}

# Function to check health
check_health() {
    local url="$1"
    local response
    local http_code
    
    log "${YELLOW}Checking health at: $url${NC}"
    
    # Make the health check request
    response=$(curl -s -w "\n%{http_code}" --max-time $TIMEOUT "$url" 2>/dev/null)
    
    if [ $? -ne 0 ]; then
        log "${RED}❌ Health check failed: Unable to connect to $url${NC}"
        return 1
    fi
    
    # Extract HTTP status code
    http_code=$(echo "$response" | tail -n1)
    response_body=$(echo "$response" | head -n -1)
    
    # Check HTTP status code
    if [ "$http_code" != "200" ]; then
        log "${RED}❌ Health check failed: HTTP $http_code${NC}"
        return 1
    fi
    
    # Check response body
    if echo "$response_body" | grep -q '"status":"healthy"'; then
        log "${GREEN}✅ Health check passed: API is healthy${NC}"
        log "Response: $response_body"
        return 0
    else
        log "${RED}❌ Health check failed: Invalid response${NC}"
        log "Response: $response_body"
        return 1
    fi
}

# Function to check API endpoints
check_endpoints() {
    local base_url="$1"
    local endpoints=(
        "/health"
        "/api/v1/pdf/merge"
        "/api/v1/pdf/split"
        "/api/v1/pdf/watermark"
        "/api/v1/pdf/compress"
    )
    
    log "${YELLOW}Checking API endpoints...${NC}"
    
    for endpoint in "${endpoints[@]}"; do
        local url="$base_url$endpoint"
        local response
        local http_code
        
        # For non-health endpoints, we expect different responses
        if [ "$endpoint" = "/health" ]; then
            check_health "$url"
        else
            # For other endpoints, check if they exist (should return 405 Method Not Allowed for GET)
            response=$(curl -s -w "\n%{http_code}" --max-time $TIMEOUT "$url" 2>/dev/null)
            http_code=$(echo "$response" | tail -n1)
            
            if [ "$http_code" = "405" ] || [ "$http_code" = "400" ]; then
                log "${GREEN}✅ Endpoint $endpoint is accessible${NC}"
            else
                log "${YELLOW}⚠️  Endpoint $endpoint returned HTTP $http_code${NC}"
            fi
        fi
    done
}

# Function to run comprehensive health check
comprehensive_check() {
    local base_url="$1"
    
    log "${YELLOW}Starting comprehensive health check...${NC}"
    
    # Check basic health
    if ! check_health "$base_url/health"; then
        return 1
    fi
    
    # Check all endpoints
    check_endpoints "$base_url"
    
    log "${GREEN}✅ Comprehensive health check completed${NC}"
    return 0
}

# Main execution
main() {
    local mode="${1:-basic}"
    local url="${2:-$API_URL}"
    
    case "$mode" in
        "basic")
            check_health "$url/health"
            ;;
        "comprehensive")
            comprehensive_check "$url"
            ;;
        "endpoints")
            check_endpoints "$url"
            ;;
        *)
            echo "Usage: $0 [basic|comprehensive|endpoints] [url]"
            echo "  basic: Basic health check (default)"
            echo "  comprehensive: Full health and endpoint check"
            echo "  endpoints: Check all API endpoints"
            echo "  url: API base URL (default: $API_URL)"
            exit 1
            ;;
    esac
}

# Run the health check
main "$@"
exit $?