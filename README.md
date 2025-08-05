# AniFetch

A neofetch-like system information tool that displays anime girls holding programming books instead of the traditional OS logo.

<!-- Screenshot will be added here -->
<!-- TODO: Add screenshot showing AniFetch output with anime girl and system info -->

## Features

- Displays system information (OS, kernel, uptime, packages, shell, CPU, memory, disk)
- Fetches random anime girls from [Anime-Girls-Holding-Programming-Books](https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books)
- **Dynamic terminal size detection** for optimal image display
- **High-quality rendering** with block symbols and 256 colors
- Caches images locally for faster runs
- Cross-platform support (Linux, macOS, Windows)

## Quick Start

```bash
# Clone and build
git clone https://github.com/robbsbro69/AniFetch.git
cd anifetch
go build -o anifetch main.go

# Install chafa for best image quality
sudo pacman -S chafa  # Arch
sudo apt install chafa # Ubuntu/Debian
brew install chafa     # macOS

# Run
./anifetch
```

## Usage

```bash
./anifetch                    # Run with image
./anifetch --no-image        # Disable image display
./anifetch --show-cache      # Show cached images
./anifetch --clear-cache     # Clear cached images
./anifetch --check-token     # Check GitHub token status
./anifetch --size 40x20      # Set custom image size
```

## GitHub Token (Optional)

For higher rate limits (5,000 vs 60 requests/hour):

```bash
export GITHUB_TOKEN="your_token_here"
```

## Troubleshooting

- **Images not showing?** Install `chafa` or try `--no-image`
- **Rate limit errors?** Set up GitHub token or use cached images
- **Wrong size?** Use `--size` to customize dimensions

## Acknowledgments

- **[Anime-Girls-Holding-Programming-Books](https://github.com/cat-milk/Anime-Girls-Holding-Programming-Books)** - Image source
- **[chafa](https://hpjansson.org/chafa/)** - Terminal image viewer
- **[neofetch](https://github.com/dylanaraps/neofetch)** - Inspiration

## License

[MIT License](LICENSE)
