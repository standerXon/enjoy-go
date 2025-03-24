package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLine(f *os.File, counts map[string]int) {
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		counts[scan.Text()]++
	}
}

func dup2() {
	// to stone result
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, fileName := range files {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLine(f, counts)
			err1 := f.Close()
			if err1 != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err1)
			}
		}
	}

	for key, value := range counts {
		if value >= 2 {
			fmt.Printf("%s:%d\n", key, value)
		}
	}

}
