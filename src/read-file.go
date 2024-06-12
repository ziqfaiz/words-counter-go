package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(name string) {
	fmt.Println("Reading file: ", name)
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	count := 0
	countMap := make(map[string]int)
	for scanner.Scan() { // internally, it advances token based on sperator
		//fmt.Println(scanner.Text()) // token in unicode-char
		//fmt.Println(scanner.Bytes()) // token in bytes
		count++
	}
	fmt.Println("Total words in file: ", count)
}
