While grsd is highly configurable when it comes to the network configuration,
the following is intended to be a quick reference for the default ports used so
port forwarding can be configured as required.

grsd provides a `--upnp` flag which can be used to automatically map the groestlcoin
peer-to-peer listening port if your router supports UPnP.  If your router does
not support UPnP, or you don't wish to use it, please note that only the groestlcoin
peer-to-peer port should be forwarded unless you specifically want to allow RPC
access to your grsd from external sources such as in more advanced network
configurations.

| - | Bitcoin mainnet | Groestlcoin mainnet | Bitcoin testnet | Groestlcoin testnet
 ---------------------- | ---- | ---- | ----- | -----
**Wallet/Original RPC** | 8332 | 1441 | 18332 | 17766
**P2P RPC**             | 8333 | 1331 | 18333 | 17777
**btcd/grsd RPC**       | 8334 | 1444 | 18334 | 17764
