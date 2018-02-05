package main

import "github.com/Cidan/pepper/plan"

func main() {
	p := plan.New()
	p.ReadDir("./examples")
	err := p.Generate()
	p.Execute()
	if err != nil {
		panic(err)
	}
}
