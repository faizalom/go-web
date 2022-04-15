package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"helper/api/coindcx"
	"log"
	"os"
	"sync"
	"time"

	"github.com/faizalom/go-web/controllers/marketcontroller"
	"github.com/faizalom/go-web/lib"
)

var BaseCoin string

var wg sync.WaitGroup
var network bytes.Buffer // Stand-in for a network connection

func main() {

	logFile, err := os.OpenFile(lib.LogPath+"/candle-store-error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Started")

	BaseCoin = "USDT"
	marketsDetails, _ := coindcx.GetMarketsDetails()
	for _, v := range marketsDetails {
		if v.BaseCurrencyShortName == "USDT" || v.TargetCurrencyShortName == "USDT" {
			wg.Add(1)
			go GrabCandle(v)
		}
	}
	wg.Wait()

	for range time.Tick(time.Hour * 1) {
		for _, v := range marketsDetails {
			if v.BaseCurrencyShortName == "USDT" || v.TargetCurrencyShortName == "USDT" {
				wg.Add(1)
				go GrabCandle(v)
			}
		}
		wg.Wait()
	}
}

func GrabCandle(v coindcx.MarketsDetails) {
	defer wg.Done()
	candleMean := marketcontroller.CandleMean{}
	marketcontroller.GetCandles(v.Pair, &candleMean, "5")

	enc := gob.NewEncoder(&network) // Will write to network.
	err := enc.Encode(candleMean)
	if err != nil {
		log.Println("encode error:", err)
		return
	}

	f, err := os.Create(lib.TempCandPath + "/" + v.Symbol)
	if err != nil {
		log.Println(err)
		return
	}
	f.Write(network.Bytes())

	_, err = lib.MDB.Candle5Model().InsertOne(context.Background(), candleMean)
	if err != nil {
		log.Println(err)
	}
}
