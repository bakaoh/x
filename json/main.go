package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSONableSlice ...
type JSONableSlice []uint8

// MarshalJSON ...
func (u JSONableSlice) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = strings.Join(strings.Fields(fmt.Sprintf("%d", u)), ",")
	}
	return []byte(result), nil
}

// Object ...
type Object struct {
	I JSONableSlice `json:"i,omitempty"`
}

func main() {
	o := &Object{I: JSONableSlice{1, 2, 3}}
	s, _ := json.Marshal(o)
	fmt.Println(string(s))
}
