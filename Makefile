GOVERSION := 1.6.2
VERSION := 0.6.2-1

all: build

clean:
	rm -f gobu
	rm -rf pkg
	rm -rf bin
	find src/* -maxdepth 0 ! -name 'gobu' -type d | xargs rm -rf

build:
	go get -v -d gobu
	env GOOS=linux GOARCH=arm go build --ldflags '-w -X main.globalVersion=$(GOVERSION)' -o gobu-Linux-armv7l gobu
	env GOOS=linux GOARCH=amd64 go build --ldflags '-w -X main.globalVersion=$(GOVERSION)' -o gobu-Linux-x86_64 gobu
	env GOOS=darwin GOARCH=amd64 go build --ldflags '-w -X main.globalVersion=$(GOVERSION)' -o gobu-Darwin-x86_64 gobu
	env GOOS=windows GOARCH=amd64 go build --ldflags '-w -X main.globalVersion=$(GOVERSION)' -o gobu-Windows-x86_64.exe gobu

install:
	sudo mv gobu-linux-amd64 /usr/sbin/gobu

upload:
	go get github.com/aktau/github-release
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
