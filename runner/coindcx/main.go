package main

import (
	"fmt"
	"helper/api/coindcx"
	"log"
	"strconv"
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
					marketcontroller.GetCandles(pair, &candleMean, "5")
					if candleMean.Min < 4 || candleMean.VariencePer > 6 {
						return
					}

					LastPrice, _ := strconv.ParseFloat(t.LastPrice, 64)

					if (candleMean.Mean + candleMean.Variance) > LastPrice {
						fmt.Println(t.Market, candleMean.VariencePer, candleMean.Min, candleMean.Max, candleMean.Mean, candleMean.Mean+candleMean.Variance, t.LastPrice)
					}
				}(m.Pair, t)
			}
		}
	}
	wg.Wait()
}
