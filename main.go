package main

import (
	"github.com/wodm8/wodm8-core/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
