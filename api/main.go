package main

import (
	"fmt"

	"go.khulnasoft.com/api/provenance"
)

// TODO: delete this when we find a better way to generate release notes.
func main() {
	fmt.Println(`This 'main' exists only to create release notes for the API.`)
	fmt.Println(provenance.GetProvenance())
}
