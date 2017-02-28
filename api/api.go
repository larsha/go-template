package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Store maps data from json api to struct
type Store struct {
	ID                string   `json:"nr"`
	City              string   `json:"city"`
	Name              string   `json:"name"`
	Address           string   `json:"address"`
	AdditionalAddress string   `json:"additional_address"`
	ZipCode           string   `json:"zip_code"`
	County            string   `json:"county"`
	Phone             string   `json:"phone"`
	ShopType          string   `json:"shop_type"`
	Services          string   `json:"services"`
	Labels            []string `json:"labels"`
	OpeningHours      []string `json:"opening_hours"`
	RT90X             int      `json:"RT90x"`
	RT90Y             int      `json:"RT90y"`
}

// API struct containing "constructor" variables
type API struct {
	URL     string
	Timeout time.Duration
}

// New initialises API struct with default data
func New() *API {
	return &API{
		URL:     "https://bolaget.io",
		Timeout: time.Duration(10 * time.Second),
	}
}

// GetStores returns slices of Store
func (a *API) GetStores() ([]Store, error) {
	body, err := a.request(a.URL + "/stores")
	if err != nil {
		return nil, err
	}
	s, err := parseStores(body)
	return s, err
}

func (a *API) request(url string) ([]byte, error) {
	client := http.Client{
		Timeout: a.Timeout,
	}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return []byte(body), err
}

func parseStores(body []byte) ([]Store, error) {
	var s = make([]Store, 0)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
