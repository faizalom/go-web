package main

import (
	"fmt"
	"helper/api/coindcx"
	"time"
)

func main() {

	// tname := "/home/ubuntu/coindcx_candle/test"
	// f, err := os.Create(tname)
	// if err != nil {
	// 	telegram.SendMessage(fmt.Sprint(err))
	// }
	// f.Write([]byte("Started"))

	for range time.Tick(time.Second * 3) {
		//sec := time.Now()
		pair := "B-BTC_USDT"
		candles, _ := coindcx.GetCandles(pair, "1m", "2")

		fmt.Println(candles)

		// b, _ := json.Marshal(candles)
		// fname := "/home/ubuntu/coindcx_candle/" + fmt.Sprint(sec.Year(), sec.Month(), sec.Day(), sec.Hour(), sec.Minute(), sec.Second()) + "_" + pair + ".json"
		// f, err := os.Create(fname)
		// if err != nil {
		// 	telegram.SendMessage(fmt.Sprint(err))
		// }
		// f.Write(b)
	}
}
