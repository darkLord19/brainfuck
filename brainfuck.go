package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	buffSize = 100
)

var (
	mem    = [30000]byte{}
	ptr    = 0
	reader = bufio.NewReader(os.Stdin)
)

func interpret(code []byte) {
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '[':
			if mem[ptr] == 0 {
				for code[i] != ']' && i < len(code) {
					i++
				}
				i++
			}
		case ']':
			if mem[ptr] != 0 {
				for code[i] != '[' && i > -1 {
					i--
				}
				i--
			}
		case ',':
			input, _ := reader.ReadString('\n')
			mem[ptr] = []byte(input)[0]
		case '.':
			fmt.Print(string(mem[ptr]))
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		code, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		interpret(code)
	}
}
