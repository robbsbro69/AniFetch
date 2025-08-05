#!/bin/bash

# AniFetch Installation Script

echo "Building AniFetch..."
go build -o anifetch main.go

if [ $? -eq 0 ]; then
    echo "Build successful!"
    
    # Check if user wants to install globally
    read -p "Do you want to install AniFetch globally? (y/n): " -n 1 -r
    echo
    
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo "Installing AniFetch to /usr/local/bin/..."
        sudo cp anifetch /usr/local/bin/
        
        if [ $? -eq 0 ]; then
            echo "AniFetch installed successfully!"
            echo "You can now run 'anifetch' from anywhere."
        else
            echo "Installation failed. You can still run './anifetch' from this directory."
        fi
    else
        echo "AniFetch built successfully! Run './anifetch' to use it."
    fi
else
    echo "Build failed. Please check the error messages above."
    exit 1
fi 