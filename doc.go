// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
grsd is a full-node groestlcoin implementation written in Go.

The default options are sane for most users.  This means grsd will work 'out of
the box' for most users.  However, there are also a wide variety of flags that
can be used to control it.

The following section provides a usage overview which enumerates the flags.  An
interesting point to note is that the long form of all of these options
(except -C) can be specified in a configuration file that is automatically
parsed when grsd starts up.  By default, the configuration file is located at
~/.grsd/grsd.conf on POSIX-style operating systems and %LOCALAPPDATA%\grsd\grsd.conf
on Windows.  The -C (--configfile) flag, as shown below, can be used to override
this location.

Usage:
  grsd [OPTIONS]

Application Options:
  -V, --version               Display version information and exit
  -C, --configfile=           Path to configuration file
                              (/home/yura/.grsd/grsd.conf)
  -b, --datadir=              Directory to store data (/home/yura/.grsd/data)
      --logdir=               Directory to log output. (/home/yura/.grsd/logs)
  -a, --addpeer=              Add a peer to connect with at startup
      --connect=              Connect only to the specified peers at startup
      --nolisten              Disable listening for incoming connections --
                              NOTE: Listening is automatically disabled if the
                              --connect or --proxy options are used without
                              also specifying listen interfaces via --listen
      --listen=               Add an interface/port to listen for connections
                              (default all interfaces port: 1331, testnet:
                              17777)
      --maxpeers=             Max number of inbound and outbound peers (125)
      --nobanning             Disable banning of misbehaving peers
      --banduration=          How long to ban misbehaving peers.  Valid time
                              units are {s, m, h}.  Minimum 1 second (24h0m0s)
      --banthreshold=         Maximum allowed ban score before disconnecting
                              and banning misbehaving peers. (100)
      --whitelist=            Add an IP network or IP that will not be banned.
                              (eg. 192.168.1.0/24 or ::1)
  -u, --rpcuser=              Username for RPC connections (yourrpcuser)
  -P, --rpcpass=              Password for RPC connections
      --rpclimituser=         Username for limited RPC connections
      --rpclimitpass=         Password for limited RPC connections
      --rpclisten=            Add an interface/port to listen for RPC
                              connections (default port: 1444, testnet: 17764)
      --rpccert=              File containing the certificate file
                              (/home/yura/.grsd/rpc.cert)
      --rpckey=               File containing the certificate key
                              (/home/yura/.grsd/rpc.key)
      --rpcmaxclients=        Max number of RPC clients for standard
                              connections (10)
      --rpcmaxwebsockets=     Max number of RPC websocket connections (25)
      --rpcmaxconcurrentreqs= Max number of concurrent RPC requests that may be
                              processed concurrently (20)
      --rpcquirks             Mirror some JSON-RPC quirks of Groestlcoin Core
                              -- NOTE: Discouraged unless interoperability
                              issues need to be worked around
      --norpc                 Disable built-in RPC server -- NOTE: The RPC
                              server is disabled by default if no
                              rpcuser/rpcpass or rpclimituser/rpclimitpass is
                              specified
      --notls                 Disable TLS for the RPC server -- NOTE: This is
                              only allowed if the RPC server is bound to
                              localhost
      --nodnsseed             Disable DNS seeding for peers
      --externalip=           Add an ip to the list of local addresses we claim
                              to listen on to peers
      --proxy=                Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)
      --proxyuser=            Username for proxy server
      --proxypass=            Password for proxy server
      --onion=                Connect to tor hidden services via SOCKS5 proxy
                              (eg. 127.0.0.1:9050)
      --onionuser=            Username for onion proxy server
      --onionpass=            Password for onion proxy server
      --noonion               Disable connecting to tor hidden services
      --torisolation          Enable Tor stream isolation by randomizing user
                              credentials for each connection.
      --testnet               Use the test network
      --regtest               Use the regression test network
      --simnet                Use the simulation test network
      --addcheckpoint=        Add a custom checkpoint.  Format:
                              '<height>:<hash>'
      --nocheckpoints         Disable built-in checkpoints.  Don't do this
                              unless you know what you're doing.
      --dbtype=               Database backend to use for the Block Chain
                              (ffldb)
      --profile=              Enable HTTP profiling on given port -- NOTE port
                              must be between 1024 and 65536
      --cpuprofile=           Write CPU profile to the specified file
  -d, --debuglevel=           Logging level for all subsystems {trace, debug,
                              info, warn, error, critical} -- You may also
                              specify
                              <subsystem>=<level>,<subsystem2>=<level>,... to
                              set the log level for individual subsystems --
                              Use show to list available subsystems (info)
      --upnp                  Use UPnP to map our listening port outside of NAT
      --minrelaytxfee=        The minimum transaction fee in GRS/kB to be
                              considered a non-zero fee. (1e-05)
      --limitfreerelay=       Limit relay of transactions with no transaction
                              fee to the given amount in thousands of bytes per
                              minute (15)
      --norelaypriority       Do not require free or low-fee transactions to
                              have high priority for relaying
      --trickleinterval=      Minimum time between attempts to send new
                              inventory to a connected peer (10s)
      --maxorphantx=          Max number of orphan transactions to keep in
                              memory (100)
      --generate              Generate (mine) groestlcoins using the CPU
      --miningaddr=           Add the specified payment address to the list of
                              addresses to use for generated blocks -- At least
                              one address is required if the generate option is
                              set
      --blockminsize=         Mininum block size in bytes to be used when
                              creating a block
      --blockmaxsize=         Maximum block size in bytes to be used when
                              creating a block (750000)
      --blockminweight=       Mininum block weight to be used when creating a
                              block
      --blockmaxweight=       Maximum block weight to be used when creating a
                              block (3000000)
      --blockprioritysize=    Size in bytes for high-priority/low-fee
                              transactions when creating a block (50000)
      --uacomment=            Comment to add to the user agent -- See BIP 14
                              for more information.
      --nopeerbloomfilters    Disable bloom filtering support
      --nocfilters            Disable committed filtering (CF) support
      --dropcfindex           Deletes the index used for committed filtering
                              (CF) support from the database on start up and
                              then exits.
      --sigcachemaxsize=      The maximum number of entries in the signature
                              verification cache (100000)
      --blocksonly            Do not accept transactions from remote peers.
      --txindex               Maintain a full hash-based transaction index
                              which makes all transactions available via the
                              getrawtransaction RPC
      --droptxindex           Deletes the hash-based transaction index from the
                              database on start up and then exits.
      --addrindex             Maintain a full address-based transaction index
                              which makes the searchrawtransactions RPC
                              available
      --dropaddrindex         Deletes the address-based transaction index from
                              the database on start up and then exits.
      --relaynonstd           Relay non-standard transactions regardless of the
                              default settings for the active network.
      --rejectnonstd          Reject non-standard transactions regardless of
                              the default settings for the active network.

Help Options:
  -h, --help                  Show this help message

*/
package main
