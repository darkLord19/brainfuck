package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type cpu struct {
	mem [30000]byte
	pc  uint32
	in  *bufio.Reader
	out *bufio.Writer
}

func (c *cpu) getchar() {
	input, _ := c.in.ReadString('\n')
	c.mem[c.pc] = []byte(input)[0]
}

func (c *cpu) putchar(ch byte) {
	c.out.Write([]byte{ch})
}

func (c *cpu) getMatchingLoopStartPos() uint32 {
	var start, end byte
	start = '['
	end = ']'
	cnt := 1
	verbose("matchigStart", c.pc)
	for i := c.pc; i > 0; i-- {
		if c.mem[i] == start {
			cnt--
		} else if c.mem[i] == end {
			cnt++
		}
		if cnt == 0 {
			return i
		}
	}
	return 0
}

func (c *cpu) getMatchingLoopEndPos() uint32 {
	var start, end byte
	start = '['
	end = ']'
	cnt := 1
	verbose("matchigEnd", c.pc)
	for i := c.pc; i >= 0; i-- {
		if c.mem[i] == start {
			cnt++
		} else if c.mem[i] == end {
			cnt--
		}
		if cnt == 0 {
			return i
		}
	}
	return 0
}

func verbose(op string, args ...interface{}) {
	if *vFlag {
		fmt.Println(op, args)
	}
}

func (c *cpu) run(prog []byte) {
	for i := 0; i < len(prog); i++ {
		switch prog[i] {
		case '>':
			c.pc++
			verbose(">", c.pc)
		case '<':
			c.pc--
			verbose("<", c.pc)
		case '+':
			c.mem[c.pc]++
			verbose("+", c.mem[c.pc], c.pc)
		case '-':
			c.mem[c.pc]--
			verbose("-", c.mem[c.pc], c.pc)
		case '[':
			if c.mem[c.pc] == 0 {
				pos := c.getMatchingLoopEndPos()
				c.pc = pos + 1
				i = int(c.pc)
				verbose("[", c.mem[c.pc], pos, i)
			}
		case ']':
			if c.mem[c.pc] > 0 {
				pos := c.getMatchingLoopStartPos()
				c.pc = pos + 1
				i = int(c.pc)
				verbose("]", c.mem[c.pc], c.pc, pos, i)
			}
		case ',':
			c.getchar()
			verbose(",", c.mem[c.pc], c.pc)
		case '.':
			c.putchar(c.mem[c.pc])
			verbose(".", c.mem[c.pc], c.pc)
		}
	}
	c.out.Flush()
}

var (
	vFlag *bool
)

func main() {
	vFlag = flag.Bool("v", false, "verbose output")
	flag.Parse()
	if len(os.Args) > 1 {
		prog, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
		c := cpu{in: bufio.NewReader(os.Stdin), out: bufio.NewWriter(os.Stdout)}
		c.run(prog)
	}
}
