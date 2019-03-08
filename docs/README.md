### Table of Contents
1. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying grsd via grsctl](#grsctlConfig)
    4. [Mining](#Mining)
2. [Help](#Help)
    1. [Startup](#Startup)
        1. [Using bootstrap.dat](#BootstrapDat)
    2. [Network Configuration](#NetworkConfig)
    3. [Wallet](#Wallet)
3. [Developer Resources](#DeveloperResources)
    1. [JSON-RPC Reference](#JSONRPCReference)

<a name="GettingStarted" />

### 1. Getting Started

<a name="Installation" />

**1.1 Installation**

See [here](https://github.com/Groestlcoin/grsd/blob/grssuite/README.md#installation).

<a name="Configuration" />

**1.2 Configuration**

grsd has a number of [configuration](http://godoc.org/github.com/Groestlcoin/grsd)
options, which can be viewed by running: `$ grsd --help`.

<a name="grsctlConfig" />

**1.3 Controlling and Querying grsd via grsctl**

grsctl is a command line utility that can be used to both control and query grsd
via [RPC](http://www.wikipedia.org/wiki/Remote_procedure_call).  grsd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password or both an RPC limited username and password:

* grsd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
* grsctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
OR
```
[Application Options]
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
For a list of available options, run: `$ grsctl --help`

<a name="Mining" />

**1.4 Mining**

grsd supports the `getblocktemplate` RPC.
The limited user cannot access this RPC.


**1.4.1. Add the payment addresses with the `miningaddr` option.**

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX
miningaddr=1M83ju3EChKYyysmM2FXtLNftbacagd8FR
```

**1.4.2. Add grsd's RPC TLS certificate to system Certificate Authority list.**

`cgminer` uses [curl](http://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `grsd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**

1. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/user/.grsd/rpc.cert /usr/share/ca-certificates/grsd.crt`
2. Add grsd.crt to /etc/ca-certificates.conf: `# echo grsd.crt >> /etc/ca-certificates.conf`
3. Update the CA certificate list: `# update-ca-certificates`

**1.4.3. Set your mining software url to use https.**

`$ cgminer -o https://127.0.0.1:1444 -u rpcuser -p rpcpassword`

<a name="Help" />

### 2. Help

<a name="Startup" />

**2.1 Startup**

Typically grsd will run and start downloading the block chain with no extra
configuration necessary, however, there is an optional method to use a
`bootstrap.dat` file that may speed up the initial block chain download process.

<a name="BootstrapDat" />

**2.1.1 bootstrap.dat**

* [Using bootstrap.dat](https://github.com/Groestlcoin/grsd/tree/grssuite/docs/using_bootstrap_dat.md)

<a name="NetworkConfig" />

**2.2 Network Configuration**

* [What Ports Are Used by Default?](https://github.com/Groestlcoin/grsd/tree/grssuite/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/Groestlcoin/grsd/tree/grssuite/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/Groestlcoin/grsd/tree/grssuite/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring grsd with Tor](https://github.com/Groestlcoin/grsd/tree/grssuite/docs/configuring_tor.md)

<a name="Wallet" />

**2.3 Wallet**

grsd was intentionally developed without an integrated wallet for security
reasons.  Please see [grswallet](https://github.com/Groestlcoin/grswallet) for more
information.

<a name="DeveloperResources" />

### 3. Developer Resources

<a name="JSONRPCReference" />

* [JSON-RPC Reference](https://github.com/Groestlcoin/grsd/tree/grssuite/docs/json_rpc_api.md)

