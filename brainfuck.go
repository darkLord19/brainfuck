package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

type cpu struct {
	mem [30000]byte
	pc  uint32
	in  io.Reader
	out io.Writer
}

func (c *cpu) getchar() {
	var input []byte
	c.in.Read(input)
	c.mem[c.pc] = []byte(input)[0]
}

func (c *cpu) putchar(ch byte) {
	c.out.Write([]byte{ch})
}

func getMatchingLoopStartPos() uint32 {
	return 0
}

func getMatchingLoopEndPos() uint32 {
	return 0
}

func (c *cpu) run(prog []byte) {
	for i := 0; i < len(prog); i++ {
		switch prog[i] {
		case '>':
			c.pc++
		case '<':
			c.pc--
		case '+':
			c.mem[c.pc]++
		case '-':
			c.mem[c.pc]--
		case '[':
			if c.mem[c.pc] == 0 {
				pos := getMatchingLoopEndPos()
				c.pc = pos + 1
				i = int(c.pc)
			}
		case ']':
			if c.mem[c.pc] > 0 {
				pos := getMatchingLoopStartPos()
				c.pc = pos + 1
				i = int(c.pc)
			}
		case ',':
			c.getchar()
		case '.':
			c.putchar(c.mem[c.pc])
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		prog, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		c := cpu{in: bufio.NewReader(os.Stdin), out: bufio.NewWriter(os.Stdout)}
		c.run(prog)
	}
}
