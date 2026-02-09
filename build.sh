#!/bin/bash

# Cross-platform build script for ComplianceGuard
# Usage: ./build.sh [all|platform]

set -e

BINARY_NAME="complianceguard"
VERSION="1.0.0"
BUILD_DIR="build"

# Create build directory
mkdir -p "$BUILD_DIR"

# Build flags
LDFLAGS="-s -w -X complianceguard/cmd.version=$VERSION"

# Function to build for a specific platform
build_platform() {
    local os=$1
    local arch=$2
    local output="${BUILD_DIR}/${BINARY_NAME}-${os}-${arch}"
    
    if [ "$os" = "windows" ]; then
        output="${output}.exe"
    fi
    
    echo "Building for $os/$arch..."
    GOOS=$os GOARCH=$arch go build -ldflags="$LDFLAGS" -o "$output" .
    
    if [ $? -eq 0 ]; then
        echo "✓ Built: $output"
    else
        echo "✗ Failed to build for $os/$arch"
        return 1
    fi
}

# Main build logic
if [ "$1" = "all" ] || [ -z "$1" ]; then
    echo "Building for all platforms..."
    echo "================================"
    
    # macOS
    build_platform darwin amd64
    build_platform darwin arm64
    
    # Linux
    build_platform linux amd64
    build_platform linux arm64
    
    # Windows
    build_platform windows amd64
    build_platform windows arm64
    
    echo "================================"
    echo "All builds completed successfully!"
    echo ""
    echo "Build artifacts:"
    ls -lh "$BUILD_DIR"
    
elif [ "$1" = "current" ]; then
    echo "Building for current platform..."
    go build -ldflags="$LDFLAGS" -o "${BUILD_DIR}/${BINARY_NAME}" .
    echo "✓ Built: ${BUILD_DIR}/${BINARY_NAME}"
    
else
    # Build for specific platform (e.g., ./build.sh linux amd64)
    if [ -z "$2" ]; then
        echo "Error: Please specify both OS and architecture"
        echo "Usage: $0 <os> <arch>"
        echo "Example: $0 linux amd64"
        exit 1
    fi
    
    build_platform "$1" "$2"
fi
