package states

import (
	"fmt"
	"os/exec"

	"github.com/Cidan/pepper/action"
	"github.com/blang/semver"
	"github.com/rs/zerolog/log"
)

// Apt state for handling apt installs
type Apt struct {
	AllowNoVersion bool     `mapstructure:"allow_no_version"`
	Packages       []string `mapstructure:"packages"`
	shell          *action.Shell
	installed      map[string]semver.Version // name and version
	cmd            string
}

// Merge two apt states together
func (a *Apt) Merge(b States) {

}

func (a *Apt) Execute() {
	a.pre()
	err := a.run()
	if err != nil {
		panic(err)
	}
	a.post()
}

// Pre runs apt update, collects installed packages
// and excludes already installed packages.
func (a *Apt) pre() {
	log.Info().Msg("Updating APT")
	cmd := exec.Command("apt-get", "update")
	cmd.Run()
}

// Generate a command line run for what actions
// will be taken.
func (a *Apt) run() error {

	// TODO: install, remove, purge, update options
	// root@jinked:/home/alobato/go/src/github.com/Cidan/pep
	// -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold"
	log.Info().Strs("packages", a.Packages).Msg("Installing packages")
	args := append([]string{
		"-q",
		"-y",
		"--force-yes",
		"-o",
		"Dpkg::Options::=\"--force-confdef\"",
		"-o",
		"Dpkg::Options::=\"--force-confold\"",
		"-o",
		"Dpkg::Options::=\"-i\"",
		"install",
	}, a.Packages...)
	log.Debug().Strs("args", args).Msg("apt args")
	cmd := exec.Command("apt-get", args...)
	//cmd.Env = []string{"DEBIAN_FRONTEND=noninteractive"}
	b, err := cmd.CombinedOutput()
	fmt.Printf("output: %s\n", string(b))

	return err
	// TODO:
	// validate version is in packages or no version is set
	// dpkg -l
	// parse semver, add to list

	// TODO:
	// Check for installed packages/create a package cache.
}

// Post idk yet.
func (a *Apt) post() {

}
