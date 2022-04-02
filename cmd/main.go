package main

import (
	"log"

	"github.com/shkurenko-sa/hcount_go/internal/server"
)

func main() {
	s := server.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
