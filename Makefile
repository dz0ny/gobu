all: prepare build upload

prepare:
	go get github.com/aktau/github-release
	go get gobu

build:
	env GOOS=linux GOARCH=amd64 go build -o gobu-linux-amd64 gobu
	env GOOS=darwin GOARCH=amd64 go build -o gobu-darwin-amd64 gobu
	env GOOS=windows GOARCH=amd64 go build -o gobu-windows-amd64.exe gobu

install: build
	sudo mv gobu-linux-amd64 /usr/sbin/gobu

upload:
	bin/github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag v0.1.6 \
	    --name "gobu-linux-amd64" \
	    --file gobu-linux-amd64
	bin/github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag v0.1.6 \
	    --name "gobu-darwin-amd64" \
	    --file gobu-darwin-amd64
	bin/github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag v0.1.6 \
	    --name "gobu-windows-amd64.exe" \
	    --file gobu-windows-amd64.exe
