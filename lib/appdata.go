package lib

import (
	"helper/model"
	"net/http"
	"theme"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	Store = sessions.NewCookieStore(key)
)

var SideMenu = []theme.Navlink{
	{
		Link: "/index",
		Icon: "fas fa-tachometer-alt",
		Text: "Dashboard",
		Children: []theme.Navlink{
			{
				Link: "/index",
				Icon: "fas fa-tachometer-alt",
				Text: "Dashboard",
			},
			{
				Link: "#",
				Icon: "fas fa-th",
				Text: "Simple Link 2",
			},
			{
				Link: "#",
				Icon: "fas fa-th",
				Text: "Simple Link 3",
			},
			{
				Link: "https://www.google.com/",
				Icon: "fas fa-th",
				Text: "Google",
			},
		},
	},
	{
		Link: "#",
		Icon: "fas fa-th",
		Text: "Simple Link 2",
	},
	{
		Link: "/index",
		Icon: "fas fa-tachometer-alt",
		Text: "Dashboard",
	},
	{
		Link: "https://www.google.com/",
		Icon: "fas fa-th",
		Text: "Google",
	},
}

var Theme theme.AdminThemeTemplete
var MDB model.MongoDB

func init() {
	Theme = theme.CoreUITheme
	Theme.SideMenu = SideMenu
	MDB.Database = model.MongoDBLive()
}

func Auth(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "auth")
}

// func GetAuthUser(r *http.Request) model.User {
// 	Auth(r)
// }
