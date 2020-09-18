package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Group string `json:"group"`
}

func main() {

	u1 := user{Name: "Kaire", Age: 12, Group: "Mamsoum"}

	fmt.Println(u1)

	res, err := json.Marshal(u1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
	jsonData := string(res)

	fmt.Println(jsonData)

	u1Again := user{}

	// err := json.Unmarshal(res, u1_again)
	err = json.Unmarshal([]byte(jsonData), &u1Again)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u1Again)

}
