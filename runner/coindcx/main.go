package main

import (
	"fmt"
	"helper/api/coindcx"
	"helper/model"
	"log"
	"strings"
	"sync"

	"github.com/faizalom/go-web/controllers/marketcontroller"
)

var BaseCoin string
var wg sync.WaitGroup

func main() {
	BaseCoin = "USDT"
	marketsDetails, _ := coindcx.GetMarketsDetails()

	ticker, err := coindcx.GetExchange()
	if err != nil {
		log.Println(err)
	}

	for _, t := range ticker {
		if !strings.Contains(t.Market, BaseCoin) {
			continue
		}
		for _, m := range marketsDetails {
			if m.CoindcxName == t.Market {
				wg.Add(1)
				go func(pair string, t coindcx.Ticker) {
					defer wg.Done()
					candleMean := marketcontroller.CandleMean{}
					marketcontroller.GetCandles(pair, &candleMean)
					// if candleMean.Max < 5 {
					// 	fmt.Printf("%#v\n", candleMean.Max)
					// }
					tickerCalc := model.GetTickerCalc(t)
					fmt.Println(tickerCalc.LowHighPer)
				}(m.Pair, t)
			}
		}
	}
	wg.Wait()
}
