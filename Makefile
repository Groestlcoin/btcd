build:
	GO111MODULE=on go build -o grsd .
	GO111MODULE=on go build -o grsctl ./cmd/btcctl

install: build
	cp grsd grsctl `go env GOPATH`/bin

reset-mod:
	git checkout go.mod go.sum

clean:
	rm -f grsd grsctl btcd
