package remote

import (
	"errors"
	"fmt"
	"gobu/archive"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
	req "github.com/levigross/grequests"
)

// Version holds status about Golang download
type Version struct {
	remoteURL string
	name      string
	hash      string
	os        string
	arch      string
	beta      bool
	Release   string
}

func (v *Version) parseVersion() error {
	// `go1.7.3.windows-amd64.zip`
	var re = regexp.MustCompile(`go([\.\d\w]+)\.(\w+)-(\w+)\.`)

	for _, match := range re.FindAllStringSubmatch(v.name, 1) {
		v.Release = match[1]
		v.os = match[2]
		v.arch = match[3]
		return nil
	}
	return errors.New("Failed parsing version")
}

// Compatible filter for compatible versions
func (v *Version) Compatible() bool {

	// handle different naming for arm
	os := runtime.GOOS
	if os == "arm" {
		os = "armv6l"
	}
	isOS := os == v.os
	isARCH := runtime.GOARCH == v.arch
	return isOS && isARCH
}

// String interfacer for Version struct
func (v *Version) String() string {
	return fmt.Sprintf("%s - %s", v.Release, v.remoteURL)
}

// Setup download selected version to target dir and unpacks it
func (v *Version) Setup(rootDir string) error {
	sdkRoot := filepath.Join(rootDir, v.Release)
	os.MkdirAll(sdkRoot, 0755)
	targetFile := filepath.Join(rootDir, v.name)

	if _, err := os.Stat(targetFile); os.IsNotExist(err) {
		resp, err := req.Get(
			v.remoteURL,
			&req.RequestOptions{
				UserAgent: "Mozilla/5.0 (compatible; Gobu; +https://github.com/dz0ny/gobu)",
			},
		)
		if err != nil {
			return err
		}

		log.Infoln("Starting download of ", v.String())
		err = resp.DownloadToFile(targetFile)
		if err != nil {
			return err
		}
		log.Infoln("Starting extraction of ", targetFile)
		if strings.HasSuffix(v.name, ".zip") {
			return archive.Unzip(targetFile, sdkRoot)
		}
		if strings.HasSuffix(v.name, ".tar.gz") {
			return archive.Untar(targetFile, sdkRoot)
		}
		return errors.New("Unsuported archive type")

	}
	return nil
}
