package utils

import (
	"math/rand"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyz0123456789"

func CreateId() string {
	var id string = ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 25; i++ {
		id += string(chars[rand.Intn(len(chars))])
	}
	return id
}
