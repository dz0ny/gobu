package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/dz0ny/gobu/version"

	"github.com/dz0ny/gobu/remote"

	ps "github.com/mitchellh/go-ps"
	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var app = kingpin.New("gobu", "Bootstrap your golang environment")
var showVerbose = app.Flag("debug", "Verbose mode.").Bool()
var mod = app.Flag("mod", "Enable modules(GO111MODULE=on)").Bool()

var shell = app.Command("shell", "Start a shell with Golang environment.").Default()
var shellVersion = shell.Flag("release", "Override Golang version used in new shell").String()

var goPath = ""
var envPath = ".gobu"

func init() {

	app.Command("versions", "List all remote golang versions")

	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	goPath = p
	envPath = filepath.Join(userHomeDir(), envPath)
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func runShell(version string) {
	if os.Getenv("GOBU") != "" {
		log.Fatalln("Already in boostraped env!")
	}
	os.Setenv("GOBU", "1")

	log.Println(">> You are now in a new GOBU shell. To exit, type 'exit'")

	shell := os.Getenv("SHELL")

	if runtime.GOOS == "windows" {

		parentID := os.Getppid()

		parentProcess, err := ps.FindProcess(parentID)

		if err != nil {
			log.Fatalln(err)
		}

		shell = parentProcess.Executable()
	}

	shellBinary := resolveBinary(shell)

	run(version, shellBinary, []string{shellBinary})
	log.Println("Exited gobu shell")
}

func resolveBinary(bin string) string {
	if !filepath.IsAbs(bin) {
		rbin, err := exec.LookPath(bin)
		if err != nil {
			log.Panic(err)
		}
		return rbin
	}
	return bin
}

func run(version, cmd string, cmdArgs []string) *os.ProcessState {

	path := os.Getenv("PATH")
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	goroot := filepath.Join(envPath, version, "go")

	if *mod {
		os.Setenv("GO111MODULE", "on")
		goPath = filepath.Join(envPath, "global")
		log.Println("Go Module support is enabled")
	}

	os.Setenv("GOPATH", goPath)
	log.Println("GOPATH", goPath)
	os.Setenv("GOROOT", goroot)
	log.Println("GOROOT", goroot)

	// Sometimes we want to use tools from local install
	os.Setenv("PATH", goPath+"/bin:"+goroot+"/bin:"+path)

	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	proc, err := os.StartProcess(cmd, cmdArgs, &pa)
	if err != nil {
		log.Panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		log.Panic(err)
	}
	return state
}

func main() {
	app.Author("dz0ny")
	app.Version(version.String())
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	log.SetFormatter(&log.TextFormatter{DisableTimestamp: true})

	if *showVerbose {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug mode enabled")
	}

	r := remote.Remote{}

	switch cmd {
	case "versions":
		r.Update()
		fmt.Println("List of supported Go lang releases for this platform:")
		for _, v := range r.Versions {
			fmt.Println(v.String())
		}
	case "shell":

		if *shellVersion != "" {
			r.Update()
			selected, err := r.GetVersion(*shellVersion)
			if err != nil {
				log.Fatalln(err)
			}
			if err := selected.Setup(envPath); err != nil {
				log.Fatalln(err)
			}
			runShell(selected.Release)
		} else {
			r.Update()
			latest := r.LatestVersion()
			if err := latest.Setup(envPath); err != nil {
				log.Fatalln(err)
			}
			runShell(latest.Release)
		}

	}
}
