package states

import (
	"github.com/Cidan/pepper/action"
	"github.com/blang/semver"
)

// Apt state for handling apt installs
type Shell struct {
	AllowNoVersion bool     `mapstructure:"allow_no_version"`
	Args           []string `mapstructure:"args"`
	Cmd            string   `mapstructure:"cmd"`
	shell          *action.Shell
	installed      map[string]semver.Version // name and version
}

// Merge two apt states together
func (a *Shell) Merge(b States) {

}

func (a *Shell) Execute() {
	a.pre()
	a.generate()
	a.post()
}

// Pre runs apt update, collects installed packages
// and excludes already installed packages.
func (a *Shell) pre() {
	// TODO:
	// validate version is in packages or no version is set
	// dpkg -l
	// parse semver, add to list
}

// Generate a command line run for what actions
// will be taken.
func (a *Shell) generate() string {

	// TODO:
	// Check for installed packages/create a package cache.
	return "apt install -y htop"
}

// Post idk yet.
func (a *Shell) post() {

}
