package main

import ("fmt"
"net/http"
"io/ioutil")


func main() {
  // want to get information from the internet
  // convert information to a string, then parsing/formatting
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
      // underscores are used anytime that you define a variable that
      // you don't intend to use (initializing a variable)
  bytes, _ := ioutil.ReadAll(resp.Body)
    // converting to a string
  string_body := string(bytes)
  fmt.Println(string_body)
  resp.Body.Close()
// returns XML


}
