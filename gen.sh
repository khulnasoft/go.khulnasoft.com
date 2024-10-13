#!/bin/bash

# Usage: ./gen.sh <repo_name>
# Calls the module function to add the repo.

repo_name="$1"

# Optional: You can fetch the full repo data from GitHub here if needed
module "$repo_name"
