package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":6000", server))
}
