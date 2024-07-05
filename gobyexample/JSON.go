package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type JSONObject struct {
		Page    int      `json:"page"`
		Strings []string `json:"fruits"`
	}

	p := &JSONObject{
		Page:    1,
		Strings: []string{"apple", "babana", "orange"},
	}
	p1, _ := json.Marshal(p)
	fmt.Println(string(p1))

	str := `{"page":1,"fruits":["apple","babana","orange"]}`
	res := JSONObject{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("%T: %#v", res, res)
}
