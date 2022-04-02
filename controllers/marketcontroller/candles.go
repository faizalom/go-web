package marketcontroller

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"helper/api/coindcx"
	"log"
	"math"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var Candles []map[string]interface{}
var csvwriter *csv.Writer

// type Candle struct {
//     Open       float64   `json:"open"`
//     High       float64   `json:"high"`
//     Low        float64   `json:"low"`
//     LowHighPer float64   `json:"low_high_per"`
//     Volume     int       `json:"volume"`
//     Close      float64   `json:"close"`
//     Time       int64     `json:"time"`
//     Timestamp  time.Time `json:"timestamp"`
// }

func allCoin() {
	csvFile, err := os.Create("all.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter = csv.NewWriter(csvFile)
	csvwriter.Write([]string{"pair", "mean", "variance", "VariencePer", "MinMargin", "MaxMargin"})

	BaseCoin := "USDT"
	marketsDetails, _ := coindcx.GetMarketsDetails()
	for _, t := range marketsDetails {
		if !strings.Contains(t.Symbol, BaseCoin) {
			continue
		}
		wg.Add(1)
		//go GetCandles(t.Pair)
	}
	wg.Wait()
	csvwriter.Flush()
	csvFile.Close()
}

func getVariance(datas ...float64) (float32, float32, float32) {
	var sum float64
	for _, v := range datas {
		sum += v
	}
	mean := sum / float64(len(datas))
	var xim2 float64
	for _, v := range datas {
		xim2 += math.Pow(v-mean, 2)
	}
	Varience := math.Sqrt(xim2 / float64(len(datas)))
	VariencePer := (Varience / mean) * 100
	return float32(mean), float32(Varience), float32(VariencePer)
}

type CandleMean struct {
	Mean        float32
	Variance    float32
	VariencePer float32
	Min         float64
	SecMin      float64
	Max         float64
	Candles10   []coindcx.Candle
}

func GetCandles(pair string, candleMean *CandleMean) {
	var network bytes.Buffer // Stand-in for a network connection
	now := time.Now()        // current local time
	sec := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	candles10 := []coindcx.Candle{}

	fname := "candles/" + fmt.Sprint(sec.Unix()*1000) + "_" + pair
	if _, err := os.Stat(fname); err == nil {
		f, err := os.Open(fname)
		if err != nil {
			log.Fatal("fopen error:", err)
		}

		dec := gob.NewDecoder(f) // Will read from network.
		err = dec.Decode(&candles10)
		if err != nil {
			log.Fatal("decode error:", err)
		}
	} else {
		candles10, _ = coindcx.GetCandles(pair, "1d", "10")
		// open output file
		f, err := os.Create("candles/" + fmt.Sprint(candles10[0].Time) + "_" + pair)
		if err != nil {
			log.Println(err)
		}

		enc := gob.NewEncoder(&network) // Will write to network.
		err = enc.Encode(candles10)
		if err != nil {
			log.Println("encode error:", err)
		}
		f.Write(network.Bytes())
	}

	x10Low := make([]float64, 0, 10)
	for _, v := range candles10 {
		x10Low = append(x10Low, v.Low)
	}
	mean, variance, VariencePer := getVariance(x10Low...)
	candleMean.Min = GetMinPercent(candles10)
	candleMean.SecMin = GetMinPercent(candles10)
	candleMean.Max = GetMaxPercent(candles10)

	candleMean.Mean = mean
	candleMean.Variance = variance
	candleMean.VariencePer = VariencePer
	candleMean.Candles10 = candles10
}

func GetMinPercent(candles []coindcx.Candle) float64 {
	min := candles[0].LowHighPer
	for _, c := range candles {
		if c.LowHighPer < min {
			min = c.LowHighPer
		}
	}
	return min
}

func GetMaxPercent(candles []coindcx.Candle) float64 {
	max := candles[0].LowHighPer
	for _, c := range candles {
		if c.LowHighPer > max {
			max = c.LowHighPer
		}
	}
	return max
}
