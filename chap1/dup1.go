package main

import("os";"fmt";"bufio")

func dup1() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for key,value := range counts {
		if value > 2 {
			fmt.Printf("%s:%d\n",key,value)
		}
	}
}
