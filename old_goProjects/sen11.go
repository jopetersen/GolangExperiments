package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")



// defining XML structure
type SitemapIndex struct{
  Locations []Location `xml:"sitemap"`
      //REMEMBER TO CAPITALIZE LOCATIONS
      // [] is an array

}

//[5]type == array, of a fixed size
//[]type == slice

type Location struct{
  Loc string `xml:"loc"`
}

func (L Location) String() string{
  return fmt.Sprintf(l.Loc)
}

func main() {
  // want to get information from the internet
  // convert information to a string, then parsing/formatting
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
      // underscores are used anytime that you define a variable that
      // you don't intend to use (initializing a variable)
  bytes, _ := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

    // converting to a string
  var s SitemapIndex
  xml.Unmarshal(bytes, &s)
  fmt.Println(s.Locations)
// returns XML


}
