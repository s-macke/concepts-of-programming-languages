package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var attempts []int

func main() {
	//Generate the random integer
	secret := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter an integer between 1 and 100: ")

	for {
		entry, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		line := strings.Fields(entry)[0]
		num, err := strconv.Atoi(line)

		attempts = append(attempts, num)

		if err != nil {
			fmt.Println("Enter a number between 1 and 100.")
		} else if num < 1 || num > 100 {
			fmt.Println("Enter a number between 1 and 100.")
		} else {
			if secret-5 <= num && num < secret {
				fmt.Println("You are extremely close from the left.")
			} else if secret-10 <= num && num < secret-5 {
				fmt.Println("You are very close from the left.")
			} else if secret-20 <= num && num < secret-10 {
				fmt.Println("You are somewhat close from the left.")
			} else if num < secret-20 {
				fmt.Println("You are way off from the left.")
			} else if secret+5 >= num && num > secret {
				fmt.Println("You are extremely close from the right.")
			} else if secret+10 >= num && num > secret+5 {
				fmt.Println("You are very close from the right.")
			} else if secret+20 >= num && num > secret+10 {
				fmt.Println("You are somewhat close from the right.")
			} else if num > secret+20 {
				fmt.Println("You are way off from the right.")
			} else if secret == num {
				message := fmt.Sprintf("Congrats!!!, You got the number after %d attempts!!!\n", len(attempts))
				fmt.Println(message)
				break
			}
		}

		if secret != num {
			fmt.Println("Try again: ")
		}
	}

}
