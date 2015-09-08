# gobu
[![wercker status](https://app.wercker.com/status/db2136ecdcb6c98f23d442af3d42e7d8/m "wercker status")](https://app.wercker.com/project/bykey/db2136ecdcb6c98f23d442af3d42e7d8)

Painless bootstrapping of golang. It downloads and installs GO,
sets GOROOT, GOPATH or GO15VENDOREXPERIMENT and runs default shell.
```
Usage of gobu:
  -env_path string
    	Location of golang installation (default "/home/dz0ny/.gobu")
  -vendor
    	Start with GO15VENDOREXPERIMENT
  -version string
    	Version of golang you wish to use (default "1.5")
```

You can find compiled binaries for your platform under "Releases"


## What it does
Downloads latest version from releases and runs it in shell. After the GO is
bootstrapped you will end up in shell with "Vendoring" enabled and with fresh
env set up for you.

## Why
Because I hate setting environment every time for every project.
