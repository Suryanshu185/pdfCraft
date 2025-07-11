# pdfCraft API Deployment Guide

This guide provides step-by-step instructions for deploying pdfCraft API to various cloud platforms for public use.

## Prerequisites

- GitHub account with access to the pdfCraft repository
- Docker installed (for local testing)
- Domain name (optional, for custom domains)

## Quick Deployment (Recommended)

### 1. Heroku (Easiest)

1. **One-click deploy**: Click the "Deploy to Heroku" button in the README
2. **Manual deploy**:
   ```bash
   # Clone repository
   git clone https://github.com/Suryanshu185/pdfCraft.git
   cd pdfCraft
   
   # Install Heroku CLI
   curl https://cli-assets.heroku.com/install.sh | sh
   
   # Login to Heroku
   heroku login
   
   # Create app
   heroku create your-app-name
   
   # Set environment variables
   heroku config:set GO_ENV=production
   heroku config:set MAX_FILE_SIZE=50MB
   heroku config:set CACHE_TTL=3600
   
   # Deploy
   git push heroku main
   ```

3. **Custom domain** (optional):
   ```bash
   heroku domains:add api.yourdomain.com
   # Then add CNAME record in your DNS: api -> your-app-name.herokuapp.com
   ```

### 2. Railway (Fast & Modern)

1. **One-click deploy**: Click the "Deploy to Railway" button in the README
2. **Manual deploy**:
   ```bash
   # Install Railway CLI
   npm install -g @railway/cli
   
   # Login and deploy
   railway login
   railway link
   railway up
   ```

### 3. Render (Free Tier Available)

1. **One-click deploy**: Click the "Deploy to Render" button in the README
2. **Manual deploy**:
   - Connect GitHub repository to Render
   - Select "Web Service"
   - Choose "Docker" as build method
   - Deploy automatically

## Advanced Deployment

### Google Cloud Platform (Cloud Run)

1. **Setup**:
   ```bash
   # Install gcloud CLI
   curl https://sdk.cloud.google.com | bash
   
   # Login and set project
   gcloud auth login
   gcloud config set project YOUR_PROJECT_ID
   ```

2. **Deploy**:
   ```bash
   # Build and push to Container Registry
   gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/pdfcraft-api
   
   # Deploy to Cloud Run
   gcloud run deploy pdfcraft-api \
     --image gcr.io/YOUR_PROJECT_ID/pdfcraft-api \
     --platform managed \
     --region us-central1 \
     --allow-unauthenticated \
     --port 8080 \
     --memory 1Gi \
     --max-instances 100
   ```

### AWS (Elastic Container Service)

1. **Setup**:
   ```bash
   # Install AWS CLI
   curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
   unzip awscliv2.zip
   sudo ./aws/install
   
   # Configure credentials
   aws configure
   ```

2. **Deploy**:
   ```bash
   # Create ECR repository
   aws ecr create-repository --repository-name pdfcraft-api
   
   # Build and push
   docker build -t pdfcraft-api .
   docker tag pdfcraft-api:latest YOUR_AWS_ACCOUNT_ID.dkr.ecr.YOUR_REGION.amazonaws.com/pdfcraft-api:latest
   docker push YOUR_AWS_ACCOUNT_ID.dkr.ecr.YOUR_REGION.amazonaws.com/pdfcraft-api:latest
   
   # Deploy using ECS (use AWS Console or CLI)
   ```

## Configuration for Production

### Environment Variables

Copy `.env.example` to `.env` and configure:

```bash
# Required
PORT=8080
GO_ENV=production
MAX_FILE_SIZE=50MB

# Security
ENABLE_CORS=true
ALLOWED_ORIGINS=https://yourdomain.com
FORCE_HTTPS=true

# Performance
CACHE_TTL=3600
RATE_LIMIT_REQUESTS=100
```

### SSL/TLS Configuration

Most cloud platforms provide automatic SSL/TLS:

- **Heroku**: Automatic SSL for custom domains
- **Railway**: Automatic SSL for all deployments
- **Render**: Automatic SSL for all deployments
- **Google Cloud Run**: Automatic SSL for custom domains
- **AWS**: Use Application Load Balancer with SSL certificate

### Custom Domain Setup

1. **Add domain to platform** (platform-specific commands)
2. **Configure DNS**:
   ```
   Type: CNAME
   Name: api
   Value: your-app-url.platform.com
   ```
3. **Verify SSL certificate** (automatic on most platforms)

## Monitoring and Maintenance

### Health Monitoring

Use the provided health check script:

```bash
# Basic health check
./health-check.sh basic https://your-api-url.com

# Comprehensive check
./health-check.sh comprehensive https://your-api-url.com
```

### Logging

Configure logging based on your platform:

- **Heroku**: `heroku logs --tail`
- **Railway**: Built-in logging dashboard
- **Render**: Built-in logging dashboard
- **Google Cloud**: Cloud Logging
- **AWS**: CloudWatch

### Performance Monitoring

Set up monitoring:

- **Uptime monitoring**: Use services like UptimeRobot, Pingdom
- **Performance monitoring**: Use APM tools like New Relic, Datadog
- **Error tracking**: Use Sentry, Rollbar

## Security Checklist

- [ ] HTTPS enabled
- [ ] Rate limiting configured
- [ ] File size limits set
- [ ] CORS configured properly
- [ ] Security headers enabled
- [ ] API key protection (if needed)
- [ ] Regular security updates

## Cost Optimization

### Free Tiers

- **Heroku**: 550-1000 dyno hours/month
- **Railway**: $5/month credit
- **Render**: Free tier available
- **Google Cloud**: $300 credit for new accounts
- **AWS**: Free tier for 12 months

### Optimization Tips

1. **Use appropriate instance sizes**
2. **Enable auto-scaling**
3. **Implement caching**
4. **Optimize Docker image size**
5. **Use CDN for static assets**

## Troubleshooting

### Common Issues

1. **Memory errors**: Increase memory allocation
2. **File upload failures**: Check file size limits
3. **SSL errors**: Verify domain configuration
4. **Rate limiting**: Adjust rate limit settings

### Debug Mode

Enable debug mode for troubleshooting:

```bash
# Set environment variable
DEBUG=true
LOG_LEVEL=debug
```

### Getting Help

- **GitHub Issues**: Report bugs and request features
- **Documentation**: Check API documentation
- **Community**: Join discussions in repository

## Scaling

### Horizontal Scaling

- **Heroku**: `heroku ps:scale web=3`
- **Railway**: Auto-scaling available
- **Google Cloud Run**: Auto-scaling built-in
- **AWS ECS**: Configure auto-scaling group

### Database Scaling

If using database:

- **Connection pooling**
- **Read replicas**
- **Database sharding**

### Caching Strategy

- **Redis for distributed caching**
- **CDN for static assets**
- **Application-level caching**

## Maintenance

### Regular Tasks

1. **Monitor logs for errors**
2. **Update dependencies**
3. **Check SSL certificate expiration**
4. **Monitor resource usage**
5. **Backup configurations**

### Updates

```bash
# Update dependencies
go mod tidy
go mod download

# Rebuild and redeploy
docker build -t pdfcraft-api .
# Deploy using your platform's method
```

## Support

- **Documentation**: [README.md](README.md)
- **Issues**: [GitHub Issues](https://github.com/Suryanshu185/pdfCraft/issues)
- **API Reference**: Available at `/health` endpoint