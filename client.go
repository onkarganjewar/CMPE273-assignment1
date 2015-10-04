package main

import (
  "os"
  "log"
  "fmt"
  "net/rpc"
  "strconv"
  )


type Arguments struct{
StockSymbolAndPercentage string
Budget float64
}


type My_Quote struct {
Stocks string `json:"stocksymbol"`
Investedamount float64	`json:"stockprice"`
Tradeid int `json:"id"`
}

type Id struct{
Tradeid int `json:"id"`
}

type Latest_Quote struct {
Stocks string `json:"stocksymbol"`
Investedamount float64	`json:"stockprice"`
}

var a Arguments

  // Create Client
  func main() {

    if len(os.Args) != 2 {
          fmt.Println("Usage: ", os.Args[0], "server")
          os.Exit(1)
        }

      serverAddress := os.Args[1]

      client, err := rpc.DialHTTP("tcp", serverAddress+":1550")

      if err != nil {
          log.Fatal("dialing:", err)
          }




  fmt.Println("enter the stock symbol and percentage of share to buy");
  fmt.Scanln(&a.StockSymbolAndPercentage);
  fmt.Println("Enter the budget");
  fmt.Scanln(&a.Budget)
  fmt.Println("Entered details are")
  fmt.Println(a)

  var ge My_Quote

  err = client.Call("Services.Get_Price",a,&ge)

  if err != nil {
                  log.Fatal("service error", err)
                }

  fmt.Println("Stocks : ",ge.Stocks)
  fmt.Println("Trade ID : ",ge.Tradeid)
  fmt.Println("UnInvested amount : ",ge.Investedamount)

  f, _ := strconv.ParseFloat("ge.Stocks", 64)


// For generating the portfolio


fmt.Println("\nEnter Trade Id")
var s int
fmt.Scan(&s)
if s == 1 {


  var ga My_Quote

  err = client.Call("Services.Get_Price",a,&ga)

  if err != nil {
                    log.Fatal("service error", err)
                }

  x, _ := strconv.ParseFloat("ga.Stocks", 64)
  y := x - f
  fmt.Println("Stocks:", ga.Stocks)
  fmt.Println("Profit/Loss per stock:", y)
  fmt.Println("Uninvested amount:", ga.Investedamount)
}

}
