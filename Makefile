build:
	GO111MODULE=on go build -o grsd .
	GO111MODULE=on go build -o grsctl ./cmd/btcctl

build-all: build
	GO111MODULE=on go build ./cmd/addblock
	GO111MODULE=on go build ./cmd/findcheckpoint
	GO111MODULE=on go build ./cmd/gencerts

install: build
	cp grsd grsctl `go env GOPATH`/bin

reset-mod:
	git checkout go.mod go.sum

test:
	GO111MODULE=on go test -tags="rpctest" ./...

clean:
	rm -f grsd grsctl btcd addblock findcheckpoint gencerts
