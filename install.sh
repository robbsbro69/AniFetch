#!/bin/bash

# AniFetch Installation Script

echo "🔍 Checking prerequisites..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed!"
    echo ""
    echo "Please install Go first:"
    echo "  Ubuntu/Debian: sudo apt install golang-go"
    echo "  Arch Linux:    sudo pacman -S go"
    echo "  macOS:         brew install go"
    echo ""
    echo "After installing Go, run this script again."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | grep -o 'go[0-9]\+\.[0-9]\+' | cut -c3-)
REQUIRED_VERSION="1.21"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo "❌ Go version $GO_VERSION is too old!"
    echo "Please install Go 1.21 or later."
    exit 1
fi

echo "✅ Go $(go version | grep -o 'go[0-9]\+\.[0-9]\+') found"

# Check if chafa is installed (optional)
if command -v chafa &> /dev/null; then
    echo "✅ chafa found (for best image quality)"
else
    echo "⚠️  chafa not found (optional but recommended)"
    echo "   Install with: sudo apt install chafa (Ubuntu/Debian)"
    echo "   Install with: sudo pacman -S chafa (Arch)"
    echo "   Install with: brew install chafa (macOS)"
fi

echo ""
echo "🔨 Building AniFetch..."
go build -o anifetch main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    
    # Check if user wants to install globally
    echo ""
    read -p "Do you want to install AniFetch globally? (y/n): " -n 1 -r
    echo
    
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "📦 Installing AniFetch to /usr/local/bin/..."
        sudo cp anifetch /usr/local/bin/
        
        if [ $? -eq 0 ]; then
            echo "✅ AniFetch installed successfully!"
            echo "🚀 You can now run 'anifetch' from anywhere."
        else
            echo "❌ Installation failed. You can still run './anifetch' from this directory."
        fi
    else
        echo "✅ AniFetch built successfully! Run './anifetch' to use it."
    fi
else
    echo "❌ Build failed. Please check the error messages above."
    exit 1
fi 