package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

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
}

func goDownloadPage() io.Reader {
	client := &http.Client{Timeout: time.Duration(5 * time.Second)}
	req, err := http.NewRequest("GET", "https://golang.org/dl/", nil)

	req.Header.Set("Accept", "text/html,application/xhtml+xml,*/*;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Gobu; +https://github.com/dz0ny/gobu)")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	(&buf).ReadFrom(resp.Body)

	return &buf
}

func availableVersions(stream io.Reader) []string {
	var line string
	var parts []string
	var versions []string

	scanner := bufio.NewScanner(stream)

	for scanner.Scan() {
		line = scanner.Text()

		if strings.Contains(line, "<div\x20class=\"toggle") &&
			strings.Contains(line, "\"\x20id=\"go") {
			parts = strings.Split(line, "\x22")

			if len(parts) >= 3 && parts[3][0:2] == "go" {
				versions = append(versions, parts[3])
			}
		}
	}

	return versions
}

func latestVersion(stream io.Reader) string {
	var versions = availableVersions(stream)

	if len(versions) > 0 {
		return versions[0]
	}

	return ""
}

func latestVersionUrl(stream io.Reader) string {
	var versions = availableVersions(stream)
	var tpl = "https://storage.googleapis.com/golang/%s.%s-%s.tar.gz"

	if len(versions) > 0 {
		return fmt.Sprintf(tpl, versions[0], runtime.GOOS, runtime.GOARCH)
	}

	return ""
}
