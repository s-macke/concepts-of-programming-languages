package main

import (
	"encoding/json"
	"strconv"
)

type Test struct {
	Stringg string         `json:"string"`
	Integer int            `json:"integer"`
	Array   []string       `json:"array"`
	Mapp    map[string]int `json:"map"`
}

func main() {
	for i := 0; i < 10000000; i++ {
		str := `{"string": "string", "integer": ` + strconv.Itoa(i) + `, "array": ["array", "array", "array"], "map":{"map":1,"map":2,"map":3}}`
		var t Test
		json.Unmarshal([]byte(str), &t)
	}
}
