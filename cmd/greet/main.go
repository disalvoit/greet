// Command greet prints a greeting for the name given on the command line.
package main

import (
	"fmt"
	"os"

	"github.com/disalvoit/greet"
)

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greet.Hello(name))
}
