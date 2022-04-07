package main

import (
	"fmt"
	"helper/api/coindcx"
	"helper/model"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/faizalom/go-web/controllers/marketcontroller"
)

var BaseCoin string
var wg sync.WaitGroup
var LastPrice5 map[string][]float64

func SetLastPrice5() {
	LastPrice5 = make(map[string][]float64)
	ticker, err := coindcx.GetExchange()
	if err != nil {
		log.Println(err)
	}
	for _, t := range ticker {
		if !strings.Contains(t.Market, BaseCoin) {
			continue
		}
		LastPrice5[t.Market] = make([]float64, 5)
	}
}

func getRank(xc []float64) int {
	a := xc[4]
	b := xc[3]
	c := xc[2]
	d := xc[1]
	e := xc[0]

	rank := 0
	if a >= b {
		rank += 1
	}
	if b >= c {
		rank += 1
	}
	if c >= d {
		rank += 1
	}
	if d >= e {
		rank += 1
	}
	if a > e {
		rank += 1
	}
	return rank
}

func main() {
	BaseCoin = "USDT"
	SetLastPrice5()

	for range time.Tick(time.Second * 30) {

		ticker, err := coindcx.GetExchange()
		if err != nil {
			log.Println(err)
		}
		for _, t := range ticker {
			if !strings.Contains(t.Market, BaseCoin) {
				continue
			}
			lastPrice, _ := strconv.ParseFloat(t.LastPrice, 64)
			LastPrice5[t.Market] = append(LastPrice5[t.Market][1:], lastPrice)
			rank := getRank(LastPrice5[t.Market])
			if rank == 5 {
				fmt.Println(t.Market, LastPrice5[t.Market][0])
			}
		}
		fmt.Println("-----------------------------")
	}
}

func main3() {
	s := make([]int, 5)
	fmt.Println(s[1:])

	s = append(s[1:], 4)
	fmt.Println(s)

	s = append(s[1:], 4)
	fmt.Println(s)

	s = append(s[1:], 4)
	fmt.Println(s)

	x := make(map[string][]int)
	x["XX"] = make([]int, 5)

	x["XX"] = append(x["XX"][1:], 4)
	fmt.Println(x)

	x["XX"] = append(x["XX"][1:], 4)
	fmt.Println(x)

	x["XX"] = append(x["XX"][1:], 4)
	fmt.Println(x)

	// fmt.Println(len(s), s)
	// for len(s) > 0 {
	// 	x, s = s[0], s[1:] // undefined: x
	// 	fmt.Println(x)     // undefined: x
	// }
	// fmt.Println(len(s), s)
}

func main2() {
	fmt.Println("Works")
	a := 5
	b := 4
	c := 3
	d := 2
	e := 1

	rank := 0
	if a >= b {
		rank += 1
	}
	if b >= c {
		rank += 1
	}
	if c >= d {
		rank += 1
	}
	if d >= e {
		rank += 1
	}
	if a > e {
		rank += 1
	}

	fmt.Println(rank)
}

func main1() {
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

					WatchLow(t)

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

func WatchLow(t coindcx.Ticker) {
	tickerCalc := model.GetTickerCalc(t)
	fmt.Println(tickerCalc.LowHighPer)
}
