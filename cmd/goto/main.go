package main

import (
	"github.com/jhachmer/gotothemovies/pkg/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "goto:", log.LstdFlags)
	svr := server.NewServer(":8080", logger)
	svr.Serve()
}
