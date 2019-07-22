package main

import (
	"flag"
	"log"

	"github.com/MISTikus/ChatilGo/api"
)

func main() {
	startServer := flag.Bool("server", false, "For starting new server")
	if *startServer {
		log.Println("Starting server of ChatilGo...")
		server := api.NewServer()
		server.Start()
	}
}
