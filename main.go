package main

import (
	"github.com/Practice/DataFetcher/ifsc"
)

func main() {
	ifscInit()
}

func ifscInit() {
	ifsc := &ifsc.Ifsc{}
	ifsc.Init()
}
