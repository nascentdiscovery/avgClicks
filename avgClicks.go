package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
)

type Record struct {
	Value  string
	Clicks int64
}
type Client struct {
	ApiUrl   *url.URL
	Bearer   string
	Client   *http.Client
	Lookback int
}

var lookback int

func init() {
	flag.IntVar(&lookback, "lookback", 30, "How many days of historical data to average")
}

func main() {
	flag.Parse()
	ApiUrl, _ := url.Parse("https://api-ssl.bitly.com/v4")

	client := &Client{
		ApiUrl:   ApiUrl,
		Client:   http.DefaultClient,
		Bearer:   "Bearer " + os.Getenv("BITLY_API_TOKEN"),
		Lookback: lookback,
	}
	path := "user"
	guid := getJSfromPath(path, "", client).Get("default_group_guid").MustString()
	path = "groups/" + guid + "/countries"
	units := strconv.Itoa(client.Lookback)
	i := getJSfromPath(path, units, client).Get("metrics").MustArray()
	cpc := getClicksPerCountry(i)
	fmt.Printf("Record of average clicks per day, by country for the Bitly Group ID %v.\n", guid)
	fmt.Printf("Using %v days worth of data.\n", units)

	for k, v := range cpc {
		fmt.Printf("%v: %v clicks per day\n", k, float64(v)/float64(client.Lookback))
	}
}

func getClicksPerCountry(i []interface{}) (Records map[string]int64) {
	Records = make(map[string]int64)
	rec := &Record{}
	for _, v := range i {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Map {
			set := 0
			for _, key := range rv.MapKeys() {
				strct := rv.MapIndex(key)
				switch key.Interface() {
				case "value":
					rec.Value = strct.Interface().(string)
					set += 1
				case "clicks":
					rec.Clicks, _ = strct.Interface().(json.Number).Int64()
					set += 1
				}
				if set == 2 {
					Records[rec.Value] += rec.Clicks
					set = 0
				}

			}
		}

	}
	return Records

}

func getJSfromPath(path string, units string, client *Client) (js *simplejson.Json) {
	client.ApiUrl.Path = "v4/" + path
	if units != "" {
		q := client.ApiUrl.Query()
		q.Set("units", units)
		client.ApiUrl.RawQuery = q.Encode()
	}

	req, err := http.NewRequest("GET", client.ApiUrl.String(), nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", client.Bearer)
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	njs, err := simplejson.NewJson(body)
	js = njs
	if err != nil {
		log.Fatalln(err)
	}
	return js
}
