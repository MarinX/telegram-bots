package main

import (
	"math/rand"
	"time"
)

const maxRoomNameSize = 6

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

//GenerateRandomString returns random string
func GenerateRandomString() string {
	b := make([]rune, maxRoomNameSize)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
