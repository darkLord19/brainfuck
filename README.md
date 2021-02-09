# Brainfuck interpreter written in Go is Awesome

Brainfuck is an esoteric programming language created in 1993 by Urban Müller.

The language consists of eight commands, listed below. A brainfuck program is a sequence of these commands, possibly interspersed with other characters (which are ignored). The commands are executed sequentially, with some exceptions: an instruction pointer begins at the first command, and each command it points to is executed, after which it normally moves forward to the next command. The program terminates when the instruction pointer moves past the last command.

The brainfuck language uses a simple machine model consisting of the program and instruction pointer, as well as an array of at least 30,000 byte cells initialized to zero; a movable data pointer (initialized to point to the leftmost byte of the array); and two streams of bytes for input and output (most often connected to a keyboard and a monitor respectively, and using the ASCII character encoding).

The eight language commands each consist of a single character:

|   Character   |   Meaning   |
|---------------|-------------|
|       >       | increment the data pointer (to point to the next cell to the right).|
|       <       | decrement the data pointer (to point to the next cell to the left).|
|       + 	    | increment (increase by one) the byte at the data pointer.|
|       - 	    | decrement (decrease by one) the byte at the data pointer.|
|       . 	    | output the byte at the data pointer.|
|       , 	    | accept one byte of input, storing its value in the byte at the data pointer.|
|       [ 	    | if the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching ] command.|
|       ] 	    | if the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching [ command.|



# Run Examples

To run examples, run following command
```bash
go run brainfuck.go examples/<filename>
```
