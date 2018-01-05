package states

type Apt struct {
	Requires string `hcl:"requires"`
}

func NewApt() *Apt {
	return &Apt{}
}
