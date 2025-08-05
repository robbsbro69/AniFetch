package anime

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	BaseURL    = "https://api.github.com/repos/cat-milk/Anime-Girls-Holding-Programming-Books/contents"
	RawBaseURL = "https://raw.githubusercontent.com/cat-milk/Anime-Girls-Holding-Programming-Books/master"
)

type GitHubContent struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}

type Fetcher struct {
	cacheDir string
}

func NewFetcher(cacheDir string) *Fetcher {
	os.MkdirAll(cacheDir, 0755)
	return &Fetcher{cacheDir: cacheDir}
}

func (f *Fetcher) GetRandomAnimeGirl() (string, error) {
	// Create HTTP client with GitHub token if available
	client := &http.Client{}
	req, err := http.NewRequest("GET", BaseURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	
	// Add GitHub token if available
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "token "+token)
	}
	
	// Get list of directories (programming languages) from GitHub API
	resp, err := client.Do(req)
	if err != nil {
		// Fallback to cached images if API is unreachable
		return f.getRandomCachedImage()
	}
	defer resp.Body.Close()

	var contents []GitHubContent
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		// Fallback to cached images if API response is invalid
		return f.getRandomCachedImage()
	}

	// Filter for directories
	var directories []GitHubContent
	for _, content := range contents {
		if content.Type == "dir" {
			directories = append(directories, content)
		}
	}

	if len(directories) == 0 {
		return "", fmt.Errorf("no directories found")
	}

	// Select random directory
	randomDirIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(directories))))
	if err != nil {
		return "", fmt.Errorf("error generating random number: %v", err)
	}
	selectedDir := directories[randomDirIndex.Int64()]

	// Get images from the selected directory
	dirURL := fmt.Sprintf("%s/%s", BaseURL, selectedDir.Name)
	req, err = http.NewRequest("GET", dirURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating directory request: %v", err)
	}
	
	// Add GitHub token if available
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		req.Header.Set("Authorization", "token "+token)
	}
	
	resp, err = client.Do(req)
	if err != nil {
		// Fallback to cached images if directory API is unreachable
		return f.getRandomCachedImage()
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Check if response is an error
	var errorResponse struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
	
	// Try to decode as error first
	if err := json.Unmarshal(body, &errorResponse); err == nil && errorResponse.Message != "" {
		// Fallback to cached images if API returns an error (like rate limit)
		return f.getRandomCachedImage()
	}

	var dirContents []GitHubContent
	if err := json.Unmarshal(body, &dirContents); err != nil {
		// Fallback to cached images if directory response is invalid
		return f.getRandomCachedImage()
	}

	// Filter for image files
	var images []GitHubContent
	for _, content := range dirContents {
		if strings.HasSuffix(content.Name, ".png") || strings.HasSuffix(content.Name, ".jpg") || strings.HasSuffix(content.Name, ".jpeg") {
			images = append(images, content)
		}
	}

	if len(images) == 0 {
		// Try to use a cached image as fallback
		cachedImages, err := f.GetCachedImages()
		if err == nil && len(cachedImages) > 0 {
			// Return a random cached image
			randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(cachedImages))))
			if err == nil {
				return cachedImages[randomIndex.Int64()], nil
			}
		}
		return "", fmt.Errorf("no images found in directory %s", selectedDir.Name)
	}

	// Select random image
	randomImgIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(images))))
	if err != nil {
		return "", fmt.Errorf("error generating random number: %v", err)
	}
	selected := images[randomImgIndex.Int64()]
	
	// Download image to cache
	cachePath := filepath.Join(f.cacheDir, selected.Name)
	
	// Check if already cached - but let's get a random one each time for variety
	// if _, err := os.Stat(cachePath); err == nil {
	// 	return cachePath, nil
	// }

	// Download the image
	resp, err = http.Get(selected.DownloadURL)
	if err != nil {
		return "", fmt.Errorf("error downloading image: %v", err)
	}
	defer resp.Body.Close()

	file, err := os.Create(cachePath)
	if err != nil {
		return "", fmt.Errorf("error creating cache file: %v", err)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
	return cachePath, nil
}

func (f *Fetcher) GetCachedImages() ([]string, error) {
	files, err := os.ReadDir(f.cacheDir)
	if err != nil {
		return nil, err
	}

	var images []string
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			if strings.HasSuffix(name, ".png") || strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") {
				images = append(images, filepath.Join(f.cacheDir, name))
			}
		}
	}

	return images, nil
}

func (f *Fetcher) ClearCache() error {
	return os.RemoveAll(f.cacheDir)
}

func (f *Fetcher) getRandomCachedImage() (string, error) {
	cachedImages, err := f.GetCachedImages()
	if err != nil {
		return "", fmt.Errorf("error getting cached images: %v", err)
	}
	
	if len(cachedImages) == 0 {
		return "", fmt.Errorf("no cached images available")
	}
	
	// Return a random cached image
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(cachedImages))))
	if err != nil {
		return "", fmt.Errorf("error generating random number: %v", err)
	}
	
	return cachedImages[randomIndex.Int64()], nil
} 