// Copyright 2016 The Solvist Author(s). All rights reserved.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nuronialBlock/solvist/solvist/data"
	"github.com/nuronialBlock/solvist/solvist/ui"
)

func main() {
	err := data.OpenDBSession(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", ui.NewServer())

	log.Printf("Lightening on %s", os.Getenv("ADDR"))

	err = http.ListenAndServe(os.Getenv("ADDR"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
