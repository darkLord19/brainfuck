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
	sp  uint32
	in  *bufio.Reader
	out *bufio.Writer
}

func (c *cpu) getchar() {
	input, _ := c.in.ReadString('\n')
	c.mem[c.sp] = []byte(input)[0]
}

func (c *cpu) putchar(ch byte) {
	c.out.Write([]byte{ch})
}

func (c *cpu) getMatchingLoopStartPos(prog []byte) uint32 {
	var start, end byte
	start = '['
	end = ']'
	cnt := 1
	verbose("matchigStart", c.pc, c.sp)
	for i := c.pc - 1; i > 0; i-- {
		if prog[i] == start {
			cnt--
		} else if prog[i] == end {
			cnt++
		}
		if cnt == 0 {
			return i
		}
	}
	return 0
}

func (c *cpu) getMatchingLoopEnspos(prog []byte) uint32 {
	var start, end byte
	start = '['
	end = ']'
	cnt := 1
	verbose("matchigEnd", c.pc)
	for i := c.pc - 1; i >= 0; i-- {
		if prog[i] == start {
			cnt++
		} else if prog[i] == end {
			cnt--
		}
		if cnt == 0 {
			return i
		}
	}
	return 0
}

func verbose(op string, args ...interface{}) {
	if vFlag {
		fmt.Println(op, args)
	}
}

func (c *cpu) run(prog []byte) {
	for c.pc = 0; c.pc < uint32(len(prog)); c.pc++ {
		switch prog[c.pc] {
		case '>':
			c.sp++
			verbose(">", c.sp)
		case '<':
			c.sp--
			verbose("<", c.sp)
		case '+':
			c.mem[c.sp]++
			verbose("+", c.mem[c.sp], c.sp)
		case '-':
			c.mem[c.sp]--
			verbose("-", c.mem[c.sp], c.sp)
		case '[':
			if c.mem[c.sp] == 0 {
				pos := c.getMatchingLoopEnspos(prog[c.pc:])
				c.pc = pos + 1
				verbose("[", c.mem[c.sp], pos)
			}
			verbose("[", c.mem[c.sp], c.sp)
		case ']':
			if c.mem[c.sp] > 0 {
				pos := c.getMatchingLoopStartPos(prog[:c.pc])
				c.pc = pos + 1
				verbose("]", c.mem[c.sp], c.sp, pos)
			}
			verbose("]", c.mem[c.sp], c.sp)
		case ',':
			c.getchar()
			verbose(",", c.mem[c.sp], c.sp)
		case '.':
			c.putchar(c.mem[c.sp])
			verbose(".", c.mem[c.sp], c.sp)
		}
		// fmt.Println(string(prog[c.pc]), c.mem[:10], c.pc, c.sp)
	}
	c.out.Flush()
}

var (
	vFlag bool
)

func main() {
	var fname string
	flag.BoolVar(&vFlag, "v", false, "verbose output")
	flag.StringVar(&fname, "f", "", "filename containing brainfuck program")
	flag.Parse()
	if fname != "" {
		prog, err := ioutil.ReadFile(fname)
		if err != nil {
			panic(err)
		}
		c := cpu{in: bufio.NewReader(os.Stdin), out: bufio.NewWriter(os.Stdout)}
		c.run(prog)
	}
}
