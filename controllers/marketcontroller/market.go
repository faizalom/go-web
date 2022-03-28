package marketcontroller

import (
	"encoding/json"
	"fmt"
	"helper/api/coindcx"
	"helper/api/telegram"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type MarketController struct{}

func NewMarketController() *MarketController {
	return &MarketController{}
}

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

func (c MarketController) MarketCoin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	MarketName := ps.ByName("market")
	MarketsDetail := coindcx.MarketsDetails{}

	marketsDetails, _ := coindcx.GetMarketsDetails()
	for _, t := range marketsDetails {
		if t.CoindcxName == MarketName {
			MarketsDetail = t
		}
	}

	Ticker, TickerCalc, _ := coindcx.GetTicker(MarketName)
	candles10, _ := coindcx.GetCandles(MarketsDetail.Pair, "1d", "12")
	//Get Varience
	x10Low := make([]float64, 0, 10)
	x10Hig := make([]float64, 0, 10)
	for _, v := range candles10 {
		x10Low = append(x10Low, v.Low)
		x10Hig = append(x10Hig, v.High)
	}
	meanLow, varianceLow, variencePerLow := getVariance(x10Low...)
	meanHigh, varianceHigh, variencePerHigh := getVariance(x10Hig...)
	//mean, variance, VariencePer := getVariance(9, 2, 5, 4, 12, 7, 8, 11, 9, 3, 7, 4, 12, 5, 4, 10, 9, 6, 9, 4)

	candles10Statics := struct {
		MeanMin        float32
		VarianceMin    float32
		VariencePerMin float32
		MeanMax        float32
		VarianceMax    float32
		VariencePerMax float32
	}{
		meanLow,
		varianceLow,
		variencePerLow,
		meanHigh,
		varianceHigh,
		variencePerHigh,
	}

	candlesToday, _ := coindcx.GetCandles(MarketsDetail.Pair, "30m", "48")
	candlesLast, _ := coindcx.GetCandles(MarketsDetail.Pair, "1m", "60")

	d := coindcx.GetCoinBalance(MarketsDetail.TargetCurrencyShortName)
	mData := make(map[string]string)
	mData["coinBalance"] = fmt.Sprintf("%f", d.Balance)
	LastPrice, _ := strconv.ParseFloat(Ticker.LastPrice, 64)
	mData["conversionBalance"] = fmt.Sprintf("%f", d.Balance*LastPrice)
	mData["BuyCoin"] = "0"
	mData["Watch"] = "0"
	mData["BuyCoinBuyPrice"] = "0"
	mData["BuyCoinSellPrice"] = "0"
	mData["BuyCoinTotal"] = "0"
	mData["Profit"] = "0"
	mData["ProfitMargin"] = "0"
	if BuyCoin[MarketName] {
		mData["BuyCoin"] = "1"
		mData["BuyCoinBuyPrice"] = fmt.Sprintf("%f", BuyCoinBuyPrice[MarketName])
		mData["BuyCoinSellPrice"] = fmt.Sprintf("%f", BuyCoinSellPrice[MarketName])
		Profit := (d.Balance * LastPrice) - (BuyCoinBuyPrice[MarketName] * d.Balance)
		mData["Profit"] = fmt.Sprintf("%f", Profit)
		mData["ProfitMargin"] = fmt.Sprintf("%.2f", Profit/(d.Balance*LastPrice)*100)
	}
	if WatchList[MarketName] {
		mData["Watch"] = "1"
	}

	data := struct {
		BaseCoin         string
		Ticker           interface{}
		TickerCalc       interface{}
		MarketsDetail    interface{}
		Candles10        interface{}
		Candles10Statics interface{}
		CandlesToday     interface{}
		CandlesLast      interface{}
		MarketName       string
		Data             map[string]string
	}{
		BaseCoin,
		Ticker,
		TickerCalc,
		MarketsDetail,
		candles10,
		candles10Statics,
		candlesToday,
		candlesLast,
		MarketName,
		mData,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}

	// Path := "/home/ubuntu/coindcx"
	// if runtime.GOOS == "windows" {
	// 	Path, _ = os.Getwd()
	// }
	// t := adminlte.GetTemplate()
	// t.Funcs(fm).ParseFiles(Path + "/templates/market.html")
	// t.ExecuteTemplate(w, "market.html", data)
}

