package main

import "fmt"

func main() {
	num := 42   // implicit type "int" with value 42
	ptr := &num // implicit type of "pointer to type int points to num"
	//var num int = 42 // explicit type definition
	//var ptr *int = &num // explicit type definition

	fmt.Println("Value of num:", num)
	// Expected Output: Value of num: 42

	fmt.Println("Address of num:", &num)
	// Expected Output: Address of num: 0xSOME_MEMORY_ADDRESS (this will vary every run)

	fmt.Println("Address stored in ptr:", ptr)
	// Expected Output: Address stored in ptr: 0xSOME_MEMORY_ADDRESS (same as address of num)

	fmt.Println("Value via pointer:", *ptr)
	// Expected Output: Value via pointer: 42

	*ptr = 43
	fmt.Println("Value of num:", num)
	// Expected output: New value in num = 43
}
