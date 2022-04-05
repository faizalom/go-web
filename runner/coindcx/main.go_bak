package main

import (
	"fmt"
	"helper/api/coindcx"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/faizalom/go-web/controllers/marketcontroller"
)

// var fm = template.FuncMap{
// 	"strReplace":  tplfunc.StrReplace,
// 	"strContains": tplfunc.StrContains,
// 	"strToLower":  tplfunc.StrToLower,
// 	"add":         tplfunc.Add,
// 	"dateFormat":  tplfunc.DateFormat,
// 	"dFormat":     tplfunc.DFormat,
// }

var BaseCoin string
var WatchList map[string]bool
var BuyCoin map[string]bool
var BuyCoinBuyPrice map[string]float64
var BuyCoinSellPrice map[string]float64

var wg sync.WaitGroup
var Candles []map[string]interface{}

func main() {
	BaseCoin = "USDT"
	ticker, err := coindcx.GetExchange()
	if err != nil {
		log.Println(err)
	}
	//MaxLeverageShort interface{}
	var markets []struct {
		coindcx.Ticker
		Coin         string
		LowNowMargin float64
		Timestamp    time.Time
		LowHighPer   float64
		marketcontroller.CandleMean
	}

	marketsDetails, _ := coindcx.GetMarketsDetails()

	for _, t := range ticker {
		if !strings.Contains(t.Market, BaseCoin) {
			continue
		}

		lastPrice, _ := strconv.ParseFloat(t.LastPrice, 64)
		low, _ := strconv.ParseFloat(t.Low, 64)
		high, _ := strconv.ParseFloat(t.High, 64)
		timeStamp, _ := t.Timestamp.Int64()
		timeUnix := time.Unix(timeStamp, 0)
		lowHighPer := (high/low - 1) * 100
		market := struct {
			coindcx.Ticker
			Coin         string
			LowNowMargin float64
			Timestamp    time.Time
			LowHighPer   float64
			marketcontroller.CandleMean
		}{
			Ticker:       t,
			Coin:         strings.Replace(t.Market, BaseCoin, "", -1),
			LowNowMargin: ((lastPrice - low) / low * 100),
			Timestamp:    timeUnix,
			LowHighPer:   lowHighPer,
		}

		for _, m := range marketsDetails {
			if m.CoindcxName == t.Market {
				wg.Add(1)
				go func(pair string) {
					defer wg.Done()
					marketcontroller.GetCandles(pair, &market.CandleMean)
					markets = append(markets, market)

					// if market.Min > 5 {
					// 	fmt.Println(pair, market.Min)
					// }

					//fmt.Println(market.CandleMean.Candles10[0].Open)

					//if pair == "B-CHR_USDT" {
					//	fmt.Printf("%#v", &market)
					//}

				}(m.Pair)
			}
		}

	}
	wg.Wait()

	for _, m := range markets {
		fmt.Println(m.Low, "\t", m.CandleMean.Candles10[0].Low)
	}

	fmt.Println("Done")
}
