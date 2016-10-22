package main

import (
	"bytes"
	"reflect"
	"regexp"
	"testing"
)

var content = `
	<div class="toggle" id="go1.2.3">
	<div class="toggleVisible" id="go1.2.4">
	<div class="toggle" id="go1.2.5">
	<div class="toggleVisible" id="go2.0.3">
	<div class="toggle" id="go3.0">`

func CheckVersionName(t *testing.T, name string) {
	re := regexp.MustCompile(`^go[0-9]+\.[0-9]+(\.[0-9]+)?$`)

	if !re.MatchString(name) {
		t.Fatalf("Version name is invalid: %s", name)
	}
}

func TestAvailableVersions(t *testing.T) {
	stream := bytes.NewBufferString(content)
	versions := availableVersions(stream)
	expected := []string{
		"go1.2.3",
		"go1.2.4",
		"go1.2.5",
		"go2.0.3",
		"go3.0",
	}

	for _, name := range versions {
		CheckVersionName(t, name)
	}

	if !reflect.DeepEqual(versions, expected) {
		t.Fatalf("Incorrect extracted versions: %s", versions)
	}
}

func TestLatestVersion(t *testing.T) {
	stream := bytes.NewBufferString(content)
	CheckVersionName(t, latestVersion(stream))
}

func TestLatestVersionUrl(t *testing.T) {
	stream := bytes.NewBufferString(content)
	url := latestVersionUrl(stream)
	re := regexp.MustCompile(`go[0-9\.]{1,11}.\S+-\S+\.tar\.gz$`)

	if !re.MatchString(url) {
		t.Fatalf("Latest version URL is invalid: %s", url)
	}
}
