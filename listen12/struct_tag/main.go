package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Sex      string `json:"sex"`
	Score    float32
}

func main() {
	user := &User{
		Username: "user01",
		Sex:      "男",
		Score:    92.2,
	}
	data, _ := json.Marshal(user)
	fmt.Printf("json str:%s\n", string(data))
}
