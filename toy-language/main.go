package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var info *log.Logger = log.New(io.Discard, "[INFO] ", log.LstdFlags)
var output *log.Logger = log.New(os.Stdout, "", 0)
var assembly *log.Logger = log.New(os.Stdout, "", 0)

func main() {

	var interpreter = false

	if !interpreter {
		assembly.Printf(`
        format ELF64 executable 3

        ; syscalls
            SYS_EXIT = 60 
            SYS_WRITE = 1

        ; constants
            LN = 10
            STD_OUT = 1

            `)

	}

	var stack []int64 = []int64{}

	var filename string = "main.tlang"
	var sourceFile, err = os.Open(filename)

	if err != nil {
		info.Printf("Couldn't open the file %s. %s\n", filename, err.Error())
	}

	content, err := io.ReadAll(sourceFile)

	tokens := strings.Split(string(content), " ")

	for _, token := range tokens {
		token = strings.Trim(token, "\n")
		info.Printf("Token: %v\n", token)
		info.Printf("Stack: %v\n", stack)

		if strings.Trim(token, " ") == "" {
			continue
		}

		token_b := []byte(token)
		if token_b[0] >= '0' && token_b[0] <= '9' {
			number, err := strconv.ParseInt(token, 10, 32)
			if err != nil {
				panic("Couldn't parse number")
			}
			if interpreter {
				stack = append(stack, number)
			} else {
				assembly.Printf("; Token: %s", token)
				assembly.Printf("push %d", number)
			}

		} else if token == "+" {

			if interpreter {
				if len(stack) < 2 {
					panic("Not enough item on stack for '+' operator")
				}
				// pop num1
				num1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				// pop num2
				num2 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				// add
				sum := num1 + num2
				// push to stack
				stack = append(stack, sum)
			} else {
				assembly.Printf("; Token: %s", token)

				// pop num1
				assembly.Printf("pop rax")
				// pop num2
				assembly.Printf("pop rdx")
				// add
				assembly.Printf("add rax, rdx")
				// push result to stack
				assembly.Printf("push rax")
			}

		} else if token == "-" {

			assembly.Printf("; Token: %s", token)
			if interpreter {
				if len(stack) < 2 {
					panic("Not enough item on stack for '+' operator")
				}
				// pop num1
				num1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				// pop num2
				num2 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]

				// sub
				sum := num1 - num2
				// push to stack
				stack = append(stack, sum)
			} else {

				// pop num1
				assembly.Printf("pop rax")
				// pop num2
				assembly.Printf("pop rdx")
				// add
				assembly.Printf("sub rax, rdx")
				// push result to stack
				assembly.Printf("push rax")
			}

		} else if token == "$" { // Print
			if interpreter {
				if len(stack) < 1 {
					panic("Not enough item on stack for '$' operator")
				}
				num1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				output.Printf("%d", num1)
			} else {

				assembly.Printf("; Token: %s", token)
				assembly.Printf("pop rax")
				assembly.Printf("mov byte [msg], al")
				assembly.Printf("mov  rax, SYS_WRITE")
				assembly.Printf("mov  rdi, STD_OUT")
				assembly.Printf("mov  rsi, msg")
				assembly.Printf("mov  rdx, msg_size")
				assembly.Printf("syscall")
			}

		} else {
			panic(fmt.Sprintf("Unrecognized token: '%s'", token))
		}
		info.Printf("After processing token Stack: %v\n", stack)

	}

	if !interpreter {
		assembly.Printf(`
            mov  rax, SYS_EXIT
            mov  rdi, 1
            syscall

        msg:
            db ' ', LN, 0
            msg_size = $ - msg
            `)

	}

}
