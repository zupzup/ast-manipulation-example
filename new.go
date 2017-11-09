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

// TODO: document exported function
func AnExportedFunction() {

}

// TODO: document exported function
func DoStuff() {

}

// DoOtherStuff does other stuff
func DoOtherStuff() {

}
