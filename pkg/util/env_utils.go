package util

import "os"

//GetEnv parses an environment variable, and has a default fallback value if the environment variable
//is not set or is blank
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
