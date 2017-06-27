# gobu
[![Build Status](https://travis-ci.org/dz0ny/gobu.svg?branch=master)](https://travis-ci.org/dz0ny/gobu)
[![Go Report Card](https://goreportcard.com/badge/github.com/dz0ny/gobu)](https://goreportcard.com/report/github.com/dz0ny/gobu)

Painless bootstrapping of GOlang. It downloads and installs GO for your OS
and architecture, sets GOROOT, GOPATH and runs default shell or your command.
```
Usage of gobu:
usage: gobu [<flags>] <command> [<args> ...]

Bootstrap your GOlang enviroment

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --debug    Verbose mode.
  --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  shell* [<flags>]
    Start a shell with Golang enviroment.

  versions
    List of all supported Golang releases.



```

You can find compiled binaries for your platform under "Releases" or if you prefer quick install:

```
$ curl -L https://github.com/dz0ny/gobu/releases/download/v0.7.3/gobu-`uname -s`-`uname -m` > /usr/local/bin/gobu
$ chmod +x /usr/local/bin/gobu
```

## Example starting new project

```
$ mkdir -p project/src/hello_world
$ cd project
$ gobu
2016/08/25 09:55:52 >> You are now in a new GOBU shell. To exit, type 'exit'
2016/08/25 09:55:52 GOPATH /home/dz0ny/project
2016/08/25 09:55:52 GOROOT /home/dz0ny/.gobu/1.7/go
$ nano src/hello_world/main.go

package main

import "fmt"
func main() {
    fmt.Println("hello world")
}

$ go build hello_world
$ tree
.
├── hello_world
└── src
    └── hello_world
        └── main.go

```


## What it does
Downloads latest version from releases and runs it in shell. After the GO is
bootstrapped you will end up in shell with "Vendoring" enabled and with fresh
env set up for you.

## Why
Because I hate setting environment every time for every project(not library).
