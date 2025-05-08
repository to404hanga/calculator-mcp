package main

import (
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/to404hanga/calculator-mcp/internal/mcp"
)

func main() {
	log.SetOutput(os.Stderr)
	log.Printf("Starting calculator-mcp server...")

	s := mcp.NewCalcServer()

	log.Printf("Starting stdio server...")
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
