package tweetmodule

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

type tweet struct {
	ID      int
	Message string

	Platform string
}

var tweets = map[int]tweet{
	1: {123, "Anyone there", "iOS"},
	2: {456, "Instagram is better", "Tweetdeck"},
}

//New intialises tweet routes
func New() {
	fmt.Println("Setup Tweets")
	tweets := web.New()
	goji.Handle("/tweets/*", tweets)

	tweets.Use(middleware.SubRouter)
	tweets.Use(conneg)

	tweets.Get("/", helloTweet)
	tweets.Get("/:id", getTweetByID)

}

func helloTweet(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Look how much I tweet"))
}

func getTweetByID(ctx web.C, w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(ctx.URLParams["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tweet := tweets[id]

	ctx.Env["model"] = tweet
}

func conneg(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		h.ServeHTTP(w, r)

		accept := r.Header.Get("Accept")
		model := c.Env["model"]

		switch accept {
		case "application/json":

			w.Header().Set("Content-Type", "application/json")

			js, err := json.Marshal(model)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(js)

		case "application/xml":
			x, err := xml.MarshalIndent(model, "", "  ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/xml")
			w.Write(x)
		}
	}
	return http.HandlerFunc(fn)
}

type ResponseProcessor interface {
	CanProcess(accept string) bool
	Process(model interface{})
}

type JsonProcessor struct {
}

func (*JsonProcessor) CanProcess(accept string) bool {
	return true
}
func (*JsonProcessor) Process(model interface{}) {}
