package main

import "fmt"

func main() {
    var num int = 42
    ptr := &num
    
    fmt.Println("Value of num:", num)           
    // Expected Output: Value of num: 42
    
    fmt.Println("Address of num:", &num)        
    // Expected Output: Address of num: 0xSOME_MEMORY_ADDRESS (this will vary every run)
    
    fmt.Println("Value via pointer:", *ptr)     
    // Expected Output: Value via pointer: 42
    
    fmt.Println("Address stored in ptr:", ptr)  
    // Expected Output: Address stored in ptr: 0xSOME_MEMORY_ADDRESS (same as address of num)
}
