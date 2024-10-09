#!/bin/sh

echo "Running golangci-lint on staged files..."

golangci-lint run --new-from-rev HEAD --fix
