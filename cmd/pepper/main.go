package main

import "github.com/Cidan/pepper/schema"

func main() {
	s := schema.New()
	s.ReadDir("./examples")
	err := s.Generate()
	if err != nil {
		panic(err)
	}
}
