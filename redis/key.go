package redis

import "strings"

// GetKey returns a key with the given key prefixes
func GetKey(key string, prefixes ...string) string {
	// Push key to the end of the prefixes string slice
	prefixes = append(prefixes, key)

	return strings.Join(prefixes, KeySeparator)
}
