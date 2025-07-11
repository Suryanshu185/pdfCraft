package cache

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// FileCache handles file caching with TTL
type FileCache struct {
	cacheDir string
	ttl      time.Duration
}

// CacheEntry represents a cached file entry
type CacheEntry struct {
	Path      string
	CreatedAt time.Time
	TTL       time.Duration
}

// NewFileCache creates a new file cache
func NewFileCache(cacheDir string, ttl time.Duration) *FileCache {
	// Create cache directory if it doesn't exist
	os.MkdirAll(cacheDir, 0755)
	
	return &FileCache{
		cacheDir: cacheDir,
		ttl:      ttl,
	}
}

// generateKey creates a cache key from input parameters
func (fc *FileCache) generateKey(params ...string) string {
	h := md5.New()
	for _, param := range params {
		h.Write([]byte(param))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Get retrieves a file from cache if it exists and is not expired
func (fc *FileCache) Get(key string) (string, bool) {
	filePath := filepath.Join(fc.cacheDir, key)
	
	// Check if file exists
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return "", false
	}
	
	// Check if file is expired
	if time.Since(info.ModTime()) > fc.ttl {
		os.Remove(filePath)
		return "", false
	}
	
	return filePath, true
}

// Set stores a file in cache
func (fc *FileCache) Set(key string, srcPath string) error {
	dstPath := filepath.Join(fc.cacheDir, key)
	
	// Copy file to cache
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	
	_, err = io.Copy(dst, src)
	return err
}

// GenerateKey creates a cache key from operation and parameters
func (fc *FileCache) GenerateKey(operation string, params ...string) string {
	allParams := append([]string{operation}, params...)
	return fc.generateKey(allParams...)
}

// CleanExpired removes expired cache entries
func (fc *FileCache) CleanExpired() error {
	return filepath.Walk(fc.cacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() && time.Since(info.ModTime()) > fc.ttl {
			return os.Remove(path)
		}
		
		return nil
	})
}

// Clear removes all cache entries
func (fc *FileCache) Clear() error {
	return os.RemoveAll(fc.cacheDir)
}

// Size returns the total size of cached files
func (fc *FileCache) Size() (int64, error) {
	var size int64
	
	err := filepath.Walk(fc.cacheDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() {
			size += info.Size()
		}
		
		return nil
	})
	
	return size, err
}