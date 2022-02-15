package main

import (
	"log"
	"net/http"
)

func main() {
	// Inits server with in memory db
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":6000", server))
}
