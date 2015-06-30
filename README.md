GoUrDHT
========

This project is part of the URDHT project to simplify and generalize DHTs.
The ultimate goal is to create a solid, accessible, and powerful backend for p2p applications.
Using the Client libraries, ideally developers should be able to use the Decentralized P2P DHT as a database and messaging layer for building their own applications.

Right now GoUrDHT is still in early development, and should be treated as such (expect bugs! Please report them in our issue tracker!). If you are interested in running a node to support the alpha testnet, read on:

## Installation

Ideally, GoUrDHT depends only on an install of Go 1.4+.
Libraries this project depends on have been incorporated into this repository.

- Clone this repository
- Edit config.json to match your network configuration
- run go run main.go to start the server (use 'screen' to run in the background right now)

## Configuring config.json

config.json looks like this:

```
{
	"bindAddr":"0.0.0.0", //set this to the interface you want the server to bind to
	"bindPort":8000, //set this to the port you want the rest-server to run on
	"wsBindAddr":"0.0.0.0", //this will match bindAddr in most cases
	"wsBindPort":8001, //this is the port you want the websocket-server to run on
	"bootstraps":"bootstrap.json", //this is a list of servers to bootstrap from
	"publicAddr":"http://127.0.0.1:8000/", //this should match the publicly accessible IP for your computer and bindPort
	"wsAddr":"ws://127.0.0.1:8001" // similar to above, matching wsBindPort
}
```

the "bindAddr" and "bindPort" indicate the network interface and port you want to bind to.
"publicAddr" is a url where your server can be reached, it should use your public ip and port utilized by NAT or port forwarding. You can use DNS names if they are configured to reach your node.
