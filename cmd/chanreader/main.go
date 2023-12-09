package main

import (
	"github.com/glebpepega/chanreader/internal/server"
	"github.com/glebpepega/chanreader/internal/server/config"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg := config.MustLoad()

	s := server.New(cfg)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
