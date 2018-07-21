# franklin
Real-time stock data and more from your command line

## Function
1. Fetches current price of Equity Security (Stocks)
2. Fetches real-time data about an Equity Security

## Usage
1. Price only
    ```$ franklin --ticker=aapl```
    ```$ franklin -t=aapl```

2. Selected details
    ```$ franklin --ticker=aapl --details```
    ```$ franklin -t=aapl -d```

## Setup
- ensure working go environment
- clone repository
- cd franklin
- go install
- run ```$ franklin``` to your hearts content

## Setting GO Path
```$ export PATH=$PATH:/usr/local/go/bin```
```$ export GOPATH=$HOME/go;```
```$ export PATH=$PATH:$GOPATH/bin;```
