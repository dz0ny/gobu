prepare:
	go get github.com/aktau/github-release

build:
	env GOOS=linux GOARCH=amd64 go build -o gobu-linux-amd64
	env GOOS=darwin GOARCH=amd64 go build -o gobu-darwin-amd64
	env GOOS=windows GOARCH=amd64 go build -o gobu-windows-amd64.exe

upload:
	github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag v0.1.0 \
	    --name "gobu-linux-amd64" \
	    --file gobu-linux-amd64
	github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag v0.1.0 \
	    --name "gobu-darwin-amd64" \
	    --file gobu-darwin-amd64
	github-release upload \
	    --user dz0ny \
	    --repo gobu \
	    --tag v0.1.0 \
	    --name "gobu-windows-amd64.exe" \
	    --file gobu-windows-amd64.exe
