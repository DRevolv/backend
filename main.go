package main

import "fmt"

// These variables will be set during build time via ldflags
var (
	Version     = "dev"
	Environment = "local"
)

func main() {
	fmt.Printf("DRevolv Backend\n")
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Environment: %s\n", Environment)
	fmt.Println("Server starting...")

	fmt.Println("Hello, World!")
}
