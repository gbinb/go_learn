package sample

import (
	"fmt"
)

func ModifyStr(str string) {
	str = "hello world"
}

func ReferenceSamp() {
	str := "go language"
	ModifyStr(str)
	fmt.Printf("%s\n", str)
}
