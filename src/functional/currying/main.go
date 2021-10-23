package main

import "os"

func LogToFile(filename string) func(message string) {
	f, _ := os.Create(filename)
	//defer f.Close() This cannot work here
	return func(message string) {
		_, _ = f.WriteString(message + "\n")
	}
}

func main() {
	LogToFile("log.txt")("Log this statement")

	Logger := LogToFile("log.txt") // partial application
	Logger("Log this statement")
}
