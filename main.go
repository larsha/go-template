package main

import (
	"encoding/json"
	"fmt"
	"github.com/larsha/go-template/api"
	"github.com/larsha/go-template/cache"
	"github.com/larsha/go-template/middlewares"
	"net/http"
	"os"
)

type fooBar struct {
	Foo string `json:"foo"`
}

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

	data := fooBar{Foo: val}

	// add := strconv.Itoa(utils.Add(1, 2))

	// fmt.Fprintf(w, "%s %s", val, add)

	json.NewEncoder(w).Encode(data)
}

func main() {
	api := api.New()
	stores, err := api.GetStores()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stores)

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), middlewares.Logger(http.DefaultServeMux))
}
