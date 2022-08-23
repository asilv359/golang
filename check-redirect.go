package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RedirectResponse struct {
	Status   string `json:status`
	Url      string `json:url`
	Code     int    `json:code`
	Location string `json:location`
	Server   string `json:server`
}

// type RedirectResponse struct {
// 	Status   string         `json:status`
// 	Response map[string]int `json:response`
// 	Location string         `json:location`
// 	Server   string         `json:server`
// }

// type Redirect struct {
// 	Url  string `json:status`
// 	Code string `json:status`
// }

func handleReq(w http.ResponseWriter, req *http.Request) {

	var server string
	// var location string

	site := req.URL.Query().Get("site")

	redirects := make(map[string]int)

	nextUrl := site
	// var i int
	// for i < 20
	for i := 0; i < 20; {

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}

		resp, err := client.Get(nextUrl)

		if err != nil {
			fmt.Println(err, "nil err from resp")

		}

		if resp.Header.Get("server") != "" {
			server = resp.Header.Get("server")
		} else {

			server = ("Can not get server name")
		}

		// location = resp.Header.Get("Location")

		status := http.StatusText(resp.StatusCode)

		url := resp.Request.URL.String()

		code := resp.StatusCode

		redirects[url] = resp.StatusCode

		if resp.StatusCode == 200 {

			fmt.Println("done")

			redirectStruct := RedirectResponse{status, url, code, nextUrl, server}
			json_data, _ := json.Marshal(redirectStruct)
			w.Write(json_data)

			break
		} else {
			nextUrl = resp.Header.Get("Location")
			// i += i
		}

		if resp.StatusCode == 404 {
			fmt.Println(err, "404 nil")

			redirectStruct := RedirectResponse{status, url, code, nextUrl, server}
			json_data, _ := json.Marshal(redirectStruct)
			w.Write(json_data)

			break
		}
		redirectStruct := RedirectResponse{status, url, code, nextUrl, server}
		json_data, _ := json.Marshal(redirectStruct)
		w.Write(json_data)
		fmt.Fprintf(w, "\n")

	}

	// json_data, _ := json.Marshal(redirects)
	// w.Write(json_data)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

}

func main() {

	http.HandleFunc("/tracer", handleReq)

	http.ListenAndServe("127.0.0.1:8181", nil)
}

// TODO: each redirect its json object
