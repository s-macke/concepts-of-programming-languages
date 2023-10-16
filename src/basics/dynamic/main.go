package main

import "fmt"

// START OMIT
func main() {
    var dynamicVar interface{}

    dynamicVar = 55
    fmt.Println(dynamicVar)  // Outputs: 55

    dynamicVar = "Hello, Golang!"
    fmt.Println(dynamicVar)  // Outputs: Hello, Golang!
}



// END OMIT
