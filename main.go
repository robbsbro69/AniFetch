package main

import (
	"flag"
	"fmt"
	"os"

	"anifetch/pkg/anime"
	"anifetch/pkg/config"
	"anifetch/pkg/display"
	"anifetch/pkg/system"
)

func main() {
	// Parse command line flags
	var (
		noImage = flag.Bool("no-image", false, "Disable image display")
		clearCache = flag.Bool("clear-cache", false, "Clear cached images")
		showCache = flag.Bool("show-cache", false, "Show cached images")
		checkToken = flag.Bool("check-token", false, "Check GitHub token status")
		imageSize = flag.String("size", "40x20", "Image size (fallback if terminal size detection fails)")
	)
	flag.Parse()

	// Initialize configuration
	cfg := config.NewConfig()
	
	// Initialize display renderer with custom image size
	renderer := display.NewRenderer(!*noImage && cfg.ShowImage)
	renderer.SetImageSize(*imageSize)

	// Handle special commands
	if *clearCache {
		fetcher := anime.NewFetcher(cfg.GetCacheDir())
		if err := fetcher.ClearCache(); err != nil {
			renderer.DisplayError(fmt.Sprintf("Failed to clear cache: %v", err))
			os.Exit(1)
		}
		renderer.DisplaySuccess("Cache cleared successfully!")
		return
	}

	if *showCache {
		fetcher := anime.NewFetcher(cfg.GetCacheDir())
		images, err := fetcher.GetCachedImages()
		if err != nil {
			renderer.DisplayError(fmt.Sprintf("Failed to get cached images: %v", err))
			os.Exit(1)
		}
		
		if len(images) == 0 {
			renderer.DisplaySuccess("No cached images found.")
		} else {
			fmt.Printf("Cached images (%d):\n", len(images))
			for _, img := range images {
				fmt.Printf("  - %s\n", img)
			}
		}
		return
	}

	if *checkToken {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			fmt.Println("❌ No GitHub token found")
			fmt.Println("Set GITHUB_TOKEN environment variable for higher rate limits")
			fmt.Println("Rate limit: 60 requests/hour (unauthenticated)")
		} else {
			fmt.Println("✅ GitHub token found")
			fmt.Println("Rate limit: 5,000 requests/hour (authenticated)")
		}
		return
	}

	// Ensure cache directory exists
	if err := cfg.EnsureCacheDir(); err != nil {
		renderer.DisplayError(fmt.Sprintf("Failed to create cache directory: %v", err))
		os.Exit(1)
	}

	// Get system information
	sysInfo := system.GetSystemInfo()

	// Get anime girl image
	var animeGirlPath string
	if !*noImage {
		fetcher := anime.NewFetcher(cfg.GetCacheDir())
		path, err := fetcher.GetRandomAnimeGirl()
		if err != nil {
			renderer.DisplayError(fmt.Sprintf("Failed to get anime girl image: %v", err))
			// Continue without image
		} else {
			animeGirlPath = path
		}
	}

	// Display the information
	renderer.DisplayInfo(sysInfo, animeGirlPath)
} 