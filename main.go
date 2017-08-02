package main

import (
	"github.com/practice/DataFetcher/ifsc"
)

func main() {
	ifscInit()
}

func ifscInit() {
	ifsc := &ifsc.Ifsc{}
	ifsc.Init()
}
