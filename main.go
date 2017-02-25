package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ledongthuc/note-down/core"
)

func main() {
	input := "# Header\ncontent\n#HEader\n*a*"
	inputInBytes := []byte(input)
	output := core.ParseRequest(inputInBytes)
	fmt.Println(string(output))

	fmt.Println("********\n")

	var err error
	output, err = ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}
	output = core.ParseRequest(output)

	fmt.Println(string(output))
}
