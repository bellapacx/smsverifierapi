package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("testescape/serviceAccount.json") // your downloaded file
	if err != nil {
		panic(err)
	}

	// Replace actual newlines with \n and escape backslashes
	jsonStr := strings.ReplaceAll(string(data), "\n", "\\n")

	fmt.Println(jsonStr) // paste this output into Railway
}
