package main

import (
	"log"

	"github.com/iriscompanyio/awslex-bot/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
