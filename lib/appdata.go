package lib

import (
	"helper/model"
	"theme"
)

var SideMenu = []theme.Navlink{
	{
		Link: "/u/",
		Icon: "fas fa-tachometer-alt",
		Text: "Dashboard",
	},
	// {
	// 	Link: "/u/staff",
	// 	Icon: "fas fa-th",
	// 	Text: "Staff",
	// 	Children: []theme.Navlink{
	// 		{
	// 			Link: "/u/staff",
	// 			Icon: "fas fa-tachometer-alt",
	// 			Text: "List",
	// 		},
	// 		{
	// 			Link: "/u/staff/add",
	// 			Icon: "fas fa-th",
	// 			Text: "Add",
	// 		},
	// 	},
	// },
	{
		Link: "/u/staff",
		Icon: "fas fa-th",
		Text: "Staff",
	},
}

var Theme theme.AdminThemeTemplete
var MDB model.MongoDB

func init() {
	Theme = theme.CoreUITheme
	Theme.SideMenu = SideMenu
	Theme.Title = "🦁FAPP"
	MDB.Database = model.MongoDBLive()
}
