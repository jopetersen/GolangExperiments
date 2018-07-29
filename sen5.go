package main
//https://www.youtube.com/watch?v=-QmdZ7821wA

import ("fmt"
        "net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
  // w is a writer, * is reading through
  fmt.Fprintf(w, "Woah, Go is neat!")
     // fprint is formatting whatever specified, based on what you specified
     // will output whatever you ask it to
// to view: http://localhost:8000/

}

func about_handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Expert About Page")

}
func main(){
  // creating a handler
  // taking a URL/path, figure out the type of function that corresponds
  http.HandleFunc("/", index_handler)
      // takes a path (/ is index)
  // to add a new page...
  http.HandleFunc("/about/", about_handler)

  http.ListenAndServe(":8000", nil)
    // creates our server at a local host with port 8000

}
