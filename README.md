# gobu
Painless bootstrapping of golang. It downloads ands installs GO,
sets GOROOT, GOPATH or GO15VENDOREXPERIMENT and runs default shell.
```
Usage of gobu:
  -env_path string
    	Location of golang instalation (default "/home/dz0ny/.gobu")
  -vendor
    	Start with GO15VENDOREXPERIMENT
  -version string
    	Version of golang you wish to use (default "1.5")
```

You can find compiled binaries for your platform under "Releases"


## What it does
Download latest version from releases and run it in shell. After the GO is
bootstrapped you will end up in shell with "Vendoring" enabled and with fresh
env setup for you.

## Why
Because I hate setting environment every time for every project.
