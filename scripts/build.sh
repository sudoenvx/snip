#!/usr/bin/bash
# -*- coding: utf-8 -*-

# Exit immediately if a command fails
set -e

# Configure Go environment (adjust these paths to your setup)
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Check if Go is installed
if ! command -v go.exe &> /dev/null
then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi


# Name of the output binary
APP_NAME="snip"

# Build the project (compiles all .go files in the current directory)
# echo "Building Go project..."
go.exe build -o $APP_NAME ./cmd/snip/

# Run the binary with any arguments passed to the script
# echo "Running $APP_NAME..."
./$APP_NAME "$@"