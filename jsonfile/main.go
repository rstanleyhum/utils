package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type container struct {
	File string `json:"file"`
}

func main() {
	input := os.Stdin
	output := os.Stdout

	data, err := ioutil.ReadAll(input)
	if err != nil {
		log.Println(err)
	}

	content := container{string(data)}

	enc := json.NewEncoder(output)
	enc.Encode(content)

}
