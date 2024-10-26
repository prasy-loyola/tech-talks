package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	ifCount := 0
	whileCount := 0

	if len(args) < 3 {
		Error.Printf("Missing arguments.\nUsage: tlangc <command> <filename>\ncommand: i - interpret, c - compile\n")
		os.Exit(1)
	}

	if args[1] != "c" && args[1] != "i" {
		Error.Printf("Invalid argument '%s'.\nUsage: tlangc <command> <filename>\ncommand: i - interpret, c - compile\n", args[1])
		os.Exit(1)
	}

	var interpreter = args[1] == "i"

	if !interpreter {
		Assembly.Printf(`
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

	var filename string = args[2]
	var sourceFile, err = os.Open(filename)

	if err != nil {
		Error.Printf("Couldn't open the file '%s'.\n%s\n", filename, err.Error())
		os.Exit(1)
	}

	contentbytes, err := io.ReadAll(sourceFile)

	content := string(contentbytes)
	content = strings.ReplaceAll(content, "\n", " ")

	tokens := strings.Split(content, " ")

	for _, token := range tokens {
		token = strings.Trim(token, "\n")
		Info.Printf("Token: %v\n", token)
		Info.Printf("Stack: %v\n", stack)

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
				Assembly.Printf("; Token: %s", token)
				Assembly.Printf("push %d", number)
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
				Assembly.Printf("; Token: %s", token)

				// pop num1
				Assembly.Printf("pop rax")
				// pop num2
				Assembly.Printf("pop rdx")
				// add
				Assembly.Printf("add rax, rdx")
				// push result to stack
				Assembly.Printf("push rax")
			}

		} else if token == "-" {

			Assembly.Printf("; Token: %s", token)
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
				Assembly.Printf("pop rdx")
				// pop num2
				Assembly.Printf("pop rax")
				// add
				Assembly.Printf("sub rax, rdx")
				// push result to stack
				Assembly.Printf("push rax")
			}

		} else if token == "print" {
			if interpreter {
				if len(stack) < 1 {
					panic("Not enough item on stack for '$' operator")
				}
				num1 := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				Output.Printf("%d", num1)
			} else {

				Assembly.Printf("; Token: %s", token)
				Assembly.Printf("pop rax")
				Assembly.Printf("call Print_Number")
			}

		} else if token == "if" {
			if interpreter {
				panic("if condition Not Implemented")
			} else {
				Assembly.Printf("; Token: %s", token)
				Assembly.Printf(`

; IF_%d starts
; Duplicate value on stack
pop rax
push rax

; Test if zero
test rax, rax`, ifCount)
				Assembly.Printf("jz IF_%d_ELSE", ifCount)
				Assembly.Printf("IF_%d:", ifCount)
				ifCount++
			}

		} else if token == "else" {
			if interpreter {
				panic("If-Else not supported in interpreter")
			} else {
				Assembly.Printf("; Token: %s", token)
				Assembly.Printf("jmp IF_%d_THEN", ifCount) // jump after then before else
				Assembly.Printf("IF_%d_ELSE:", ifCount)
			}

		} else if token == "then" {
			if interpreter {
				panic("If-Else not supported in interpreter")
			} else {
				Assembly.Printf("; Token: %s", token)
				Assembly.Printf("IF_%d_THEN:", ifCount)
			}
		} else if token == "while" {
			if interpreter {
				panic("while not supported in interpreter")
			} else {
				Assembly.Printf("; Token: %s", token)
				Assembly.Printf(`

; WHILE_%d starts
WHILE_%d:
; Duplicate value on stack
pop rax
push rax

; Test if zero
test rax, rax`, whileCount, whileCount)
				Assembly.Printf("jz WHILE_%d_END", ifCount)
				whileCount++
			}
		} else if token == "end" {
			if interpreter {
				panic("while end not supported in interpreter")
			} else {
				Assembly.Printf("; Token: %s", token)
				Assembly.Printf("jmp WHILE_%d", ifCount)
				Assembly.Printf("WHILE_%d_END:", ifCount)
			}
		} else {
			panic(fmt.Sprintf("Unrecognized token: '%s'", token))
		}
		Info.Printf("After processing token Stack: %v\n", stack)

	}

	if !interpreter {
		Assembly.Printf(`
    mov  rax, SYS_EXIT
    mov  rdi, 1
    syscall

msg:
    db ' ', LN, 0
    msg_size = $ - msg


	DECIMAL DB "00000000000000000000", LN, 0; place to hold the decimal number
	DECIMAL_SIZE = $ - DECIMAL
	COUNT   DB 0

Print_Number:
	mov rbx, 10; divisor
	xor rcx, rcx; CX=0 (number of digits)

Getting_Digits_Loop:
	xor  rdx, rdx; Attention: DIV applies also RDX!
	div  rbx; RDX:RAX / BX = AX remainder: RDX
	push dx; LIFO
	inc  rcx; increment number of digits
	test rax, rax; RAX = 0?
	jnz  Getting_Digits_Loop; no: once more

	mov byte [COUNT], cl; storing the number of digits to memory
	mov rsi, DECIMAL; target string DECIMAL

Storing_Digits_Loop:
	pop  ax; get back pushed digit
	or   al, '0'; AL to ASCII
	mov  byte [rsi], al; save AL
	inc  rsi
	loop Storing_Digits_Loop; until there are no digits left

	mov byte [rsi], '$'; End-of-string delimiter for INT 21 / FN 09h

	;Print
	mov rax, SYS_WRITE
	mov rdi, STD_OUT
	mov rsi, DECIMAL
	xor rdx, rdx
	mov dl, byte [COUNT]
	syscall
	ret
 `)

	}

}
