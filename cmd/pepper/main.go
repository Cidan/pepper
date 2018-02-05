package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Cidan/pepper/plan"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	p := plan.New()
	p.ReadDir("./examples")
	err := p.Generate()
	p.Execute()
	if err != nil {
		panic(err)
	}
}
