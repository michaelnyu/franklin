# franklin
Real-time stock data and more from your command line

## Function
1. Fetches current price of Equity Security (Stocks)
2. Fetches real-time data about an Equity Security

## Usage
1. Price only
    ```shell
    $ franklin --ticker=aapl
    $ franklin -t=aapl
    ```

2. Selected details
    ```shell
    $ franklin --ticker=aapl --details
    $ franklin -t=aapl -d
    ```

## Setup
- ensure working go environment
- install with:
```
$ go get github.com/michaelnyu/franklin
```
- cd franklin
- go install
- run ```franklin``` to your hearts content

## Setting GO Path
```shell
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go;
export PATH=$PATH:$GOPATH/bin;
```
