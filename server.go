package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
}

type village map[*citizen]bool

type citizen struct {
	name             string
	connection       *connection
	isVillageCreator bool
}

type conncection struct {
}
