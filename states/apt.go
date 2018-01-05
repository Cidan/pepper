package states

type Apt struct {
	Packages []string `mapstructure:"packages"`
}

func (a *Apt) Generate() string {
	return ""
}
