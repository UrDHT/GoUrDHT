// This file should contain all error-handling functions

package GoUrDHT

import "fmt"

func DialFailed(err error) {
	// TODO: should we continue?
	fmt.Println(err.Error())
}
