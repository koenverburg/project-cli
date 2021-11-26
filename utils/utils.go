package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyString(in interface{}) {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
