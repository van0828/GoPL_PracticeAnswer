package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		for s, n := range counts {
			if n > 1 {
				fmt.Printf("%s\t%d\n", s, n)
			}
		}
	} else {
		for _, arg := range files {
			count := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, count)
			f.Close()
			for s, n := range count {
				if n > 1 {
					fmt.Printf("%s\t%d\t%s\n", s, n, arg)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
