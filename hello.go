package main

import(
  "fmt"
  "net/http"
)

func handleReq(w http.ResponseWriter, req *http.Request){
    site := req.URL.Query().Get("site")
  nextURL := site

  var i int
  for i < 100 {
    client := &http.Client{
      CheckRedirect: func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
    } }

    resp, err := client.Get(nextURL)

    // if err != nil {
    //   fmt.Println(err)
    // }
    if resp.StatusCode == 404 {
        fmt.Println(err)
      }

    fmt.Println("StatusCode:", resp.StatusCode)
    fmt.Println(resp.Request.URL)

    if resp.StatusCode == 200 {
      fmt.Println("Done!")
      break
    } else {
      nextURL = resp.Header.Get("Location")
      i += 1
    }
  }
}


func main(){

    http.HandleFunc("/", handleReq)

    http.ListenAndServe(":8080", nil)



}