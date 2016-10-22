package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/net/html"
)

func download(version string, extension string) {

	local := filepath.Join(envPath, version+extension)
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

func versions(response *http.Response) []string {

	basePath := "//storage.googleapis.com/golang/"
	suffix := "src.tar.gz"
	versions := []string{}
	tokenizer := html.NewTokenizer(response.Body)

	for {

		next := tokenizer.Next()

		switch {

		case next == html.ErrorToken:
			return versions

		case next == html.StartTagToken:

			token := tokenizer.Token()

			if token.Data == "a" {

			tokenLoop:
				for _, element := range token.Attr {

					val := element.Val

					if strings.Contains(val, basePath) {
						if strings.Contains(val, suffix) {

							parts := strings.Split(val, "/")

							if len(parts) >= 5 {

								name := parts[4]
								name = strings.TrimLeft(name, "go")
								name = strings.TrimRight(name, "."+suffix)

								for _, v := range versions {
									if v == name {
										break tokenLoop
									}
								}
								versions = append(versions, name)
							}
						}
					}
				}
			}
		}
	}

	return versions
}

func latestVersion(versions []string) string {
	if len(versions) > 0 {
		return versions[0]
	}
	return ""
}
