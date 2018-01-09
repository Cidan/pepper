package states

type States interface {
	Merge(States)
	Pre()
	Generate() string
	Post()
}
