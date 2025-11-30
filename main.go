package main

import (
	"github.com/danrodsg/api-students/api"
	"github.com/rs/zerolog/log"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	server.Start()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")

	}

}
