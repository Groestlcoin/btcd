grsd
====

grsd is an alternative full node groestlcoin implementation written in Go
(golang).

One key difference between grsd and Groestlcoin Core is that grsd does *NOT*
include wallet functionality and this was a very intentional design decision.
This means you can't actually make or receive payments directly with grsd.
That functionality is provided by the
[grswallet](https://github.com/Groestlcoin/grswallet).

grsd is based on [btcd](https://github.com/btcsuite/btcd) and stays
very similar to it.

## Requirements

[Go](http://golang.org) 1.12 or newer.

## Installation

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain grsd, all dependencies, and install it:

```bash
$ git clone https://github.com/Groestlcoin/grsd
$ cd grsd
$ make install
```

NOTE: Do not use `go install` command, because it will install the binaries under wrong names.

- grsd (and utilities) will now be installed in `$GOPATH/bin`.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

- If you do not have GNU make in Windows just type commands from `build` section of Makefile:

```
C:\grsd> go build -o grsd .
C:\grsd> go build -o grsctl ./cmd/btcctl
```

`grsd.exe` and `grsctl.exe` can be copied to any directory you like.

## Updating

- Run the following commands to update grsd, all dependencies, and install it:

```bash
$ cd grsd
$ git pull
$ make install
```

## Getting Started

grsd has several configuration options available to tweak how it runs, but all
of the basic operations described in the intro section work with zero
configuration.

#### Linux/BSD/POSIX/Source

```bash
$ grsd
```

grsctl (analogue of groestlcoin-cli) can be used to control grsd:
```bash
$ grsctl getinfo
```

## Documentation

The documentation is a work-in-progress.  It is located in the
[docs](https://github.com/Groestlcoin/grsd/tree/grssuite/docs) folder.

## License

grsd is licensed under the [copyfree](http://copyfree.org) ISC License.
