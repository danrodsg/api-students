package main

import (
	"log"

	"github.com/danrodsg/api-students/api"

)

func main() {
	
	server := api.NewServer()

	server.ConfigureRoutes()

	server.Start()

	if err := server.Start(); err != nil {
		log.Fatal(err)

	}

}
