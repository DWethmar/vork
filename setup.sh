#!/bin/bash

# This script installs the pre-commit hook from scripts/lint.sh

# Exit immediately if any command exits with a non-zero status
set -e

# Define the path to the Git hooks directory
HOOKS_DIR=".git/hooks"
PRE_COMMIT_HOOK="$HOOKS_DIR/pre-commit"

# Check if the script is being run from the root of the Git repository
if [ ! -d "$HOOKS_DIR" ]; then
    echo "Error: .git/hooks directory not found."
    echo "Please run this script from the root of your Git repository."
    exit 1
fi

# Copy the lint.sh script to the pre-commit hook
cp scripts/lint.sh "$PRE_COMMIT_HOOK"

# Make the pre-commit hook executable
chmod +x "$PRE_COMMIT_HOOK"

echo "Pre-commit hook installed successfully."