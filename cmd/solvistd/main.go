package main

import (
	"log"
	"os"

	"github.com/nuronialBlock/solvist/solvist/data"
)

func main() {
	err := data.OpenDBSession(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal(err)
	}
}
