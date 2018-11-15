VERSION := 0.8.0
PKG := gobu
MODULE := github.com/dz0ny/gobu
COMMIT := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date -u +%FT%T)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
CURRENT_TARGET = $(PKG)-$(shell uname -s)-$(shell uname -m)
TARGETS := Linux-arm-armv7l Linux-arm-armv6l Linux-arm64-aarch64 Linux-amd64-x86_64 Darwin-amd64-x86_64 Windows-amd64-x86_64

os = $(word 1, $(subst -, ,$@))
arch = $(word 3, $(subst -, ,$@))
goarch = $(word 2, $(subst -, ,$@))
goos = $(shell echo $(os) | tr A-Z a-z)
output = $(PKG)-$(os)-$(arch)
version_flags = -X $(MODULE)/version.Version=$(VERSION) \
 -X $(MODULE)/version.CommitHash=${COMMIT} \
 -X $(MODULE)/version.Branch=${BRANCH} \
 -X $(MODULE)/version.BuildTime=${BUILD_TIME}

define localbuild
	GO111MODULE=off go get -u $(1)
	GO111MODULE=off go build $(1)
	mkdir -p bin
	mv $(2) bin/$(2)
endef

define ghupload
	bin/github-release upload \
		--user dz0ny \
		--repo $(PKG) \
		--tag "v$(VERSION)" \
		--name $(PKG)-$(1) \
		--file $(PKG)-$(1)
endef

.PHONY: $(TARGETS)
$(TARGETS):
	env CGO_ENABLED=0 GOOS=$(goos) GOARCH=$(goarch) go build -gcflags "-trimpath $(shell pwd)"  --ldflags '-s -w $(version_flags)' -o $(output)

#
# Build all defined targets
#
.PHONY: build
build: $(TARGETS)

#
# Install app for current system
#
install:
	sudo mv $(CURRENT_TARGET) /usr/local/bin/$(PKG)

bin/github-release:
	$(call localbuild,github.com/aktau/github-release,github-release)

bin/gocov:
	$(call localbuild,github.com/axw/gocov/gocov,gocov)

bin/golangci-lint:
	$(call localbuild,github.com/golangci/golangci-lint/cmd/golangci-lint,golangci-lint)

clean:
	rm -rf bin

lint: bin/golangci-lint
	bin/golangci-lint run
	go fmt

test: lint cover
	go test -v -race

cover: bin/gocov
	gocov test | gocov report

upload: bin/github-release
	$(call ghupload,Linux-armv7l)
	$(call ghupload,Linux-armv6l)
	$(call ghupload,Linux-aarch64)
	$(call ghupload,Linux-x86_64)
	$(call ghupload,Darwin-x86_64)
	$(call ghupload,Windows-x86_64)

all: build

release:
	git stash
	git fetch -p
	git checkout master
	git pull -r
	git tag v$(VERSION)
	git push origin v$(VERSION)
	git pull -r
