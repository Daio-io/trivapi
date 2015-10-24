package utils

import "os"

// GetPort - Get the port number set
// by the environment PORT
// or return default 9000
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	return port
}
