package main

import "testing"
import "regexp"

func CheckVersionName(t *testing.T, name string) {
	re := regexp.MustCompile(`^go[0-9]+\.[0-9]+(\.[0-9]+)?$`)

	if !re.MatchString(name) {
		t.Fatalf("Version name is invalid: %s", name)
	}
}

func TestAvailableVersions(t *testing.T) {
	versions := availableVersions()

	for _, name := range versions {
		CheckVersionName(t, name)
	}
}

func TestLatestVersion(t *testing.T) {
	CheckVersionName(t, latestVersion())
}

func TestLatestVersionUrl(t *testing.T) {
	url := latestVersionUrl()
	re := regexp.MustCompile(`go[0-9\.]{1,11}.\S+-\S+\.tar\.gz$`)

	if !re.MatchString(url) {
		t.Fatalf("Latest version URL is invalid: %s", url)
	}
}
