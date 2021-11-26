package main

import "fmt"
import "os"
import "vocab"

func main() {
	for _, word := range os.Args[1:] {
		otherWord, error := vocab.Hypernym(word)
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		} else {
			fmt.Printf("A %v is an %v.\n", word, otherWord)
		}
	}
}
