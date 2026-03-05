package main

import (
	"os"

	"github.com/sudoenvx/snip/internal/api"
)

func main() {
	addr := os.Getenv("SERVER_ADDR")
	server := api.NewServer(addr)

	err := server.Start()

	if err != nil {
		return
	}
}
