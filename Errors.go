// This file should contain all error-handling functions

package main

import "fmt"

func DialFailed(err error) {
	// TODO: should we continue?
	fmt.Println(err.Error())
}
