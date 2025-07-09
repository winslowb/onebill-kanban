#!/bin/bash
#
# Description: bootstrap the golang dev environment for onebill-kanban setup
# Tags: Bill gloang kanban
# Date: 2025-07-09
#

set -e

PROJECT_NAME="onebill-kanban"
GOPATH="$HOME/projects/go"
PROJECT_PATH="$GOPATH/src/github.com/winslowb/$PROJECT_NAME"
GO_VERSION="1.22.3" # Adjust if there's a newer version

echo "[*] Installing Go..."
cd /tmp
curl -OL https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C $HOME/.local/bin -xzf go${GO_VERSION}.linux-amd64.tar.gz

export PATH=$PATH:$HOME/.local/bin/go/bin
export GOPATH=$HOME/go

echo "[*] Creating project directory..."
mkdir -p "$PROJECT_PATH"/{cmd,internal,ui,model,data,sprint}
cd "$PROJECT_PATH"

echo "[*] Initializing Go module..."
go mod init github.com/winslowb/$PROJECT_NAME

echo "[*] Installing Charm dependencies..."
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/glamour

echo "[*] Initializing Git repo..."
git init
touch README.md .gitignore
echo -e "/bin/\n/pkg/\n/vendor/" > .gitignore
git add .
git commit -m "Initial commit - project structure setup"

echo "[✓] Bootstrap complete."
echo "Project path: $PROJECT_PATH"

