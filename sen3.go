// types

package main
import ("fmt")

func add(x float64, y float64) float64{
  // need to add the types of variables to expect
  // a simple addition command
  return x+y
}


func main (){
  // defining a variable:
  var num1 float64 = 5.6
  // this is similar to JS?
// float64 is the set of all IEEE-754 64-bit floating-point numbers.
// floatNUMBER is the number of sig figs(?)
  var num2 float64 = 9.5
  fmt.Println(add(num1, num2))
}

// making the main shorter...
// constants...
// const x int = 5
func add(x, y float32) float32{
  // need to add the types of variables to expect
  // a simple addition command
  return x+y
}

func main (){
  num1, num2 float32 := 5.6, 9.5
    // if inside a function, don't need to give the type, but the type can't change!
    // https://tour.golang.org/basics/10
    // := is a shorthand for assigning variables


  fmt.Println(add(num1, num2))
}

func multiple(a, b string) (string, string){
  return a,b
  // need to specify every return type, even if the same
}

func main (){
  num1, num2 := 5.6, 9.5
  w1, w2 = "Hey", "there"

  fmt.Println(multiple(num1, num2))

/*  multiple line comments
       are like this */
  var a int = 62
  var b float64 = float64(a)


}
