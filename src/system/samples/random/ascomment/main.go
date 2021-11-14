package main

// #include<stdlib.h>
// #include<stdio.h>
// void Cmain() {
//   srandom(1);
//   printf("random int is %li\n", random());
// }
import "C"

func main() {
	C.Cmain()
}
