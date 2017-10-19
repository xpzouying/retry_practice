package main

import (
	"errors"
	"fmt"
)

var (
	doTimes int
)

func main() {
	if err := retry(3, 1, doFunc); err != nil {
		fmt.Printf("retry doFunc error: %v\n", err)
	}

}

func doFunc() error {
	doTimes++

	if doTimes <= 2 {
		fmt.Println("do fail")
		return errors.New("doFunc fail")
	}

	fmt.Printf("this is not first do, success")
	return nil
}
