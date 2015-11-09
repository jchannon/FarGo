package main

import (
	"net/http"

	"github.com/jchannon/FarGo/modules/tweets"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web/middleware"
)

func main() {
	goji.Use(middleware.EnvInit)
	tweetmodule.New()

	goji.Get("/hello", HomeHandler)
	goji.Get("/*", http.FileServer(http.Dir("static")))

	goji.Serve()

}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
