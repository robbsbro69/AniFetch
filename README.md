# AniFetch

A neofetch-like system information tool that displays anime girls holding programming books instead of the traditional OS logo.

## Features

- Displays system information similar to neofetch/fastfetch
- Fetches random anime girls holding programming books from the [Anime-Girls-Holding-Programming-Books](https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books) repository
- **Random selection** from different programming language directories (JavaScript, Python, C++, etc.)
- Caches images locally for faster subsequent runs
- **Dynamic terminal size detection** for optimal image display
- **High-quality image rendering** with block symbols and 256 colors
- Supports multiple image display methods:
  - `chafa` (modern terminal image viewer) - **Recommended**
  - `imgcat` (iTerm2)
  - `kitty icat` (Kitty terminal)
  - Terminal image protocols
  - ASCII art fallback
- Cross-platform support (Linux, macOS, Windows)
- Command line options for customization
- **GitHub Personal Access Token support** for higher API rate limits

## System Information Displayed

- OS and Kernel version
- Hostname
- Uptime
- Package count (supports pacman, dpkg, rpm, brew, nix)
- Shell
- CPU information
- Memory usage
- Disk usage

## Installation

### Quick Start (Recommended)

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/anifetch.git
   cd anifetch
   ```

2. **Build the application:**
   ```bash
   go build -o anifetch main.go
   ```

3. **Install globally (optional):**
   ```bash
   sudo cp anifetch /usr/local/bin/
   ```

4. **Install chafa for best image quality:**
   ```bash
   # On Arch Linux:
   sudo pacman -S chafa
   
   # On Ubuntu/Debian:
   sudo apt install chafa
   
   # On macOS:
   brew install chafa
   ```

5. **Run AniFetch:**
   ```bash
   ./anifetch
   # or if installed globally:
   anifetch
   ```

### Prerequisites

- **Go 1.21 or later** - Required for building the application
- **`chafa`** - For optimal image display (recommended)
- **A modern terminal** - That supports image protocols for best results
- **Git** - For cloning the repository

### System Requirements

- **Linux**: Most distributions supported (Arch, Ubuntu, Debian, etc.)
- **macOS**: 10.14 or later
- **Windows**: Windows 10 or later (with WSL recommended)
- **Terminal**: kitty, iTerm2, or any terminal with image support

### GitHub API Token (Optional but Recommended)

For higher rate limits (5,000 requests/hour instead of 60), set up a GitHub Personal Access Token:

1. Go to [GitHub Settings → Developer settings → Personal access tokens](https://github.com/settings/tokens)
2. Click "Generate new token (classic)"
3. Give it a name like "anifetch-cli"
4. **No scopes needed** for public repositories
5. Copy the token and set it as an environment variable:

```bash
export GITHUB_TOKEN="your_token_here"
```

**Rate Limits:**
- **Unauthenticated:** 60 requests/hour
- **With Personal Access Token:** 5,000 requests/hour

### Alternative Installation Methods

#### Using the install script:
```bash
chmod +x install.sh
./install.sh
```

#### Manual build from source:
```bash
git clone <repository-url>
cd anifetch
go build -o anifetch main.go
```

#### Install globally:
```bash
sudo cp anifetch /usr/local/bin/
```

## Usage

### Basic Usage

Simply run:

```bash
./anifetch
```

or if installed globally:

```bash
anifetch
```

### First Run

On first run, AniFetch will:
1. Create a cache directory (`~/.anifetch/`)
2. Download anime girl images from GitHub
3. Display system information with a random anime girl

### Troubleshooting

**If images don't display:**
- Install `chafa`: `sudo pacman -S chafa` (Arch) or `sudo apt install chafa` (Ubuntu)
- Try a different terminal (kitty, iTerm2, etc.)
- Use `--no-image` to disable images temporarily

**If you get rate limit errors:**
- Set up a GitHub token (see GitHub API Token section above)
- The tool will use cached images as fallback

### Command Line Options

```bash
./anifetch --help          # Show help
./anifetch --no-image      # Disable image display
./anifetch --show-cache    # Show cached images
./anifetch --clear-cache   # Clear cached images
./anifetch --check-token   # Check GitHub token status
./anifetch --size 40x20    # Set custom image size (fallback)
```

## Image Display

The tool uses **dynamic terminal size detection** to display anime girls at the optimal size for your terminal. It automatically adapts the image size based on your terminal dimensions while leaving space for system information.

### Image Quality Features:
- **Block symbols** for maximum detail
- **256 colors** for rich, vibrant images
- **Ordered dithering** for sharp, clear rendering
- **Dynamic sizing** that adapts to your terminal
- **Smart padding** to ensure proper layout

### Recommended Setup:
- Install `chafa` for the best image quality
- Use a terminal with good color support
- The tool will automatically detect your terminal size and optimize the image accordingly

If your terminal doesn't support image display, the tool will still show all system information with an ASCII art fallback.

## Architecture

The project is organized into modular packages:

- **`pkg/system/`** - System information gathering
- **`pkg/anime/`** - Anime image fetching and caching
- **`pkg/display/`** - Image display and rendering
- **`pkg/config/`** - Configuration management

## Configuration

### Cache Management

Images are cached in `~/.anifetch/` directory. You can:

```bash
# View cached images
./anifetch --show-cache

# Clear all cached images
./anifetch --clear-cache

# Manually delete cache directory
rm -rf ~/.anifetch/
```

### Environment Variables

- `GITHUB_TOKEN` - Set your GitHub Personal Access Token for higher rate limits

### Customization

- Use `--size` to set custom image dimensions
- Use `--no-image` to disable image display
- The tool automatically adapts to your terminal size

## Dependencies

- Go standard library
- `golang.org/x/term` for terminal size detection
- `chafa` (recommended) for optimal image display

## License

This project is open source and available under the MIT License.
