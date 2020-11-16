package cli

import (
	"fmt"
)

//PrintFormat standard formatting
func PrintFormat(args ...string) {

	for _, a := range args {
		fmt.Printf("%-30v ", a)
	}
	fmt.Print("\n")
}
