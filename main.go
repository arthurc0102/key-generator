package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	lowercase   = "abcdefghijklmnopqrstuvwxyz"
	uppercase   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "0123456789"
	punctuation = "!#$%&()*+,-./:;<=>?@[]^_`{|}~"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var (
		l, u, d, p bool
		repeat     bool
		length     int
		target     string
	)

	flag.BoolVar(&l, "l", true, "Use lowercase char in target")
	flag.BoolVar(&u, "u", true, "Use uppercase char in target")
	flag.BoolVar(&d, "d", true, "Use digits char in target")
	flag.BoolVar(&p, "p", true, "Use punctuation in target")
	flag.BoolVar(&repeat, "r", true, "Char can be used for many times")
	flag.IntVar(&length, "len", 50, "Length for key")
	flag.Parse()

	if length < 1 {
		log.Fatalln("Length should be bigger than 0")
	}

	config := []bool{l, u, d, p}
	chars := []string{lowercase, uppercase, digits, punctuation}

	if len(config) != len(chars) {
		panic("Len of 'config' and 'chars' should be the same")
	}

	for i := range config {
		if !config[i] {
			continue
		}

		target += chars[i]
	}

	targetLen := len(target)

	if !repeat && int(length) > targetLen {
		fmt.Println("Your key length is smaller then target, you can:")
		fmt.Println("  - Make your key length shorter")
		fmt.Println("  - Make your target longer")
		fmt.Println("  - Allow repeat")
		os.Exit(1)
	}

	result := make([]byte, length)

	if repeat {
		for i := range result {
			result[i] = target[rand.Intn(targetLen)]
		}
	} else {
		shuffle := rand.Perm(targetLen)
		for i := range result {
			result[i] = target[shuffle[i]]
		}
	}

	fmt.Println(string(result))
}
