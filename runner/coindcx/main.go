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
					if candleMean.Max < 8 {
						return
					}

					fmt.Println(candleMean.VariencePer)

					//tickerCalc := model.GetTickerCalc(t)
					//fmt.Println(tickerCalc.NowHighPer)
					// lastPrice24, err := strconv.ParseFloat(t.LastPrice, 32)
					// if err != nil {
					// 	log.Println(err)
					// }
					//fmt.Println(candleMean.Candles10[0].Low - lastPrice24)
					//fmt.Println(t.LastPrice)

					// low24, err := strconv.ParseFloat(t.Low, 32)
					// if err != nil {
					// 	log.Println(err)
					// }
					// if candleMean.Candles10[0].Low == low24 {
					// 	fmt.Println(low24)
					// }
					tickerCalc := model.GetTickerCalc(t)
					fmt.Println(tickerCalc.LowHighPer)
				}(m.Pair, t)
			}
		}
	}
	wg.Wait()
}
