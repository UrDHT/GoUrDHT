package main

import (
	"fmt"
)

func main() {
	target := "http://envy.blamestross.com:8000/"
	nw := new(Networking)
	reply := nw.GetIP(target)
	fmt.Println(reply)
}
