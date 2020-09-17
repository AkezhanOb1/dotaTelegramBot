package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	hour := a / 30
	a = a % 30
	fmt.Printf("It is %d hours %d minutes \n", hour, a*2)
}
