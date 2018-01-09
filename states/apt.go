package states

// Apt state for handling apt installs
type Apt struct {
	Packages []string `mapstructure:"packages"`
}

// Merge two apt states together
func (a *Apt) Merge(b States) {

}

// Pre runs apt update, collects installed packages
// and excludes already installed packages.
func (a *Apt) Pre() {

}

// Generate a command line run for what actions
// will be taken.
func (a *Apt) Generate() string {

	// TODO:
	// Check for installed packages/create a package cache.
	return "apt install -y htop"
}

// Post idk yet.
func (a *Apt) Post() {

}
