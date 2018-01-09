package action

// Shell action runs a set of commands on the shell
type Shell struct {
	Action
	commands []string
}

func NewShell() *Shell {
	return &Shell{}
}

func (s *Shell) Execute() {

}
