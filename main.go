package main

import (
	"fmt"

	"github.com/practice/DataFetcher/ifsc"
)

func main() {
	fmt.Println("Hello World!")
	ifscInit()
}

func ifscInit() {
	ifsc := &ifsc.Ifsc{}
	ifsc.Init()
}
