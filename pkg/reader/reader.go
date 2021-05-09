package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadStdin() {
	r := bufio.NewReader(os.Stdin)
	for {
		line, isPrefix, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(line))
		if !isPrefix {
			fmt.Println()
		}
	}
}
