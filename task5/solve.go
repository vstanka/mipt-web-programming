package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var urls map[string]string

func generateShortLink() (result string) {
	count := len(urls) + 1
	for count > 0 {
		result += string('a' + (count % 26))
		count /= 26
	}
	return
}

func shorterHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var body []byte
		defer req.Body.Close()
		body, _ = ioutil.ReadAll(req.Body)
		var req map[string]string
		err := json.Unmarshal(body, &req)
		if err == nil {
			url, ok := req["url"]
			if ok {
				key := generateShortLink()
				urls[key] = url
				respMap := make(map[string]string)
				respMap["key"] = key
				resp, _ := json.Marshal(respMap)
				rw.Write(resp)
				return
			}
		}
		http.Error(rw, "", 400)
	} else {
		url, ok := urls[req.RequestURI[1:]]
		if !ok {
			http.NotFound(rw, req)
			return
		}
		http.Redirect(rw, req, url, 301)
	}

}

func main() {
	urls = make(map[string]string)
	http.HandleFunc("/", shorterHandler)
	http.ListenAndServe(":8082", nil)
}