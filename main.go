package main

import (
	"fmt"
	"syscall"
)

func main() {
	syscall.Mprotect()
}

func a() {

	fmt.Println("hello world")

}
