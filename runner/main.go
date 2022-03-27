package main

import (
	"context"
	"fmt"
	"helper/model"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/faizalom/go-web/controllers"
	"github.com/faizalom/go-web/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var NilDate time.Time
var verifyStartedAt time.Time
var wg sync.WaitGroup

func GetXV() {

	opts := options.Find()
	opts.SetLimit(int64(1000))
	opts.SetSort(bson.M{"_id": -1})

	filter := bson.D{
		{"dates.verifiedAt", bson.M{"$lt": verifyStartedAt}},
		{"dates.deletedAt", NilDate},
	}

	cursor, err := lib.MDB.XVidModel().Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}

	var sxv []model.XVid
	if err = cursor.All(context.TODO(), &sxv); err != nil {
		log.Fatal(err)
	}

	for _, s := range sxv {
		wg.Add(1)
		go updateXv(s)
	}
	wg.Wait()
}

func updateXv(xv model.XVid) {
	defer wg.Done()
	VideoID := strconv.Itoa(xv.VideoID)
	controllers.GrabVideo(VideoID, &xv)

	if xv.Dates.DeletedAt != NilDate {
		// bxv, _ := json.Marshal(xv)
		// s := string(bxv)
		// fmt.Println(s)
	} else {
		xv.URL = lib.XVURL + "/video" + VideoID + "/" + xv.URLName
	}

	update := bson.M{"$set": xv}
	_, err := lib.MDB.XVidModel().UpdateOne(context.TODO(), bson.M{"_id": xv.VideoID}, update)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(res)

	fmt.Println(VideoID, "\t", xv.Dates.VerifiedAt, "\t", xv.Dates.DeletedAt)
}

func main() {
	verifyStartedAt = time.Date(2022, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	GetXV()
}
