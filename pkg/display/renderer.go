package display

import (
	"fmt"
	"os"

	"anifetch/pkg/system"
)

type Renderer struct {
	showImage bool
	imageSize string
}

func NewRenderer(showImage bool) *Renderer {
	return &Renderer{showImage: showImage, imageSize: "40x20"}
}

func (r *Renderer) SetImageSize(size string) {
	r.imageSize = size
}

func (r *Renderer) DisplayInfo(info system.SystemInfo, animeGirlPath string) {
	// ANSI color codes
	const (
		reset   = "\033[0m"
		bold    = "\033[1m"
		red     = "\033[31m"
		green   = "\033[32m"
		yellow  = "\033[33m"
		blue    = "\033[34m"
		magenta = "\033[35m"
		cyan    = "\033[36m"
		white   = "\033[37m"
	)

	// Display anime girl using various terminal image protocols
	if r.showImage && animeGirlPath != "" {
		if !r.displayImage(animeGirlPath) {
			// Only show ASCII art if no image display methods work
			r.displayASCIIArt()
		}
	} else if r.showImage {
		// ASCII art fallback
		r.displayASCIIArt()
	}

	// Display system information
	fmt.Printf("%s%s%s@%s%s %s%s%s\n", bold, green, info.Hostname, reset, bold, blue, info.OS, reset)
	fmt.Printf("%s%s%s%s %s%s%s\n", bold, green, "────────", reset, bold, blue, "────")
	fmt.Printf("%sOS:%s %s%s%s\n", bold, reset, yellow, info.OS, reset)
	fmt.Printf("%sKernel:%s %s%s%s\n", bold, reset, yellow, info.Kernel, reset)
	fmt.Printf("%sUptime:%s %s%s%s\n", bold, reset, yellow, info.Uptime, reset)
	fmt.Printf("%sPackages:%s %s%s%s\n", bold, reset, yellow, info.Packages, reset)
	fmt.Printf("%sShell:%s %s%s%s\n", bold, reset, yellow, info.Shell, reset)
	fmt.Printf("%sCPU:%s %s%s%s\n", bold, reset, yellow, info.CPU, reset)
	fmt.Printf("%sMemory:%s %s%s%s\n", bold, reset, yellow, info.Memory, reset)
	fmt.Printf("%sDisk:%s %s%s%s\n", bold, reset, yellow, info.Disk, reset)
}

func (r *Renderer) displayImage(imagePath string) bool {
	// Try advanced image display methods with custom size
	imgDisplay := NewImageDisplayWithSize(r.imageSize)
	if imgDisplay.DisplayImage(imagePath) {
		return true
	}
	
	// Fallback to basic terminal protocols
	term := os.Getenv("TERM")
	
	switch term {
	case "xterm-kitty":
		// Kitty terminal image protocol
		fmt.Printf("\033]1337;File=inline=1;preserveAspectRatio=1:%s\007", imagePath)
		return true
	case "xterm-256color", "screen-256color":
		// Try iTerm2 image protocol
		fmt.Printf("\033]1337;File=inline=1;preserveAspectRatio=1:%s\007", imagePath)
		return true
	default:
		// Try generic image protocol
		fmt.Printf("\033]1337;File=inline=1;preserveAspectRatio=1:%s\007", imagePath)
		return true
	}
}

func (r *Renderer) displayASCIIArt() {
	// Cute ASCII art of an anime girl
	art := `
    ╭─────────────────────────╮
    │      (◕‿◕)             │
    │      /|\\               │
    │     / \\                │
    │                         │
    │   Holding a Programming │
    │   Book! 📚              │
    │                         │
    │   🎀  Anime Girl  🎀    │
    ╰─────────────────────────╯
`
	fmt.Print(art)
}

func (r *Renderer) DisplayError(message string) {
	const red = "\033[31m"
	const reset = "\033[0m"
	fmt.Fprintf(os.Stderr, "%sError: %s%s\n", red, message, reset)
}

func (r *Renderer) DisplaySuccess(message string) {
	const green = "\033[32m"
	const reset = "\033[0m"
	fmt.Printf("%s%s%s\n", green, message, reset)
} 