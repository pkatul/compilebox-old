package main

import (
	"fmt"
	"testbox"
)

func main() {
	lang := "Python"
	code := "import sys\nprint(len(sys.stdin.read()))\n"
	stdin := "12345\n1234567891011\n"
	fmt.Println(testbox.Test(lang, code, stdin, "11\n11\n"))
}