package main

import (
	"fmt"
	"os"
	"strings"
)

func dup3() {

	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for key,value := range counts {
		if value >= 2 {
			fmt.Printf("%d\t%s\n",value,key)
		}
	}
}
