#!/bin/bash

# Exit on error, uninitialized variables, and error in a pipeline
set -euo pipefail

# Function to generate the HTML redirect for the module
module() {
    local github_repo_name="$1"
    local module_root="${2:-$github_repo_name}"
    local module_subdir="${3:-}"
    local module_path="${module_root}${module_subdir:+/$module_subdir}"

    mkdir -p "$module_path"
    
    # Generate the HTML file
    local github_url="https://github.com/khulnasoft-lab/$github_repo_name"
    printf '%s\n' \
    "<!doctype html>" \
    "<html lang=\"en\">" \
    "<head>" \
    "    <title>go.khulnasoft.com/lab/$module_path</title>" \
    "    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/>" \
    "    <meta name=\"go-import\" content=\"go.khulnasoft.com/lab/$module_root git $github_url\">" \
    "    <meta http-equiv=\"refresh\" content=\"0; url=$github_url\">" \
    "</head>" \
    "<body>" \
    "    Redirecting to <a href=\"$github_url\">$github_url</a>" \
    "</body>" \
    "</html>" > "$module_path/index.html"

    echo "Generated module for $github_repo_name at $module_path/index.html"
}

# Main script execution
if [[ $# -lt 1 ]]; then
    echo "Usage: ./gen.sh <repo_name> [module_root] [module_subdir]"
    exit 1
fi

# Call the module function with the provided arguments
module "$@"
