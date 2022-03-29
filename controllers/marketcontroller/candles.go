package marketcontroller

import (
	"encoding/csv"
	"helper/api/coindcx"
	"log"
	"math"
	"os"
	"strings"
	"sync"
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
	Max         float64
}

func GetCandles(pair string, candleMean *CandleMean) {
	defer wg.Done()
	candles10, _ := coindcx.GetCandles(pair, "1d", "10")
	//fmt.Printf("len=%d cap=%d\n", len(candles10), cap(candles10))
	if len(candles10) == 10 {
		data := make(map[string]interface{})
		data["pair"] = pair
		data["candles"] = candles10
		Candles = append(Candles, data)

		x10Low := make([]float64, 0, 10)
		for _, v := range candles10 {
			x10Low = append(x10Low, v.Low)
		}
		mean, variance, VariencePer := getVariance(x10Low...)
		//mean, variance, VariencePer := getVariance(9, 2, 5, 4, 12, 7, 8, 11, 9, 3, 7, 4, 12, 5, 4, 10, 9, 6, 9, 4)
		candleMean.Min = GetMinPercent(candles10)
		candleMean.Max = GetMaxPercent(candles10)

		candleMean.Mean = mean
		candleMean.Variance = variance
		candleMean.VariencePer = VariencePer

		// candleMean = CandleMean{
		// 	Mean:        mean,
		// 	Variance:    variance,
		// 	VariencePer: VariencePer,
		// 	Min:         min,
		// 	Max:         max,
		// }

		//d := []string{pair, fmt.Sprint(mean), fmt.Sprint(variance), fmt.Sprint(VariencePer), fmt.Sprint(min), fmt.Sprint(max)}
		// err := csvwriter.Write(d)
		// if err != nil {
		// 	log.Fatalf("failed creating file: %s", err)
		// }
	}
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