func (c MarketController) WatchAdd(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.FormValue("coin") != "" {
		watch, _ := strconv.ParseBool(r.FormValue("watch"))
		WatchList[r.FormValue("coin")] = watch
		WatchCoins := ""
		for coin, ok := range WatchList {
			if !ok {
				continue
			}
			WatchCoins += coin + "\n"
		}
		telegram.SendMessage(WatchCoins)
	}
	data := struct {
		BaseCoin         string
		WatchList        interface{}
		BuyCoin          interface{}
		BuyCoinBuyPrice  interface{}
		BuyCoinSellPrice interface{}
	}{
		BaseCoin,
		WatchList,
		BuyCoin,
		BuyCoinBuyPrice,
		BuyCoinSellPrice,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}

	// Path := "/home/ubuntu/coindcx"
	// if runtime.GOOS == "windows" {
	// 	Path, _ = os.Getwd()
	// }
	// t := adminlte.GetTemplate()
	// t.Funcs(fm).ParseFiles(Path + "/templates/watch.html")
	// t.ExecuteTemplate(w, "watch.html", data)
}

func (c MarketController) CoinBuy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	buyCoins := ""
	if r.FormValue("coin") != "" {
		buy, _ := strconv.ParseBool(r.FormValue("buy"))
		buyPrice, _ := strconv.ParseFloat(r.FormValue("buy_price"), 64)
		sellPrice, _ := strconv.ParseFloat(r.FormValue("sell_price"), 64)
		BuyCoin[r.FormValue("coin")] = buy
		BuyCoinBuyPrice[r.FormValue("coin")] = buyPrice
		BuyCoinSellPrice[r.FormValue("coin")] = sellPrice
		if !buy {
			delete(BuyCoin, r.FormValue("coin"))
		}
		for coin, ok := range BuyCoin {
			if !ok {
				continue
			}
			buyCoins += coin + "\n"
		}
		telegram.SendMessageUp(buyCoins)
	}

	data := struct {
		BaseCoin         string
		WatchList        interface{}
		BuyCoin          interface{}
		BuyCoinBuyPrice  interface{}
		BuyCoinSellPrice interface{}
	}{
		BaseCoin,
		WatchList,
		BuyCoin,
		BuyCoinBuyPrice,
		BuyCoinSellPrice,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}

	// Path := "/home/ubuntu/coindcx"
	// if runtime.GOOS == "windows" {
	// 	Path, _ = os.Getwd()
	// }
	// t := adminlte.GetTemplate()
	// t.Funcs(fm).ParseFiles(Path + "/templates/watch.html")
	// t.ExecuteTemplate(w, "watch.html", data)
}

func (c MarketController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	BaseCoin = "USDT"
	ticker, err := coindcx.GetExchange()
	if err != nil {
		log.Println(err)
	}
	//MaxLeverageShort interface{}
	var markets []interface{}
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
		}{
			Ticker:       t,
			Coin:         strings.Replace(t.Market, BaseCoin, "", -1),
			LowNowMargin: lastPrice / low,
			Timestamp:    timeUnix,
			LowHighPer:   lowHighPer,
		}
		markets = append(markets, market)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(markets)
	if err != nil {
		log.Println(err)
	}

	// Path := "/home/ubuntu/coindcx"
	// if runtime.GOOS == "windows" {
	// 	Path, _ = os.Getwd()
	// }
	// t := adminlte.GetTemplate()
	// t.Funcs(fm).ParseFiles(Path + "/templates/dashboard.html")
	// t.ExecuteTemplate(w, "dashboard.html", data)
}
