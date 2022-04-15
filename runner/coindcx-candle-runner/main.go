package main

import (
	"encoding/gob"
	"helper/api/coindcx"
	"helper/api/telegram"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/faizalom/go-web/controllers/marketcontroller"
	"github.com/faizalom/go-web/lib"
)

var BaseCoin string

func getStoredCandle(market string) (marketcontroller.CandleMean, error) {
	fname := lib.TempCandPath + "/" + market
	candleMean := marketcontroller.CandleMean{}

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal("fopen error:", err)
		return candleMean, err
	}

	dec := gob.NewDecoder(f) // Will read from network.
	err = dec.Decode(&candleMean)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	return candleMean, err
}

func main() {
	logFile, err := os.OpenFile(lib.LogPath+"/candle-runner-error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Started")
	telegram.SendMessage("Started: Candle Runner")

	BaseCoin = "USDT"

	for range time.Tick(time.Second * 30) {
		ticker, err := coindcx.GetExchange()
		if err != nil {
			log.Println(err)
			continue
		}

		okMarket := make([]string, 0)
		for _, t := range ticker {
			if !strings.Contains(t.Market, BaseCoin) {
				continue
			}
			candleMean, err := getStoredCandle(t.Market)
			if err != nil {
				continue
			}
			if candleMean.Min < 4 || candleMean.VariencePer > 6 {
				continue
			}
			LastPrice, _ := strconv.ParseFloat(t.LastPrice, 64)
			if (candleMean.Mean + candleMean.Variance) > LastPrice {
				okMarket = append(okMarket, t.Market)

			}
		}
		if len(okMarket) > 1 {
			telegram.SendMessage(strings.Join(okMarket, "-"))
		}
	}
}
