package main

import (
	"fmt"
	"os"
)

// Env variable
var commandPrefix = getEnv("prefix", "/")

func main() {
	fmt.Println("Starting server...")
	startListening()
}

// Useful to define a default value to getEnv
func getEnv(key, def string) string {
	resp := os.Getenv(key)
	if resp != "" {
		return resp
	}
	return def
}
