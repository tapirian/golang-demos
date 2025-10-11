package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name,omitempty"`
	PName string `json:"pname,omitempty"`
}

func main() {
	zhangsan := &Person{Name: "zhangsan", PName: "zhangsansan"}
	j, _ := json.Marshal(zhangsan)
	j1, _ := json.MarshalIndent(zhangsan, "", "")

	var d = &Person{}
	json.Unmarshal(j, d)

	var d1 = &Person{}
	json.Unmarshal(j1, d1)

	fmt.Printf("%s\n", j)
	fmt.Printf("%s\n", j1)
	fmt.Printf("%s\n", d.Name)
	fmt.Printf("%s\n", d1.Name)
}
