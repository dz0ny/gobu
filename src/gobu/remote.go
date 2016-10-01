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
		untar(local, target)
	} else {
		log.Printf("Skipping download")
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
	var versions []string = availableVersions(stream)

	if len(versions) > 0 {
		return versions[0]
	}

	return ""
}

func latestVersionUrl(stream io.Reader) string {
	var versions []string = availableVersions(stream)
	var tpl string = "https://storage.googleapis.com/golang/%s.%s-%s.tar.gz"

	if len(versions) > 0 {
		return fmt.Sprintf(tpl, versions[0], runtime.GOOS, runtime.GOARCH)
	}

	return ""
}
