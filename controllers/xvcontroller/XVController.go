package xvcontroller

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"helper/model"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/faizalom/go-web/lib"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var NilDate time.Time

func XVIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Println("%s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	//for k, v := range r.Header {
	//fmt.Print("Header field %q, Value %q\n", k, v)
	// 	fmt.Println(k, "\t", v)
	// }

	// fmt.Println("Host = %q\n", r.Host)
	// fmt.Println("RemoteAddr= %q\n", r.RemoteAddr)
	// //Get value for a specified token
	// fmt.Println("\n\nFinding value of \"Accept\" %q", r.Header["Accept"])

	/*rand*/
	// pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 100}}}}}
	// opts1 := options.Aggregate().SetMaxTime(2 * time.Second)
	// cursor, err := model.XVidModel().Aggregate(context.TODO(), pipeline, opts1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	/*rand*/

	param := r.URL.Query()
	limit, _ := strconv.Atoi(param.Get("limit"))
	if limit == 0 {
		limit = 24
	}

	page, _ := strconv.Atoi(param.Get("page"))
	skip := page * limit

	sort, _ := strconv.Atoi(param.Get("sort"))
	if sort == -1 {
		sort = -1
	} else {
		sort = 1
	}

	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(skip))
	opts.SetSort(bson.M{"_id": sort})

	filter := bson.D{
		{"dates.verifiedAt", bson.M{"$ne": NilDate}},
		{"dates.deletedAt", NilDate},
	}

	var category []string
	categoryStr := param.Get("category")
	if categoryStr != "" {
		category = strings.Split(categoryStr, ",")
	}
	if len(category) > 0 {
		filter = append(filter, bson.E{"tags", bson.M{"$in": category}})
	}

	var tags []string
	tagsStr := param.Get("tags")
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}
	if len(tags) > 0 {
		filter = append(filter, bson.E{"tags", bson.M{"$in": tags}})
	}

	var stars []string
	starStr := param.Get("stars.url")
	if starStr != "" {
		stars = strings.Split(starStr, ",")
	}
	if len(stars) > 0 {
		filter = append(filter, bson.E{"stars.url", bson.M{"$in": stars}})
	}

	var uploaderURL []string
	uploaderStr := param.Get("uploader.url")
	if uploaderStr != "" {
		uploaderURL = strings.Split(uploaderStr, ",")
	}
	if len(uploaderURL) > 0 {
		filter = append(filter, bson.E{"uploader.url", bson.M{"$in": uploaderURL}})
	}

	cursor, err := lib.MDB.XVidModel().Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err)
	}

	var sxv []model.XVid
	if err = cursor.All(context.TODO(), &sxv); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(sxv)
	if err != nil {
		log.Println(err)
	}

	// if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
	// 	bs, err := json.Marshal(sxv)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Fprintln(w, string(bs))
	// 	return
	// }
	// data.XVids = sxv
	// t.ExecuteTemplate(w, "xvid.html", data)
}

