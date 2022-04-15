package lib

import (
	"helper/model"
	"log"
	"os"
	"path/filepath"
	"theme"
)

var SideMenu = []theme.Navlink{
	{
		Link: "/u/",
		Icon: "fas fa-tachometer-alt",
		Text: "Dashboard",
	},
	{
		Link: "/u/staff",
		Icon: "fas fa-th",
		Text: "Staff",
	},
	{
		Link: "/u/market",
		Icon: "fas fa-tachometer-alt",
		Text: "Market",
		Children: []theme.Navlink{
			{
				Link: "/u/market",
				Icon: "fas fa-tachometer-alt",
				Text: "Market",
			},
			{
				Link: "/u/great-trade",
				Icon: "fas fa-tachometer-alt",
				Text: "Great Trade",
			},
			{
				Link: "/u/candle-mean",
				Icon: "fas fa-tachometer-alt",
				Text: "Candle Mean",
			},
			{
				Link: "/u/low",
				Icon: "fas fa-tachometer-alt",
				Text: "Low",
			},
		},
	},
	{
		Link: "/u/book",
		Icon: "fas fa-book",
		Text: "Book",
	},
}

var Theme theme.AdminThemeTemplete
var MDB model.MongoDB
var TempCandPath string
var LogPath string

func init() {
	Theme = theme.CoreUITheme
	Theme.SideMenu = SideMenu
	Theme.Title = "🦁FAPP"
	MDB.Database = model.MongoDBLive()

	logPath, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	}
	LogPath = logPath

	// open output file
	TempCandPath = filepath.Join(LogPath, "logs/candles")
	err = os.MkdirAll(TempCandPath, os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}

	LogPath = filepath.Join(LogPath, "logs/coindcx")
	err = os.MkdirAll(LogPath, os.ModePerm)
	if err != nil {
		log.Panicln(err)
	}
}
