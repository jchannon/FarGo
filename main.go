package main

import (
	"net/http"

	"github.com/jchannon/FarGo/modules/tweets"
	"github.com/zenazn/goji"
)

func main() {

	tweetmodule.New()

	goji.Get("/hello", HomeHandler)
	goji.Get("/*", http.FileServer(http.Dir("static")))

	goji.Serve()

}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
