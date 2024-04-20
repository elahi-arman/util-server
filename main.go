package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elahi-arman/util-server/internal/server"
)

func main() {
	s, err := server.NewServer("up", "token")
	if err != nil {
		fmt.Print("unable to start server", err)
		os.Exit(1)
	}

	log.Fatal(http.ListenAndServe(":8080", s.GetRouter()))
}
