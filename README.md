# ROT13 Encoder
* This program takes a string as input and displays it with each of its letters replaced by the letter 13 spots ahead in the alphabetical order. 
* The transformation maintains the case of the letters (i.e., 'a' becomes 'n' and 'A' becomes 'N'). 
* Non-alphabetical characters remain unchanged. The output is followed by a newline (\n).

## Instructions

* The program processes a single argument, which is the string to be transformed.
* If the number of arguments is different from 1, the program will display nothing.

## Usage

The program is written in Go. To run it, use the go run command followed by the program and the string argument. Below are some examples:

```bash 

$ go run . "abc"
nop

$ go run . "hello there"
uryyb gurer

$ go run . "HELLO, HELP"
URYYB, URYC

$ go run .
$

```
