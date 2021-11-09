// Package twofer provides a function that returns a string
package twofer

import (
	"fmt"
)

// ShareWith creates a String with the name of the person and "you" if no name is provided.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
