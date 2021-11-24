package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Env variable
var commandPrefix = getEnv("prefix", "/")

func main() {
	fmt.Println("Starting server...")
	startListening()
	fmt.Println("Server OK !")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// Useful to define a default value to getEnv
func getEnv(key, def string) string {
	resp := os.Getenv(key)
	if resp != "" {
		return resp
	}
	return def
}
