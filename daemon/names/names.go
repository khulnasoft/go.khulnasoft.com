package names // import "go.khulnasoft.com/daemon/names"

import "go.khulnasoft.com/internal/lazyregexp"

// RestrictedNameChars collects the characters allowed to represent a name, normally used to validate container and volume names.
const RestrictedNameChars = `[a-zA-Z0-9][a-zA-Z0-9_.-]`

// RestrictedNamePattern is a regular expression to validate names against the collection of restricted characters.
var RestrictedNamePattern = lazyregexp.New(`^` + RestrictedNameChars + `+$`)
