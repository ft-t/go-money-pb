#!/bin/sh

# Set safe directory for Git
git config --global --add safe.directory /app/.git

# Execute the original buf-breaking command
exec buf "$@"
