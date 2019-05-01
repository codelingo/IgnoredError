//Package main is an example package
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := PrintDivision(1, 2)
	if err != nil {
		panic(err.Error())
	}

	PrintDivision(2, 0)
}

func PrintDivision(a, b int) (error) {
	if b == 0 {
		return errors.New("divide by zero error")
	}
	fmt.Println(a/b)
	return nil
}

