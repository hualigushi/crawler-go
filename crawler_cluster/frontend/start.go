package main

import (
	"crawler/frontend/controler"
	"net/http"
)

func main() {
	http.Handle("/",http.FileServer(
		http.Dir("frontend/view"),
	))
	http.Handle("/search",
		controler.CreateSearchResultHandler(
			"frontend/view/template.html",
		))
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
}
