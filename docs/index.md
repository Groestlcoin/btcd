# grsd

grsd is an alternative full node groestlcoin implementation written in Go
(golang).

One key difference between grsd and Groestlcoin Core is that grsd does *NOT*
include wallet functionality and this was a very intentional design decision.
This means you can't actually make or receive payments directly with grsd.
That functionality is provided by the
[grswallet](https://github.com/Groestlcoin/grswallet).

grsd is based on [btcd](https://github.com/btcsuite/btcd) and stays
very similar to it.

## Contents

* [Configuration](configuration.md)
* [Configuring TOR](configuring_tor.md)
* [Docker](using_docker.md)
* [Controlling](controlling.md)
* [Mining](mining.md)
* [Wallet](wallet.md)
* [JSON RPC API](json_rpc_api.md)

## License

grsd is licensed under the [copyfree](http://copyfree.org) ISC License.

