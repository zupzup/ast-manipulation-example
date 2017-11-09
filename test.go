package main

import (
	"fmt"
)

func main() {
	fmt.Println("testprogram")
	DoStuff()
}

func unexportedFunction() {

}

// Whatever does other stuff
func Whatever() {

}

func AnExportedFunction() {

}

func DoStuff() {

}

// DoOtherStuff does other stuff
func DoOtherStuff() {

}
