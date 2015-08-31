package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

var env_path = ".go15env"
var global_version = "1.5"
var online = "https://storage.googleapis.com/golang/go%s.%s-%s.tar.gz"

func init() {
	usr, _ := user.Current()
	env_path = filepath.Join(usr.HomeDir, env_path)
	online = fmt.Sprintf(online, global_version, runtime.GOOS, runtime.GOARCH)
}

func CreateStore(version string) {
	os.MkdirAll(filepath.Join(env_path, version), 0755)
	os.MkdirAll(filepath.Join(env_path, version, "local"), 0755)
}

func Run(version string) {

	if os.Getenv("GO15ENV") != "" {
		log.Fatalln("Already boostraped env!")
	}

	shell := os.Getenv("SHELL")
	cmdArgs := []string{shell}
	path := os.Getenv("PATH")

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	gopath := filepath.Join(env_path, version, "local")
	goroot := filepath.Join(env_path, version, "go")

	os.Setenv("GO15ENV", "1")
	os.Setenv("GO15VENDOREXPERIMENT", "1")
	os.Setenv("GOPATH", gopath)
	os.Setenv("GOROOT", goroot)
	os.Setenv("PATH", gopath+"/bin:"+goroot+"/bin:"+path)

	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	fmt.Printf(">> You are now in a new go15env shell. To exit, type 'exit'\n")

	if !filepath.IsAbs(shell) {
		shell, err = exec.LookPath(shell)
		if err != nil {
			log.Panic(err)
		}
	}

	// Login may work better than executing the shell manually.
	//proc, err := os.StartProcess(loginPath, []string{"login", "-fpl", u.Username}, &pa)
	proc, err := os.StartProcess(shell, cmdArgs, &pa)
	if err != nil {
		log.Panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Exited go15env shell: %s", state.String())
}

func Download(version string) {
	local := filepath.Join(env_path, version+".tar.gz")
	target := filepath.Join(env_path, version)

	log.Printf("Local path: %s", local)
	log.Printf("Online path: %s", online)
	log.Printf("Target path: %s", target)
	if _, err := os.Stat(local); os.IsNotExist(err) {
		out, _ := os.Create(local)
		defer out.Close()
		resp, _ := http.Get(online)
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
		cmd := exec.Command("tar", "xvf", local, "-C", target)
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Wait()
	}
}

func main() {
	CreateStore(global_version)
	Download(global_version)
	Run(global_version)
}
