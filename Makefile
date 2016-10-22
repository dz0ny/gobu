VERSION := 0.7.2

all: setup build lint

setup:
	go get github.com/aktau/github-release
	go get github.com/alecthomas/gometalinter
	go get -v -d gobu/cmd/gobu
	bin/gometalinter --install --update

clean:
	rm -f gobu
	rm -rf pkg
	rm -rf bin
	find src/* -maxdepth 0 ! -name 'gobu' -type d | xargs rm -rf

lint:
	bin/gometalinter --fast --disable=gotype --disable=gas --cyclo-over=15 src/gobu/...
	find src/gobu -name '*.go' | xargs gofmt -w -s

build:
	env GOOS=linux GOARCH=arm go build --ldflags '-w -X main.build=$(VERSION)' -o gobu-Linux-armv7l gobu/cmd/gobu
	env GOOS=linux GOARCH=amd64 go build --ldflags '-s -w -X main.build=$(VERSION)' -o gobu-Linux-x86_64 gobu/cmd/gobu
	env GOOS=darwin GOARCH=amd64 go build --ldflags '-w -X main.build=$(VERSION)' -o gobu-Darwin-x86_64 gobu/cmd/gobu
	env GOOS=windows GOARCH=amd64 go build --ldflags '-w -X main.build=$(VERSION)' -o gobu-Windows-x86_64.exe gobu/cmd/gobu

install:
	sudo mv gobu-`uname -s`-`uname -m` /usr/local/bin/gobu

upload:
	bin/github-release upload \
			--user dz0ny \
			--repo gobu \
			--tag "v$(VERSION)" \
			--name "gobu-Linux-armv6l" \
			--file gobu-Linux-armv7l
	bin/github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag "v$(VERSION)" \
	    --name "gobu-Linux-armv7l" \
	    --file gobu-Linux-armv7l
	bin/github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag "v$(VERSION)" \
	    --name "gobu-Linux-x86_64" \
	    --file gobu-Linux-x86_64
	bin/github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag "v$(VERSION)" \
	    --name "gobu-Darwin-x86_64" \
	    --file gobu-Darwin-x86_64
	bin/github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag "v$(VERSION)" \
	    --name "gobu-Windows-x86_64.exe" \
	    --file gobu-Windows-x86_64.exe
