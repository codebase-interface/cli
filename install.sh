# Installation script for Codebase Interface CLI
# Usage: curl -sSL https://raw.githubusercontent.com/codebase-interface/cli/main/install.sh | bash

set -e

# Default installation directory
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="codebase-interface"
REPO_OWNER="codebase-interface"
REPO_NAME="cli"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Detect OS and architecture
detect_platform() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    ARCH=$(uname -m)
    
    case $OS in
        linux)
            OS="linux"
            ;;
        darwin)
            OS="darwin"
            ;;
        msys*|mingw*|cygwin*)
            OS="windows"
            ;;
        *)
            echo -e "${RED}Error: Unsupported operating system: $OS${NC}"
            exit 1
            ;;
    esac
    
    case $ARCH in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        arm64|aarch64)
            ARCH="arm64"
            ;;
        *)
            echo -e "${RED}Error: Unsupported architecture: $ARCH${NC}"
            exit 1
            ;;
    esac
}

# Get the latest release version
get_latest_version() {
    echo -e "${BLUE}Fetching latest release...${NC}"
    LATEST_VERSION=$(curl -s "https://api.github.com/repos/${REPO_OWNER}/${REPO_NAME}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    
    if [ -z "$LATEST_VERSION" ]; then
        echo -e "${RED}Error: Could not fetch latest version${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}Latest version: $LATEST_VERSION${NC}"
}

# Download and install binary
install_binary() {
    if [ "$OS" = "windows" ]; then
        BINARY_NAME="${BINARY_NAME}.exe"
    fi
    
    DOWNLOAD_URL="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/download/${LATEST_VERSION}/${BINARY_NAME}-${OS}-${ARCH}"
    if [ "$OS" = "windows" ]; then
        DOWNLOAD_URL="${DOWNLOAD_URL}.exe"
    fi
    
    echo -e "${BLUE}Downloading from: $DOWNLOAD_URL${NC}"
    
    # Create temporary file
    TMP_FILE=$(mktemp)
    
    # Download binary
    if ! curl -L "$DOWNLOAD_URL" -o "$TMP_FILE"; then
        echo -e "${RED}Error: Failed to download binary${NC}"
        exit 1
    fi
    
    # Check if we can write to install directory
    if [ ! -w "$INSTALL_DIR" ]; then
        echo -e "${YELLOW}Warning: Cannot write to $INSTALL_DIR, trying with sudo...${NC}"
        SUDO="sudo"
    else
        SUDO=""
    fi
    
    # Install binary
    $SUDO mv "$TMP_FILE" "$INSTALL_DIR/$BINARY_NAME"
    $SUDO chmod +x "$INSTALL_DIR/$BINARY_NAME"
    
    # Create symlink for short alias
    if [ "$OS" != "windows" ]; then
        $SUDO ln -sf "$INSTALL_DIR/$BINARY_NAME" "$INSTALL_DIR/cbi"
        echo -e "${GREEN}Created alias: cbi -> $BINARY_NAME${NC}"
    fi
    
    echo -e "${GREEN}✅ Successfully installed $BINARY_NAME to $INSTALL_DIR${NC}"
}

# Verify installation
verify_installation() {
    if command -v "$BINARY_NAME" >/dev/null 2>&1; then
        VERSION_OUTPUT=$("$BINARY_NAME" version 2>/dev/null || echo "version command not available")
        echo -e "${GREEN}✅ Installation verified${NC}"
        echo -e "${BLUE}Run '$BINARY_NAME --help' or 'cbi --help' to get started!${NC}"
        
        if [ "$VERSION_OUTPUT" != "version command not available" ]; then
            echo -e "${BLUE}Installed version: $VERSION_OUTPUT${NC}"
        fi
    else
        echo -e "${YELLOW}Warning: $BINARY_NAME not found in PATH${NC}"
        echo -e "${YELLOW}You may need to add $INSTALL_DIR to your PATH${NC}"
        echo -e "${YELLOW}Or run: export PATH=\"$INSTALL_DIR:\$PATH\"${NC}"
    fi
}

# Main installation flow
main() {
    echo -e "${BLUE}🚀 Installing Codebase Interface CLI...${NC}"
    echo
    
    # Allow custom install directory
    if [ -n "$1" ]; then
        INSTALL_DIR="$1"
        echo -e "${BLUE}Custom install directory: $INSTALL_DIR${NC}"
    fi
    
    # Create install directory if it doesn't exist
    if [ ! -d "$INSTALL_DIR" ]; then
        echo -e "${BLUE}Creating install directory: $INSTALL_DIR${NC}"
        mkdir -p "$INSTALL_DIR" 2>/dev/null || {
            echo -e "${YELLOW}Creating directory with sudo...${NC}"
            sudo mkdir -p "$INSTALL_DIR"
        }
    fi
    
    detect_platform
    echo -e "${BLUE}Detected platform: $OS-$ARCH${NC}"
    
    get_latest_version
    install_binary
    verify_installation
    
    echo
    echo -e "${GREEN}🎉 Installation complete!${NC}"
    echo -e "${BLUE}Quick start:${NC}"
    echo -e "  $BINARY_NAME init-config basic    # Create a configuration file"
    echo -e "  $BINARY_NAME validate            # Validate your codebase"
    echo -e "  $BINARY_NAME --help              # See all available commands"
}

# Run main function with all arguments
main "$@"