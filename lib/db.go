package lib

import (
	"github.com/faizalom/go-web/model"
)

var MDB model.MongoDB
var TempCandPath string
var LogPath string

func init() {
	MDB.Database = model.ConnDB()

	// logPath, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// LogPath = logPath

}
