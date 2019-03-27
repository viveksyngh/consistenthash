package main

import (
	"fmt"

	"github.com/viveksyngh/consistenthash/consistenthash"
)

func main() {
	chash := consistenthash.New(10, nil)
	chash.Add("S1", "S2", "S3")
	keys := []string{
		"vivek@gmail.com",
		"viveksyngh@gmail.com",
		"john@gmail.com",
		"mark@gmail.com",
		"adam@gmail.com",
		"steve@gmail.com",
		"smith@gmail.com",
	}

	for _, key := range keys {
		fmt.Printf("%s : %s \n", chash.Get(key), key)
	}

	fmt.Println("Removing server S1")
	chash.Remove("S1")
	for _, key := range keys {
		fmt.Printf("%s : %s \n", chash.Get(key), key)
	}
}
