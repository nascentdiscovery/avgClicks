package main

import (
//	"strconv"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

type Client struct {
	ApiUrl *url.URL
	Bearer string
	Client *http.Client
}

func main() {
	Records := make(map[string]int)
	ApiUrl, _ := url.Parse("https://api-ssl.bitly.com/v4")
	client := &Client{
		ApiUrl: ApiUrl,
		Client: http.DefaultClient,
		Bearer: "Bearer " + os.Getenv("BITLY_API_TOKEN"),
	}
	path := "user"
	guid := getJSfromPath(path, client).Get("default_group_guid").MustString()
	path = "groups/" + guid + "/countries"
	i := getJSfromPath(path, client).Get("metrics").MustArray()

	for _, v := range i {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Map {
			set := 0
			value, clicks := "", ""
			for _, key := range rv.MapKeys() {
				strct := rv.MapIndex(key)
				switch key.Interface() {
				case "value":
					value := strct.Interface()
					fmt.Println(value)
					set += 1
				case "clicks":
					clicks := strct.Interface()
					fmt.Println(clicks)
					set += 1
				}
				if set == 2 {
					fmt.Printf("Value: %v, Clicks: %v", value, clicks)
					set = 0
				}

			}
		}

	}
	fmt.Println(Records)

}

func getJSfromPath(path string, client *Client) (js *simplejson.Json) {
	client.ApiUrl.Path = "v4/" + path
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
