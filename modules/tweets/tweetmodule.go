package tweetmodule

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

type tweet struct {
	Id      int
	Message string

	Platform string
}

var tweets = map[int]tweet{
	1: {123, "Anyone there", "iOS"},
}

func New() {
	fmt.Println("Setup Tweets")
	tweets := web.New()
	goji.Handle("/tweets/*", tweets)
	tweets.Use(middleware.SubRouter)

	tweets.Get("/", helloTweet)
	tweets.Get("/:id", getTweetById)
}

func helloTweet(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Look how much I tweet"))
}

func getTweetById(ctx web.C, w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(ctx.URLParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tweet := tweets[id]

	js, err := json.Marshal(tweet)

	if err != nil {
		w.Write([]byte("oops" + err.Error()))
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
