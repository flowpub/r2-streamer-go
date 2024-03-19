package cli

import (
	"encoding/json"
	"fmt"

	"github.com/flowpub/r2-streamer-go/parser"
)

func RunParse(filepath string) {
	publication, err := parser.Parse(string(filepath))
	if err != nil {
		panic(err)
	}

	jsonStr, err := json.Marshal(publication)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonStr))
}
