package main

import "fmt"

const greeting = "hello, wolrd!"

func main() {
	printMsg(greeting)
}

func printMsg(msg string) {
	fmt.Println(msg)
}
