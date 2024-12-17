#!/bin/bash

# Build the binary
echo "Building the binary..."
go build -o RaceC ./cmd/main.go

# Move to /usr/local/bin
echo "Moving binary to /usr/local/bin/"
sudo mv ./RaceC /usr/local/bin/

echo "Installation complete!"
