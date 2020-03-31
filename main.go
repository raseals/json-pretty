package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file-name>\n", os.Args[0])
		return
	}

	fileContents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file: %s\n", os.Args[1])
		return
	}

	m := map[string]interface{}{}
	err = json.Unmarshal(fileContents, &m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid json input: %s\n", os.Args[1])
		return
	}

	b, _ := json.MarshalIndent(m, "", "    ")
	fmt.Println(string(b))
}
