package main

import (
	"log"

	"github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/setup"
)

func main() {

	if err := run(); err != nil {
		log.Fatalf("err in main: %v", err.Error())
	}
}

func run() error {
	errs := make(chan error)
	go func() {
		errs <- setup.Main(nil, nil)
	}()
	return <-errs
}