func GrabVideo(videoID string, modelXvid *model.XVid) {

	resp, err := http.Get(lib.XVURL + "/video" + videoID + "/vid")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	modelXvid.Dates.DeletedAt = NilDate
	modelXvid.Dates.UpdatedAt = time.Now()
	modelXvid.Dates.VerifiedAt = time.Now()

	if resp.StatusCode == 404 {
		modelXvid.Dates.DeletedAt = time.Now()
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)
		//ioutil.WriteFile("body", body, 777)

		// searchHtml5Player := []string{
		// 	"setVideoTitle",
		// 	"setVidURLLow",
		// 	"setVidURLHigh",
		// 	"setThumbUrl",
		// 	"setThumbUrl169",
		// 	"setThumbSlide",
		// 	"setThumbSlideBig",
		// 	"setThumbSlideMinute",
		// 	"setUploaderName",
		// 	"setVidURL",
		// }

		temp := strings.Split(bodyStr, "\n")
		for _, s := range temp {
			s = strings.TrimSpace(s)

			find := strings.Contains(s, "html5player.setVideoTitle(")
			if find {
				modelXvid.Title = strings.ReplaceAll(s, "html5player.setVideoTitle('", "")
				modelXvid.Title = strings.ReplaceAll(modelXvid.Title, "');", "")
			}

			find = strings.Contains(s, "html5player.setThumbUrl(")
			if find {
				modelXvid.Img.Thumbslll = strings.ReplaceAll(s, "html5player.setThumbUrl('", "")
				modelXvid.Img.Thumbslll = strings.ReplaceAll(modelXvid.Img.Thumbslll, "');", "")
				modelXvid.Img.Thumbs = strings.ReplaceAll(modelXvid.Img.Thumbslll, "/thumbslll/", "/thumbs/")
			}

			find = strings.Contains(s, "html5player.setThumbUrl169(")
			if find {
				modelXvid.Img.Thumbs169Tile = strings.ReplaceAll(s, "html5player.setThumbUrl169('", "")
				modelXvid.Img.Thumbs169Tile = strings.ReplaceAll(modelXvid.Img.Thumbs169Tile, "');", "")
				modelXvid.Img.Thumbs169 = strings.ReplaceAll(modelXvid.Img.Thumbs169Tile, "/thumbs169lll/", "/thumbs169lll/")
				modelXvid.Img.Thumbs169 = strings.ReplaceAll(modelXvid.Img.Thumbs169Tile, "/thumbs169poster/", "/thumbs169/")
			}

			find = strings.Contains(s, "html5player.setThumbSlide(")
			if find {
				modelXvid.Img.Slide = strings.ReplaceAll(s, "html5player.setThumbSlide('", "")
				modelXvid.Img.Slide = strings.ReplaceAll(modelXvid.Img.Slide, "');", "")
				modelXvid.Vid.VidPreview = strings.ReplaceAll(modelXvid.Img.Slide, "/thumbs169/", "/videopreview/")
				modelXvid.Vid.VidPreview = strings.ReplaceAll(modelXvid.Vid.VidPreview, "/mozaique.jpg", "_169.mp4")
			}

			find = strings.Contains(s, "html5player.setThumbSlideBig(")
			if find {
				modelXvid.Img.SlideBig = strings.ReplaceAll(s, "html5player.setThumbSlideBig('", "")
				modelXvid.Img.SlideBig = strings.ReplaceAll(modelXvid.Img.SlideBig, "');", "")
			}

			find = strings.Contains(s, "html5player.setThumbSlideMinute(")
			if find {
				modelXvid.Img.SlideMinute = strings.ReplaceAll(s, "html5player.setThumbSlideMinute('", "")
				modelXvid.Img.SlideMinute = strings.ReplaceAll(modelXvid.Img.SlideMinute, "');", "")
			}

			find = strings.Contains(s, "html5player.setVideoUrlLow")
			if find {
				modelXvid.Vid.VidURLLow = strings.ReplaceAll(s, "html5player.setVideoUrlLow('", "")
				modelXvid.Vid.VidURLLow = strings.ReplaceAll(modelXvid.Vid.VidURLLow, "');", "")
			}

			find = strings.Contains(s, "html5player.setVideoUrlHigh(")
			if find {
				modelXvid.Vid.VidURLHigh = strings.ReplaceAll(s, "html5player.setVideoUrlHigh('", "")
				modelXvid.Vid.VidURLHigh = strings.ReplaceAll(modelXvid.Vid.VidURLHigh, "');", "")
			}

			find = strings.Contains(s, "<script>var video_related=")
			if find {
				modelXvid.VideoRelated = []model.XVideoRelated{}
				str := strings.ReplaceAll(s, "<script>var video_related=", "")
				re := regexp.MustCompile(`;window.wpn_categories =.*`)
				bvr := re.ReplaceAll([]byte(str), []byte(""))

				srv := []model.VideoRelated{}
				json.Unmarshal(bvr, &srv)
				//fmt.Printf("%+v", srv)
				for _, v := range srv {
					vr := model.XVideoRelated{}
					vr.VideoId = v.VideoId
					vr.URL = v.URL
					vr.Thumbs169 = v.Thumbs169
					vr.Title = v.Title
					vr.Duration = v.Duration
					vr.Rate = v.Rate
					vr.Size = v.Size
					vr.Uploader.Slug = v.ProfileSlug
					vr.Uploader.Name = v.ProfileName
					vr.Uploader.URL = v.ProfileURL
					modelXvid.VideoRelated = append(modelXvid.VideoRelated, vr)
				}
			}

			find = strings.Contains(s, `<div class="video-metadata video-tags-list ordered-label-list cropped`)
			if find {
				metadata := s

				reg, err := regexp.Compile(`<div class="video-metadata video-tags-list ordered-label-list cropped.*<ul>`)
				if err != nil {
					log.Fatal(err)
				}
				safe := reg.ReplaceAllString(metadata, "<ul>")
				reg = regexp.MustCompile(`</ul>.*`)
				safe = reg.ReplaceAllString(safe, "</ul>")
				//fmt.Println(safe)

				modelXvid.Stars = []struct {
					URL         string `json:"url"`
					Name        string `json:"name"`
					ID          string `json:"id"`
					Profile     string `json:"profile"`
					Subscribers string `json:"subscribers"`
					Verified    bool   `json:"verified"`
				}{}
				modelXvid.Tags = []string{}

				ul := model.Ul{}
				xml.Unmarshal([]byte(safe), &ul)
				for _, v := range ul.Li {

					find = strings.Contains(v.A.Class, "uploader-tag")
					if find {
						modelXvid.Uploader.URL = v.A.Href
						for _, v := range v.A.Span {
							if v.Class == "name" {
								modelXvid.Uploader.Name = v.Text
							}

							if strings.Contains(v.Class, "verified") {
								modelXvid.Uploader.Verified = true
							}
						}
					}

					find = strings.Contains(v.A.Class, "profile")
					if find {
						star := struct {
							URL         string `json:"url"`
							Name        string `json:"name"`
							ID          string `json:"id"`
							Profile     string `json:"profile"`
							Subscribers string `json:"subscribers"`
							Verified    bool   `json:"verified"`
						}{
							URL: v.A.Href,
							//Name: v.A.Span.Text,
						}
						for _, v := range v.A.Span {
							if v.Class == "name" {
								star.Name = v.Text
							}

							if strings.Contains(v.Class, "verified") {
								star.Verified = true
							}
						}
						modelXvid.Stars = append(modelXvid.Stars, star)
					}

					find = strings.Contains(v.A.Href, "/tags/")
					if find {
						modelXvid.Tags = append(modelXvid.Tags, v.A.Text)
					}
				}
				//<strong id="nb-views-number">8,677,903</strong> views	</span><strong class="nb-views-number">9M</strong></span><span class="vote-actions"><a class="btn btn-default vote-action-good"><span class="icon thumb-up black black-hover">&nbsp;</span><span class="rating-inbtn hide-if-zero-11,685">8.7k</span></a><a class="btn btn-default vote-action-bad"><span class="icon thumb-down grey black-hover">&nbsp;</span><span class="rating-inbtn hide-if-zero-11,685">3k</span></a></span></div><ul class="tab-buttons"><li><a class="btn btn-default" id="tabComments_btn"><span class="icon comments-small nb-video-comments-26"></span><span class="visible-lg-inline"> Comments</span><span class="navbadge nb-video-comments nb-video-comments-26 ">26</span></a></li><li><a class="btn btn-default"><span class="icon download"></span><span class="visible-lg-inline"> Download</span></a></li><li><a class="btn btn-default"><span class="icon favorites-add"></span><span class="visible-lg-inline"> Add to my favorites</span></a></li><li><a class="btn btn-default"><span class="icon report"></span><span class="visible-lg-inline"> Report</span></a></li><li><a class="btn btn-default"><span class="visible-lg-inline"><span class="icon-f icf-embed"></span> Embed</span><span class="seperator visible-lg-inline-block">/</span><span class="icon share-small"></span><span class="visible-lg-inline"> Share</span></a></li></ul><a id="watch-later-video" class="btn btn-default visible-lg-inline-block"><span class="icon-f icf-clock"></span> Watch later</a></div><div class="tabs overflow"><div id="tabComments" class="tab overflow"></div><div id="tabDownload" class="tab"></div><div id="tabFavs" class="tab"><h4 class="bg-title grey-dark top no-border">Add this video to one of my favorites list:</h4><div id="favs-container"></div></div><div id="tabReport" class="tab"><h4 class="bg-title grey-dark top no-border">Report this video:</h4></div><div id="tabShareAndEmbed" class="tab"><h4 class="bg-title grey-dark top no-border">Copy page link</h4><div class="copy-link"><label for="copy-video-link" class="btn btn-default copy-btn">Copy</label><input id="copy-video-link" type="text" readonly value="http://www.xvideos.com/video25192641/please_fuck_us_three_in_the_ass" class="form-control"></div><h4 class="bg-title grey-dark no-border">Embed this video to your page with this code:</h4><div class="copy-link"><label for="copy-video-embed" class="btn btn-default copy-btn">Copy</label><input id="copy-video-embed" type="text" readonly value="&lt;iframe src=&quot;https://www.xvideos.com/embedframe/25192641&quot; frameborder=0 width=510 height=400 scrolling=no allowfullscreen=allowfullscreen&gt;&lt;/iframe&gt;" class="form-control"></div><h4 class="bg-title grey-dark no-border">Share this video:</h4></div></div></div>			<div id="ad-footer2" class="mobile-only-show"></div>
				//fmt.Println(v.A.Class)
				// reg = regexp.MustCompile(`<li>.*?(</li>)`)
				// slice := reg.FindAllString(safe, -1)
				// for _, v := range slice {
				// 	fmt.Println(v)
				// }
				// fmt.Println(slice)
				// fmt.Println(len(slice))
			}

			find = strings.Contains(s, `<h2 class="page-title">`)
			if find {
				h2 := model.H2{}
				xml.Unmarshal([]byte(s), &h2)
				for _, v := range h2.Span {
					if v.Class == "duration" {
						modelXvid.Duration = v.Text
					}
					if v.Class == "video-hd-mark" {
						modelXvid.Height = v.Text
					}
				}
			}
		}
	}
}

func XVidGetVideo(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	videoID := ps.ByName("videoId")
	id, err := strconv.Atoi(videoID)
	if err != nil {
		log.Println(err)
	}

	var xv model.XVid
	lib.MDB.XVidModel().FindOne(context.Background(), bson.D{{"_id", int(id)}}).Decode(&xv)
	GrabVideo(videoID, &xv)

	//update := bson.D{{"$set", xv}}
	//lib.MDB.XVidModel().UpdateOne(context.TODO(), bson.D{{"_id", int(id)}}, update)

	//opts := options.FindOneAndReplace().SetUpsert(true)
	//lib.MDB.XVidModel().FindOneAndReplace(context.TODO(), bson.D{{"videoId", int(id)}}, xv, opts)
	// var replacedDocument bson.M
	// e := models.CXVid.FindOneAndReplace(context.TODO(), bson.D{{"videoId", int(id)}}, xv, opts).Decode(&replacedDocument)
	// fmt.Println(e, replacedDocument)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(xv)
	if err != nil {
		log.Println(err)
	}
}
