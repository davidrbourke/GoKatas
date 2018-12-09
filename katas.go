package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	discoverpin "github.com/davidrbourke/katas/discoverPin"
)

func main() {
	fmt.Println("*** Katas ***")
	keys, err := ioutil.ReadFile("discoverPin/keylogger.txt")
	check(err)

	lines := strings.Split(string(keys), "\r\n")

	discoverpin.GetPin(lines)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
