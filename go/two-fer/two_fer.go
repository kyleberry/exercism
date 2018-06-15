// Package twofer returns a string of sharing
package twofer

import "strings"

// ShareWith tests for blank input and sets `name` to "you", other wise pass
// `name` along to construct our string
func ShareWith(name string) string {
	if strings.TrimSpace(name) == "" {
		name = "you"
	}
	sharingWith := "One for " + name + ", one for me."
	return sharingWith
}
