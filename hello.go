package main

import(
  "fmt"
  "net/http"
  "encoding/json"
  // "strconv"
)
type RedirectsResponse struct {
  Status string `json:"status"`
  Response map[string]int `json:"response"`
}

func handleReq(w http.ResponseWriter, req *http.Request){
  site := req.URL.Query().Get("site")

  // var redirects map[string]int 

  nextURL := site
  redirects :=  make(map[string]int)
  var i int

  for i < 100 {
    client := &http.Client{
      CheckRedirect: func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
        
    } }

    resp, err := client.Get(nextURL)


    url := resp.Request.URL.String()

    // fmt.Fprintln(w, resp, err )
  
    // if err != nil {
    // fmt.Fprintln(w, resp.Request.URL , resp.StatusCode)
      
    // }
    if resp.StatusCode == 404 {
        fmt.Println(err)
        fmt.Fprintln(w, resp.Request.URL , resp.StatusCode)
      }

    fmt.Println("StatusCode:", resp.StatusCode)
    fmt.Println(resp.Request.URL)

    // str := url + strconv.Itoa(i)

    // redirects[str] = 1


    redirects[url] = resp.StatusCode



    if resp.StatusCode == 200 {
      fmt.Println("Done!")
      break
    } else {
      nextURL = resp.Header.Get("Location")
      i += 1
    }

  }
  redirectsStruct := RedirectsResponse{"ok", redirects}

  jsonString, _ := json.Marshal(redirectsStruct)
  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonString)

}


func main(){

    http.HandleFunc("/", handleReq)

    http.ListenAndServe(":8080", nil)



}