# gobu
[![wercker status](https://app.wercker.com/status/db2136ecdcb6c98f23d442af3d42e7d8/m "wercker status")](https://app.wercker.com/project/bykey/db2136ecdcb6c98f23d442af3d42e7d8)
[![Go Report Card](https://goreportcard.com/badge/github.com/dz0ny/gobu)](https://goreportcard.com/report/github.com/dz0ny/gobu)

Painless bootstrapping of golang. It downloads and installs GO for your OS
and arhitecture, sets GOROOT, GOPATH or GO15VENDOREXPERIMENT and runs default
 shell.
```
Usage of ./gobu-linux-amd64:
  -GOPATH string
    	Overide GOPATH
  -env_path string
    	Location of GO instalation (default "/home/dz0ny/.gobu")
  -vendor
    	Start with GO15VENDOREXPERIMENT
  -version string
    	Version of Golang you wish to use (default "1.5.2")
```

You can find compiled binaries for your platform under "Releases"

## Example starting new project

```
mkdir -p project/src/hello_world
cd project
gobu -GOPATH $(pwd) -vendor

cat src/hello_world/main.go

package main

import "fmt"
func main() {
    fmt.Println("hello world")
}

go build hello_world

```


## What it does
Downloads latest version from releases and runs it in shell. After the GO is
bootstrapped you will end up in shell with "Vendoring" enabled and with fresh
env set up for you.

## Why
Because I hate setting environment every time for every project.
