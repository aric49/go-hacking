package main
import "fmt"
import "os"
import "flag"

//Main Function
func main() {
  //Declare and define variables
  var MyInt int
  MyInt = 2
  fmt.Println("Hello World!  I am a Go Program!")
  fmt.Println("=========================================")
  fmt.Println("Also - Hello World; I am another line of code!")

  //Test for Environment variable
  var MyEnvironmentVar string = os.Getenv("SETTING1")

  //If-else
  if MyEnvironmentVar == "" {
    //Don't use := here since we are assinging a new value. Not initializing.
    MyEnvironmentVar = "TestValue"
    fmt.Println("The Environment Variable is not defined. Using default value: ", MyEnvironmentVar)

  } else {
    fmt.Println("The Environment Variable is defined. The value is: ", MyEnvironmentVar)
  }

  //CLI Flags
  // When defining CLI flags, variables must be initialized first and referenced in the flag. 
  var cliInt int
  var cliStr string
  flag.IntVar(&cliInt, "number", 0, "A Random Integer")
  flag.StringVar(&cliStr, "text", "default string", "A string value to use at the CLI")
  flag.Parse()


  // fmt.Println("Using a cli arg: ", cliInt)
  fmt.Println("The CLI option for '--number' is: ", cliInt, " - ", evalInteger(cliInt))
  fmt.Println("The CLI option for '--text' is: ", cliStr)

  //String Interpolation
  fmt.Println("The SETTING1 configuration option is: ", MyEnvironmentVar)

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

func evalInteger(integer_value int) string {
  if integer_value > 10 {
    return "The provided value is greater than 10"
  } else if integer_value < 10 {
    return "The provided value is less than 10!"
  }
  return ""
}
