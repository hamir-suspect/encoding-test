package main

import (
    "fmt"
    "time"
    "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func main() {
    rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 100; i++ {
		fmt.Print(randSeq(30)+"\n")
	}
}