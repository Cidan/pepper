apt install base_system {
  packages = [
    "htop",
    "atop",
    "ntop", # All the tops
    "top"
    ]
}

apt install dep {
  requires = ["apt.install.other_stuff"]
}

apt install other_stuff {
  requires = "apt.install.base_system"
}

apt install toroot{}