package main
import "fmt"

//Main Function
func main() {
  var MyInt int
  MyInt = 2
  fmt.Println("Hello World!  I am a Go Program!")
  fmt.Println("Also - Hello World; I am another line of code!")
  //String Interpolation
  fmt.Println("The value of my variable is ", MyInt)
}


//TODO: Define a function and call it from main:
//Define another function
func addition_func(x) {
  return x+1
}
