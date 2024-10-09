#!/bin/sh

# Redirect output to stderr.
exec 1>&2

echo "Running golangci-lint on staged files..."

# Get list of staged Go files
STAGED_GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')

if [ -z "$STAGED_GO_FILES" ]; then
    echo "No Go files to lint."
    exit 0
fi

# Run golangci-lint on staged files
echo "$STAGED_GO_FILES" | xargs golangci-lint run --fix --files

# Check the exit code
if [ $? -ne 0 ]; then
    echo "golangci-lint found issues. Commit aborted."
    exit 1
fi

echo "golangci-lint passed. Proceeding with commit."
exit 0