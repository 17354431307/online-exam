package core_test

import (
	"backend/core"
	"backend/global"
	"flag"
	"fmt"
	"github.com/magiconair/properties/assert"
	"os"
	"strings"
	"testing"
)

var wordPtr = flag.String("word", "foo", "a string")
var numbPtr = flag.Int("numb", 42, "an int")
var forkPtr = flag.Bool("fork", false, "a bool")

func TestFlag(t *testing.T) {

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numbPtr:", *numbPtr)
	fmt.Println("forkPtr:", *forkPtr)
	fmt.Println("tail:", flag.Args())
}

func TestGetenvAndSetEnv(t *testing.T) {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

func TestInitViper(t *testing.T) {
	core.InitViper("../etc/config.yaml")
	assert.Equal(t, 8081, global.OE_CONFIG.App.Port)
}
