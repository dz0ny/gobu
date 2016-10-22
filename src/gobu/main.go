package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gpahal/shlex"
	"github.com/mitchellh/go-ps"
)

var goPath = ""
var execCmd = ""
var envPath = ".gobu"
var available = false
var globalVersion = "1.7"
var onlinePath = "https://storage.googleapis.com/golang/go%s.%s-%s%s"
var unixExtension = ".tar.gz"
var winExtension = ".zip"

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

func init() {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	goPath = p
	envPath = filepath.Join(userHomeDir(), envPath)

	flag.StringVar(&envPath, "env_path", envPath, "Location of GO instalation")
	flag.StringVar(&globalVersion, "version", globalVersion, "Version of Golang you wish to use")
	flag.StringVar(&goPath, "GOPATH", goPath, "Overide GOPATH")
	flag.StringVar(&execCmd, "exec", "", "Run command instead of default shell")
	flag.BoolVar(&available, "available", false, "List available stable versions")
}

func createStore(version string) {
	os.MkdirAll(filepath.Join(envPath, version), 0755)
	os.MkdirAll(filepath.Join(envPath, version, "local"), 0755)
}

func runShell(version string) {
	if os.Getenv("GOBU") != "" {
		log.Fatalln("Already in boostraped env!")
	}
	os.Setenv("GOBU", "1")

	log.Println(">> You are now in a new GOBU shell. To exit, type 'exit'")

	shell := os.Getenv("SHELL")

	if runtime.GOOS == "windows" {

		parentId := os.Getppid()

		parentProcess, err := ps.FindProcess(parentId)

		if err != nil {
			log.Panic(err)
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

	os.Setenv("GO15VENDOREXPERIMENT", "1")
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

func manageLatestVersion(globalVersion *string) {
	url := "https://golang.org/dl/"
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	resp, _ := client.Get(url)
	versions := versions(resp)
	currentVersion := latestVersion(versions)
	if currentVersion != "" {
		*globalVersion = currentVersion
	}
}

func main() {
	flag.Parse()
	arch := runtime.GOARCH
	// Fix for special case of arm version naming
	if arch == "arm" {
		arch = "armv6l"
	}

	manageLatestVersion(&globalVersion)

	extension := unixExtension

	if runtime.GOOS == "windows" {
		extension = winExtension
	}

	onlinePath = fmt.Sprintf(onlinePath, globalVersion, runtime.GOOS, arch, extension)

	if available {
		versions := availableVersions(goDownloadPage())
		for _, version := range versions {
			fmt.Println(version[2:])
		}
		return
	}

	createStore(globalVersion)
	download(globalVersion, extension)
	if execCmd != "" {
		execCmdParsed, err := shlex.Split(execCmd)
		if err != nil {
			log.Panic(err)
		}
		bin := resolveBinary(execCmdParsed[0])
		run(globalVersion, bin, execCmdParsed[:1])
	} else {
		runShell(globalVersion)
	}

}
