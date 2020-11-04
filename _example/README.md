# Example Application

This is a demo application using `gateapi` to show how Gate APIv4 works. 
Instead of running it, it is recommended to read the source code to get a general idea of
how this SDK is used. However, you can modify this code directly to implement your own logic.

## Build

```bash
# change into your $GOPATH
cd $GOPATH

# create a new application
mkdir gateapi-demo && cd gateapi-demo

# install required dependency
go mod init
go get github.com/gateio/gateapi-go/v5
go get github.com/shopspring/decimal

# build the demo application
go build
```

## Run

**READ THIS BEFORE YOU RUN ANYTHING**

**This application is shown for demo only. It will try to use your input API key and secret to
trade, lend and borrow, etc. Make sure you know exactly what it does before running it.**

```bash
# run futures demo against TestNet
./gateapi-demo futures -k <YOUR_TESTNET_API_KEY> -s <YOUR_TESTNET_API_SECRET> -u fx-api-testnet.gateio.ws

# run futures demo against real trading
./gateapi-demo futures -k <YOUR_API_KEY> -s <YOUR_API_SECRET>

# run spot demo
./gateapi-demo spot -k <YOUR_API_KEY> -s <YOUR_API_SECRET>

# run margin demo
./gateapi-demo margin -k <YOUR_API_KEY> -s <YOUR_API_SECRET>
```
