package main

import (
	"fmt"

	"github.com/viveksyngh/consistenthash/consistenthash"
)

func main() {
	chash := consistenthash.New(10, nil)
	chash.Add("S1", "S2", "S3")
	fmt.Println(chash.HashRing, chash.HashMap)
	fmt.Println(chash.Get("viveks@gmail.com"))
	fmt.Println(chash.Get("viveksyngh@gmail.com"))
}
