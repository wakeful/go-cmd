package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	code := 0

	for _, arg := range os.Args[1:] {
		fi, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			code = 1
		}
		fmt.Println(string(fi))
	}

	os.Exit(code)

}
