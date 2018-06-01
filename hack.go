package main
import "fmt"

//Main Function
func main() {
  //Declare and define variables
  var MyInt int
  MyInt = 2
  fmt.Println("Hello World!  I am a Go Program!")
  fmt.Println("=========================================")
  fmt.Println("Also - Hello World; I am another line of code!")
  //String Interpolation
  fmt.Println("The value of my variable is ", MyInt)

  //Declare and define variables in the same line:
  var MyIntInc int = addition_func(MyInt)
  fmt.Println("The incremented value of my variable is", MyIntInc)
}


//TODO: Define a function and call it from main:
//Define another function
func addition_func(x int) int {
  return x+1
}
