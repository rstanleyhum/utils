package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/rstanleyhum/utils/jupyter"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	for {
		var m jupyter.SourceOnlyNotebook
		if err := dec.Decode(&m); err != nil {
			if err == io.EOF {
				return
			}
			log.Println("Decode Error")
			log.Println(err)
			return
		}

		if err := enc.Encode(&m); err != nil {
			log.Println("Encode Error")
			log.Println(err)
		}
	}
}
