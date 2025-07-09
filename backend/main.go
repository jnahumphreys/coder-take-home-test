package main

import (
	"fmt"
	"log"
	"time"

	"github.com/coder/registry-take-home/server"
)

const (
	port = "8080"
)

func main() {
	// Initialize the database
	db := server.NewDB()

	// Start the daemon in the background
	go server.RunDaemon(server.DaemonOptions{
		DB:           db,
		InitialCount: 1000,
		Interval:     2 * time.Second,
	})

	// Create and start the server
	server := server.NewServer(db)
	fmt.Printf("Server starting on :%s\n", port)
	err := server.Listen(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
