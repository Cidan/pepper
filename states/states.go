package states

type States interface {
	Merge(States)
	Execute()
}
