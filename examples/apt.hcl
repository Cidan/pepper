apt install base_system {
  allow_no_version = true
  packages = [
    "htop",
    "atop",
    ]
}

apt install dep {
  requires = "apt.install.other_stuff"
}

apt install other_stuff {
  requires = "apt.install.base_system"
  packages = ["test"]
}

apt install toroot{}
