package main
// packages ALWAYS need to be imported at the top
import ("fmt"
        "math/rand")
// math is standard
//func foo(){
//  fmt.Println("Foo")
  // this will NOT run automatically

//}



func main(){
  fmt.Println("A number from 1-100 is..", rand.Intn(100))
// All of the functions from packages need to be capitalized, else they'd be systems facing
// in Go, you don't need to call functions to get them to run
// main function will ALWAYS run (like init in class)
}

// if you want to know more about something, enter...
//     godoc QUERY
// into the CLI
