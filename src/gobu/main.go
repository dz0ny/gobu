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

var goPath = ""
var envPath = ".gobu"
var globalVersion = "1.5.2"
var onlinePath = "https://storage.googleapis.com/golang/go%s.%s-%s.tar.gz"
var useVendoring = false

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
	envPath = filepath.Join(userHomeDir(), envPath)
	flag.StringVar(&envPath, "env_path", envPath, "Location of GO instalation")
	flag.StringVar(&globalVersion, "version", globalVersion, "Version of Golang you wish to use")
	flag.StringVar(&goPath, "GOPATH", goPath, "Overide GOPATH")
	flag.BoolVar(&useVendoring, "vendor", false, "Start with GO15VENDOREXPERIMENT")
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

func createStore(version string) {
	os.MkdirAll(filepath.Join(envPath, version), 0755)
	os.MkdirAll(filepath.Join(envPath, version, "local"), 0755)
}

func run(version string) {

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

	if goPath == "" {
		goPath = filepath.Join(envPath, version, "local")
	}
	goroot := filepath.Join(envPath, version, "go")

	os.Setenv("GOBU", "1")
	if useVendoring {
		os.Setenv("GO15VENDOREXPERIMENT", "1")
	}
	os.Setenv("GOPATH", goPath)
	os.Setenv("GOROOT", goroot)

	// Sometimes we want to use tools from local install
	os.Setenv("PATH", goPath+"/bin:"+goroot+"/bin:"+path)

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

func download(version string) {
	local := filepath.Join(envPath, version+".tar.gz")
	target := filepath.Join(envPath, version)
	if _, err := os.Stat(local); os.IsNotExist(err) {
		log.Printf("Local path: %s", local)
		log.Printf("Online path: %s", onlinePath)
		log.Printf("Target path: %s", target)
		out, _ := os.Create(local)
		defer out.Close()
		resp, _ := http.Get(onlinePath)
		defer resp.Body.Close()
		io.Copy(out, resp.Body)
		untar(local, target)
	} else {
		log.Printf("Skipping download")
	}
}

func main() {
	flag.Parse()

	onlinePath = fmt.Sprintf(onlinePath, globalVersion, runtime.GOOS, runtime.GOARCH)

	createStore(globalVersion)
	download(globalVersion)
	run(globalVersion)
}
