package display

import (
	"fmt"
	"os"
	"os/exec"
	"golang.org/x/term"
)

// ImageDisplay handles different methods of displaying images in terminal
type ImageDisplay struct{
	size string
}

func NewImageDisplay() *ImageDisplay {
	return &ImageDisplay{size: "15x8"}
}

func NewImageDisplayWithSize(size string) *ImageDisplay {
	return &ImageDisplay{size: size}
}

func (id *ImageDisplay) DisplayImage(imagePath string) bool {
	// Try different image display methods in order of preference
	
	// 1. Try chafa (modern terminal image viewer)
	if id.tryChafa(imagePath) {
		return true
	}
	
	// 2. Try imgcat (iTerm2 image protocol)
	if id.tryImgcat(imagePath) {
		return true
	}
	
	// 3. Try kitty icat
	if id.tryKittyIcat(imagePath) {
		return true
	}
	
	// 4. Try terminal image protocols
	if id.tryTerminalProtocols(imagePath) {
		return true
	}
	
	return false
}

func (id *ImageDisplay) tryChafa(imagePath string) bool {
	// Get terminal size for optimal display
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		width, height, err = term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			// Fallback to default size
			width, height = 80, 40
		}
	}

	// Calculate display size with padding (leave space for system info)
	// Use more conservative sizing for better fit
	displayWidth := (width - 4) / 3  // Even smaller - 1/3 of terminal width
	displayHeight := (height - 6) / 3 // Even smaller - 1/3 of terminal height

	// Ensure reasonable minimum and maximum sizes
	if displayWidth < 12 {
		displayWidth = 12
	} else if displayWidth > 40 {
		displayWidth = 40
	}

	if displayHeight < 6 {
		displayHeight = 6
	} else if displayHeight > 20 {
		displayHeight = 20
	}

	sizeStr := fmt.Sprintf("%dx%d", displayWidth, displayHeight)

	// Try with dynamic terminal size for optimal quality
	cmd := exec.Command("chafa", 
		"--size", sizeStr,
		"--symbols", "block",
		"--colors", "256",
		"--dither", "none",
		"--color-space", "rgb",
		imagePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err == nil {
		return true
	}
	
	// Fallback to configured size
	cmd = exec.Command("chafa", 
		"--size", id.size,
		"--symbols", "block",
		"--colors", "256",
		"--dither", "ordered",
		imagePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err == nil {
		return true
	}
	
	// Fallback to character-based size for terminals that don't support pixel size
	cmd = exec.Command("chafa", 
		"--size", "30",
		"--symbols", "block",
		"--colors", "256",
		"--dither", "none",
		"--color-space", "rgb",
		imagePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err == nil {
		return true
	}
	
	// Final fallback with smaller size
	cmd = exec.Command("chafa", 
		"--size", "20",
		"--symbols", "block",
		"--colors", "256",
		"--dither", "none",
		"--color-space", "rgb",
		imagePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err == nil {
		return true
	}
	
	return false
}

func (id *ImageDisplay) tryImgcat(imagePath string) bool {
	cmd := exec.Command("imgcat", imagePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err == nil {
		return true
	}
	return false
}

func (id *ImageDisplay) tryKittyIcat(imagePath string) bool {
	cmd := exec.Command("kitty", "+kitten", "icat", imagePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err == nil {
		return true
	}
	return false
}

func (id *ImageDisplay) tryTerminalProtocols(imagePath string) bool {
	// Instead of trying terminal protocols that often don't work well,
	// show a nice ASCII art fallback
	fmt.Println("╭─────────────────────────╮")
	fmt.Println("│      (◕‿◕)            │")
	fmt.Println("│       /|\\             │")
	fmt.Println("│      / \\              │")
	fmt.Println("│                        │")
	fmt.Println("│  Holding a Programming │")
	fmt.Println("│      Book! 📚          │")
	fmt.Println("│                        │")
	fmt.Println("│  🎀 Anime Girl 🎀      │")
	fmt.Println("╰─────────────────────────╯")
	return true
}

func (id *ImageDisplay) GetSupportedTools() []string {
	tools := []string{}
	
	// Check for chafa
	if _, err := exec.LookPath("chafa"); err == nil {
		tools = append(tools, "chafa")
	}
	
	// Check for imgcat
	if _, err := exec.LookPath("imgcat"); err == nil {
		tools = append(tools, "imgcat")
	}
	
	// Check for kitty
	if _, err := exec.LookPath("kitty"); err == nil {
		tools = append(tools, "kitty icat")
	}
	
	return tools
} 