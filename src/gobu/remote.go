package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func download(version string) {
	local := filepath.Join(envPath, version+".tar.gz")
	target := filepath.Join(envPath, version)
	if _, err := os.Stat(local); os.IsNotExist(err) {
		log.Printf("Local path: %s", local)
		log.Printf("Online path: %s", onlinePath)
		log.Printf("Target path: %s", target)
		out, _ := os.Create(local)
		defer out.Close()
		resp, err := http.Get(onlinePath)
		if resp.StatusCode > 400 {
			log.Fatalf("go toolchain download: %s is not there", onlinePath)
		}
		if err != nil {
			log.Fatalf("go toolchain download: %v", err)
		}
		defer resp.Body.Close()
		io.Copy(out, resp.Body)

		if runtime.GOOS == "windows" {
			unzip(local, target)
		} else {
			untar(local, target)
		}
	} else {
		log.Printf("Skipping download")
	}
}
