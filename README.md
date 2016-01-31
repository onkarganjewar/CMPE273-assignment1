# CMPE273 – Assignment 1
# Virtual stock trading system in Golang
System allows user to simulate the trading of stocks using real-time pricing via Yahoo Finance API. The system supports USD currency only. 
## Features   
There are 2 components in the system: Client and Server. JSON-RPC Client will take command line input and sends request to the server. Server will have JSON-RPC interface for providing various features provided in the system such as Buying or Selling Stocks & Checking the portfolio.    
•	Buying Stocks   
•	Checking portfolio    
## Requirements  
•	Golang latest stable version (I have used go1.5)
•	You can check the official golang releases here: https://golang.org/doc/devel/release.html
## Installation
### Installing Go (if you have not already installed it)
•	There are various ways to install go as per the operating system that you’re working on.   
•	All the required files and step-by-step instructions can be found here : https://golang.org/doc/install    
### Installing Package
After you have installed the Golang then run the following command      
go get github.com/onkarganjewar/CMPE273-assignment1

## Usage
Start the server
go run server.go

Run the client with localhost
go run client.go 127.0.0.1

After successful deployment you will be getting this on client side
Enter the stock symbol and percentage of share to buy


Now, enter the details about the stocks in the following format; suppose you want to buy 50% of your budget to buy stocks of Google (GOOG) and remaining 50% to buy stocks of Yahoo (YHOO) then enter 
“GOOG:50%,YHOO:50%”

The console now looks something like this  
![img](https://cloud.githubusercontent.com/assets/14006620/12701321/d533e796-c7ba-11e5-92d2-a9a84db289da.png)


