package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const target = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*(-_=+)"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	keyLen := 50

	if len(os.Args) > 1 {
		args := os.Args

		if k, err := strconv.Atoi(args[1]); err != nil {
			fmt.Printf("Key len should be int, got '%s'\n", args[1])
			os.Exit(1)
		} else {
			keyLen = k
		}
	}

	key := make([]byte, keyLen)

	for i := range key {
		key[i] = target[rand.Intn(len(target))]
	}

	fmt.Println(string(key))
}
