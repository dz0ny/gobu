package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var env_path = ".gobu"
var global_version = "1.5"
var online_path = "https://storage.googleapis.com/golang/go%s.%s-%s.tar.gz"
var use_vendoring = false

func UserHomeDir() string {
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
	env_path = filepath.Join(UserHomeDir(), env_path)
	flag.StringVar(&env_path, "env_path", env_path, "Location of GO instalation")
	flag.StringVar(&global_version, "version", "1.5", "Version of Golang you wish to use")
	flag.BoolVar(&use_vendoring, "vendor", false, "Start with GO15VENDOREXPERIMENT")
}

func untar(source, target string) {

	file, err := os.Open(source)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(source, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			log.Fatal(err)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// get the individual filename and extract to the current directory
		filename := filepath.Join(target, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			// handle directory
			log.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				log.Fatal(err)
			}

		case tar.TypeReg:
			// handle normal file
			log.Println("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				log.Fatal(err)
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				log.Fatal(err)
			}

			writer.Close()
		default:
			log.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}

}

func CreateStore(version string) {
	os.MkdirAll(filepath.Join(env_path, version), 0755)
	os.MkdirAll(filepath.Join(env_path, version, "local"), 0755)
}

func Run(version string) {

	if os.Getenv("GOBU") != "" {
		log.Fatalln("Already in boostraped env!")
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

	os.Setenv("GOBU", "1")
	if use_vendoring {
		os.Setenv("GO15VENDOREXPERIMENT", "1")
	} else {
		os.Setenv("GOPATH", gopath)
	}

	os.Setenv("GOROOT", goroot)

	// Sometimes we want to use tools from local install
	os.Setenv("PATH", gopath+"/bin:"+goroot+"/bin:"+path)

	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	log.Printf(">> You are now in a new GOBU shell. To exit, type 'exit'\n")

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

	log.Printf("Exited gobu shell: %s", state.String())
}

func Download(version string) {
	local := filepath.Join(env_path, version+".tar.gz")
	target := filepath.Join(env_path, version)
	if _, err := os.Stat(local); os.IsNotExist(err) {
		log.Printf("Local path: %s", local)
		log.Printf("Online path: %s", online_path)
		log.Printf("Target path: %s", target)
		out, _ := os.Create(local)
		defer out.Close()
		resp, _ := http.Get(online_path)
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
		untar(local, target)
	} else {
		log.Printf("Skipping download")
	}
}

func main() {
	flag.Parse()

	online_path = fmt.Sprintf(online_path, global_version, runtime.GOOS, runtime.GOARCH)

	CreateStore(global_version)
	Download(global_version)
	Run(global_version)
}
