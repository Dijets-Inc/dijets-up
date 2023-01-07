# Dijets-up

## Note

This tool is under heavy development and the documentation/code snippets below may vary slightly from the actual code in the repository. Updates to the documentation may happen some time after an update to the codebase. Nonetheless, this README should provide valuable information about using this tool.

## Overview

This is a tool to run and interact with a local Dijets network.
This tool may be especially useful for development and testing.

## Installation

To download a binary for the latest release, run:

```sh
curl -sSfL https://raw.githubusercontent.com/lasthyphen/dijets-up/main/scripts/install.sh | sh -s
```

The binary will be installed inside the `~/bin` directory.

To add the binary to your path, run

```sh
export PATH=~/bin:$PATH
```

To add it to your path permanently, add an export command to your shell initialization script (ex: .bashrc).

## Build from source code

This is only needed by advanced users who want to modify or test Dijets-up in specific ways.

Requires golang to be installed on the system ([https://go.dev/doc/install](https://go.dev/doc/install)).

### Download

```sh
git clone https://github.com/lasthyphen/dijets-up.git
```

### Build

From inside the cloned directory:

```sh
./scripts/build.sh
```

The binary will be installed inside the `./bin` directory.

To add the binary to your path, run

```sh
export PATH=$PWD/bin:$PATH
```

### Run Unit Tests

Inside the directory cloned above:

```sh
go test ./...
```

### Run E2E tests

The E2E test checks `dijets-up` RPC communication and control. It starts a network against a fresh RPC
server and executes a set of query and control operations on it.

To start it, execute inside the cloned directory:

```sh
./scripts/tests.e2e.sh
```

## Using `dijets-up`

You can import this repository as a library in your Go program, but we recommend running `dijets-up` as a binary. This creates an RPC server that you can send requests to in order to start a network, add nodes to the network, remove nodes from the network, restart nodes, etc.. You can make requests through the `dijets-up` command or by making API calls. Requests are "translated" into gRPC and sent to the server.

**Why does `dijets-up` need an RPC server?** `dijets-up` needs to provide complex workflows such as replacing nodes, restarting nodes, injecting fail points, etc.. The RPC server exposes basic operations to enable a separation of concerns such that one team develops a test framework, and the other writes test cases and controlling logic.

**Why gRPC?** The RPC server leads to more modular test components, and gRPC enables greater flexibility. The protocol buffer increases flexibility as we develop more complicated test cases. And gRPC opens up a variety of different approaches for how to write test controller (e.g., Rust). See [`rpcpb/rpc.proto`](./rpcpb/rpc.proto) for service definition.

**Why gRPC gateway?** [gRPC gateway](https://grpc-ecosystem.github.io/grpc-gateway/) exposes gRPC API via HTTP, without us writing any code. Which can be useful if a test controller writer does not want to deal with gRPC.

## `network-runner` RPC server: examples

To start the server:

```bash
dijets-up server \
--log-level debug \
--port=":8080" \
--grpc-gateway-port=":8081"

# set "--disable-grpc-gateway" to disable gRPC gateway
```

Note that the above command will run until you stop it with `CTRL + C`. You should run further commands in a separate terminal.

To ping the server:

```bash
curl -X POST -k http://localhost:8081/v1/ping -d ''

# or
dijets-up ping \
--log-level debug \
--endpoint="0.0.0.0:8080"
```

To start a new Dijets network with five nodes (a cluster):

```bash
# replace execPath with the path to a Dijets Node Binary Files
# e.g., ${HOME}/go/src/github.com/Dijets-Inc/dijetsnodego/build/dijetsnodego
DIJETSNODE_EXEC_PATH="dijetsnodego"

curl -X POST -k http://localhost:8081/v1/control/start -d '{"execPath":"'${DIJETSNODE_EXEC_PATH}'","numNodes":5,"logLevel":"DEBUG"}'

# or
dijets-up control start \
--log-level debug \
--endpoint="0.0.0.0:8080" \
--number-of-nodes=5 \
--dijetsnodego-path ${DIJETSNODE_EXEC_PATH}
```

Additional optional parameters which can be passed to the start command:

```bash
--plugin-dir ${AVALANCHEGO_PLUGIN_PATH} \
--blockchain-specs '[{"vm_name": "subnetevm", "genesis": "/tmp/subnet-evm.genesis.json"}]'
--global-node-config '{"index-enabled":false, "api-admin-enabled":true,"network-peer-list-gossip-frequency":"300ms"}'
--custom-node-configs" '{"node1":{"log-level":"debug","api-admin-enabled":false},"node2":{...},...}'
```

Configurations:

You can use Dijets-up with many different configurations and custom parameters. Network snapshots, ability to pair-up with other bash script and run a continuous process and simulating various test environments are just some of many features included in the latest release.