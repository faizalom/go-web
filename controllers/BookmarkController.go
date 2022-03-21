package controllers

import (
	"encoding/json"
	"helper/model"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Children struct {
	Children     []Children
	DateAdded    string `json:"date_added"`
	DateModified string `json:"date_modified"`
	ID           string `json:"id"`
	URL          string `json:"url"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type ChormeBookmark struct {
	Checksum string `json:"checksum"`
	Roots    struct {
		BookmarkBar struct {
			Children     []Children
			DateAdded    string `json:"date_added"`
			DateModified string `json:"date_modified"`
			ID           string `json:"id"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"bookmark_bar"`
		Synced struct {
			Children
		} `json:"synced"`
	} `json:"roots"`
}

func XVControllerImporter(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//dat, err := ioutil.ReadFile("/home/faizal/.config/google-chrome/Default/Bookmarks")
	dat, err := ioutil.ReadFile("C:/Users/Faizal/AppData/Local/Google/Chrome/User Data/Default/Bookmarks")
	if err != nil {
		log.Println(err)
	}
	bk := ChormeBookmark{}
	json.Unmarshal(dat, &bk)

	c := bk.Roots.BookmarkBar.Children
	sxv = make([]model.XVid, 0)
	spr = make([]model.XProfile, 0)
	extractBookmark(c)
	//model.XVids().InsertUpdate(sxv)
	//model.XProfiles().InsertUpdate(spr)
}

var i int = 0
var folder []string

var sxv []model.XVid
var spr []model.XProfile

func extractBookmark(c []Children) {
	//sxv := make([]model.XVid, 0)
	for _, v := range c {

		if v.Type == "folder" {
			folder = append(folder, v.Name)
			extractBookmark(v.Children)
		}
		find := strings.Contains(v.URL, "https://www.xvideos2.com/")
		if find {
			v.Name = strings.Trim(strings.Replace(v.Name, " - XVIDEOS.COM", "", 1), " ")

			if strings.Contains(v.URL, "https://www.xvideos2.com/video") {
				URL := strings.Replace(v.URL, "https://www.xvideos2.com/video", "", 1)

				re := regexp.MustCompile(`/.*`)
				bVID := re.ReplaceAll([]byte(URL), []byte(""))
				VID, _ := strconv.ParseInt(string(bVID), 10, 32)

				re = regexp.MustCompile(`.*/`)
				bURLName := re.ReplaceAll([]byte(URL), []byte(""))

				i, err := strconv.ParseInt(v.ID, 10, 32)
				if err != nil {
					log.Println(err)
				}
				id := int32(i)

				xv := model.XVid{
					VideoID:    int(VID),
					URL:        v.URL,
					URLName:    string(bURLName),
					Title:      v.Name,
					BookmarkID: id,
					Folder:     strings.Join(folder[:], "/"),
					Dates: struct {
						CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
						UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
						DeletedAt  time.Time `json:"deletedAt" bson:"deletedAt"`
						VerifiedAt time.Time `json:"verifiedAt" bson:"verifiedAt"`
					}{
						CreatedAt: fileTimeConverter(v.DateAdded),
					},
				}
				sxv = append(sxv, xv)
				//os.Exit(1)
			} else if strings.Contains(v.URL, "https://www.xvideos2.com/profiles") ||
				strings.Contains(v.URL, "https://www.xvideos2.com/model") ||
				strings.Contains(v.URL, "https://www.xvideos2.com/pornstars") ||
				strings.Contains(v.URL, "https://www.xvideos2.com/channels") ||
				strings.Contains(v.URL, "https://www.xvideos2.com/model-channels") ||
				strings.Contains(v.URL, "https://www.xvideos2.com/amateur-channels") ||
				strings.Contains(v.URL, "https://www.xvideos2.com/pornstar-channels") {

				s := strings.Split(v.URL, "/")

				re := regexp.MustCompile(`#.*`)
				slug := re.ReplaceAll([]byte(s[4]), []byte(""))

				re = regexp.MustCompile(`-.*`)
				bT := re.ReplaceAll([]byte(v.Name), []byte(""))
				title := strings.Trim(string(bT), " ")

				v.Name = strings.Trim(strings.Replace(v.Name, " - XVIDEOS.COM", "", 1), " ")

				pr := model.XProfile{
					URL:    v.URL,
					Type:   s[3],
					Slug:   string(slug),
					Name:   title,
					BID:    v.ID,
					Folder: strings.Join(folder[:], "/"),
					Dates: model.Dates{
						CreatedAt: fileTimeConverter(v.DateAdded),
					},
				}
				spr = append(spr, pr)
				//os.Exit(1)
			} else if strings.Contains(v.URL, "https://www.xvideos2.com/tags") {
			} else if strings.Contains(v.URL, "https://www.xvideos2.com/?k=") {
			} else {
				log.Println(strings.Join(folder[:], "/"), v.Name, v.URL)
			}
		}
	}
	if len(folder) > 0 {
		folder = folder[:len(folder)-1] // [A B]
	}
}

func fileTimeConverter(bkTime string) time.Time {
	fileTime, err := strconv.ParseInt(bkTime+"0", 10, 64)
	if err != nil {
		panic(err)
	}
	winSecs := fileTime / 10000000           // divide by 10 000 000 to get seconds
	unixTimestamp := (winSecs - 11644473600) // 1.1.1600 -> 1.1.1970 difference in seconds

	return time.Unix(unixTimestamp, 0)
}
