package schema

type State struct {
	Command string
	Schema  map[string]*Schema
}
