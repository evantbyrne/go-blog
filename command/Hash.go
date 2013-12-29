package main

import "fmt"
import "os"

import "../util"

func main() {
	for _, value := range os.Args[1:] {
		fmt.Println(util.Hash(value))
	}
}