
package main

import (
          "math"
          "bytes"
          "net/http"
          "net/rpc"
          "encoding/json"
          "strings"
          "fmt"
          "io/ioutil"
          "strconv"
        )

type Js struct {
List f_list `json:"list"`
}

type f_list  struct {
Meta f_meta `json:"-"`
Resources []f_res `json:"resources"`
}

type f_meta struct {
Type string `json:"-"`
Start int32 `json:"-"`
Count  int32 `json:"-"`
}

type f_res struct {
Resource j_res `json:"resource"`
}

type j_res struct {
Classname string `json:"classname"`
Fields ff `json:"fields"`
}

type ff struct{
Price string `json:"price"`
Symbol string `json:"symbol"`

}

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


type Services int

var M map[int]My_Quote


   func (t *Services) Get_Price(args *Arguments, quote *My_Quote) error {

     a := string(args.StockSymbolAndPercentage[:])
	  a = strings.Replace(a,":",",",-1)
	  a = strings.Replace(a,"%",",",-1)
	  a = strings.Replace(a,",,",",",-1)
	  a = strings.Trim(a," ")
	  a = strings.Replace(a,"\"","",-1)
	  a = strings.TrimSpace(a)
	  a = strings.TrimSuffix(a,",")
 /*a := args.StockSymbolAndPercentage[:]
  a = strings.Replace(a,":","",-1)
  a = strings.Replace(a,",","",-1)
  a = strings.Replace(a,"%",",",-1)
  a = strings.TrimSuffix(a, ",")
  fmt.Println("string b is ",a)
*/    Stockstmp:= strings.Split(a,",")
    fmt.Println("string after split is ",Stockstmp)

	  var ReqUrl string

	for  i :=0; i < len(Stockstmp) ; i++ {
	                       i = i+1

	temp,_ := strconv.ParseFloat(Stockstmp[i],64)
  fmt.Println("Extracted temp is ",temp)
	 temp= (temp * args.Budget * 0.01)
  fmt.Println("temp is ",temp)
	ReqUrl= ReqUrl + (Stockstmp[i-1] + ",")


                                }


		ReqUrl = strings.TrimSuffix(ReqUrl,",")
fmt.Println("Requested URL is ",ReqUrl)
  UrlStr := "http://finance.yahoo.com/webservice/v1/symbols/" + ReqUrl +  "/quote?format=json"


  client := &http.Client{}
  resp, _  := client.Get(UrlStr)
  req, _ := http.NewRequest("GET", UrlStr, nil)

  req.Header.Add("If-None-Match", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")


// make request
resp, _  = client.Do(req)
if( resp.StatusCode >= 200 && resp.StatusCode < 300 ) {
          var C Js
        body, _ := ioutil.ReadAll(resp.Body)

 err := json.Unmarshal(body, &C)

 n:= len(Stockstmp)

Quo:= make([]float64,n,n)


 for  i :=0; i < n ; i++ {
	         i = i + 1

  TempFloat,_ := strconv.ParseFloat(Stockstmp[i],64)
	Quo[i] = (TempFloat * args.Budget * 0.01)

	                         }


	var buffer bytes.Buffer
	q:=0
  for _ ,Sample := range  C.List.Resources {


        temp1:= Sample.Resource.Fields.Symbol
        temp2,_:=strconv.ParseFloat(Sample.Resource.Fields.Price,64)
        temp3:= (int)(Quo[q+1]/temp2)
        temp4:= math.Mod(Quo[q+1],temp2)
        q = q + 2

        quote.Stocks = fmt.Sprintf("%s:%d:%g",temp1,temp3,temp2)
        quote.Investedamount = quote.Investedamount + temp4
        buffer.WriteString(quote.Stocks)
        buffer.WriteString(",")
                                            }


              quote.Tradeid = quote.Tradeid + 1
              quote.Stocks = (buffer.String())
              quote.Stocks = strings.TrimSuffix(quote.Stocks, ",")


              M = map[int]My_Quote {
                      quote.Tradeid : {quote.Stocks,quote.Investedamount,quote.Tradeid},
                                  }



        if( err == nil ) { fmt.Println("Completed") }
  }       else          {  fmt.Println(resp.Status); }
          return nil

 }

  func main() {
generatequote := new(Services)
rpc.Register(generatequote)
rpc.HandleHTTP()

err := http.ListenAndServe(":1550", nil)
if err != nil {
fmt.Println(err.Error())
}
}
