package main

import (
	"fmt"
	"github.com/larsha/go-template/cache"
	"github.com/larsha/go-template/middlewares"
	"github.com/larsha/go-template/utils"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := cache.Set("foo", "bar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val, err := cache.Get("foo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	add := strconv.Itoa(utils.Add(1, 2))

	fmt.Fprintf(w, "%s %s", val, add)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", middlewares.Logger(http.DefaultServeMux))
}
