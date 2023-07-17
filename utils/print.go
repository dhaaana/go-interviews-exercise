package utils

import "fmt"

func PrintResult(values ...interface{}) {
	for _, value := range values {
		fmt.Printf("%v\n", value)
	}
}
