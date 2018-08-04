// https://www.youtube.com/watch?v=OwakEEY_DFw
package main
import "fmt"

func main(){
  x:= 15
  a := &x //points to x/memory address (where the value for X is being stored)
  fmt.Println(a)
  fmt.Println(*a)
  *a = 5 // changes x to 5
  fmt.Println(x)
  // 


}
