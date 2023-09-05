package core_test

import (
	"flag"
	"fmt"
	"testing"
)

var wordPtr = flag.String("word", "foo", "a string")
var numbPtr = flag.Int("numb", 42, "an int")
var forkPtr = flag.Bool("fork", false, "a bool")
var svar = flag.StringVar(&svar, "svar", "bar", "a string var")

func TestFlag(t *testing.T) {

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numbPtr:", *numbPtr)
	fmt.Println("forkPtr:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
