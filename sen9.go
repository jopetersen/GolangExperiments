package main

import ("fmt"
"net/http")

func index_handler (w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, `
    <h1>Hey there</h1>
    <p> test re </p>
    <p> working?</p>
    `)


  fmt.Fprintf(w, "<h1>Hey there</h1>")
  fmt.Fprintf(w, "<p>Go is fast!</p>")
  fmt.Fprintf(w, "<p>and simple!</p>")
  fmt.Fprintf(w, "<p>You %s even %s</p>", "can ", "<strong>variables</strong")
// can pass HTML tags by writing them in

}


func main() {
  http.HandleFunc("/", index_handler)
  http.ListenAndServe(":8000", nil)


}
