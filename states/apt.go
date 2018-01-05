package states

type Apt struct {
	State
	requires []string
}

func NewApt() *Apt {
	return &Apt{}
}
