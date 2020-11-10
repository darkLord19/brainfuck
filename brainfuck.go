package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type cpu struct {
	mem           [30000]byte
	pc            uint32
	sp            uint32
	in            *bufio.Reader
	out           *bufio.Writer
	matchKeyStart map[uint32]uint32
	matchKeyEnd   map[uint32]uint32
}

func (c *cpu) getchar() {
	input, _ := c.in.ReadString('\n')
	c.mem[c.sp] = []byte(input)[0]
}

func (c *cpu) putchar(ch byte) {
	c.out.Write([]byte{ch})
	c.out.Flush()
}

func (c *cpu) setMatchingStartEndPairs(prog []byte) error {
	arr := []uint32{}
	for i := range prog {
		if prog[i] == '[' {
			arr = append(arr, uint32(i))
		} else if prog[i] == ']' {
			c.matchKeyStart[arr[len(arr)-1]] = uint32(i)
			c.matchKeyEnd[uint32(i)] = arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		}
	}
	if len(arr) != 0 {
		return fmt.Errorf("missing pairs of [ ]")
	}
	return nil
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
			verbose("[", c.mem[c.sp], c.sp)
			if c.mem[c.sp] == 0 {
				c.pc = c.matchKeyStart[c.pc]
				verbose("[", c.mem[c.sp], c.sp)
			}
		case ']':
			verbose("]", c.mem[c.sp], c.sp)
			if c.mem[c.sp] > 0 {
				c.pc = c.matchKeyEnd[c.pc]
				verbose("]", c.mem[c.sp], c.sp)
			}
		case ',':
			c.getchar()
			verbose(",", c.mem[c.sp], c.sp)
		case '.':
			c.putchar(c.mem[c.sp])
			verbose(".", c.mem[c.sp], c.sp)
		}
		verbose(string(prog[c.pc]), c.mem[:10], c.pc, c.sp)
	}
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
		c := cpu{in: bufio.NewReader(os.Stdin), out: bufio.NewWriter(os.Stdout),
			matchKeyStart: make(map[uint32]uint32), matchKeyEnd: make(map[uint32]uint32)}
		err = c.setMatchingStartEndPairs(prog)
		if err != nil {
			panic(err)
		}
		c.run(prog)
	}
}
