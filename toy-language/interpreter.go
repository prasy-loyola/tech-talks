package main

import "fmt"

func Interpret(program Program) {
	stack := make([]int64, 0)
	for _, instuction := range program.Instructions {
		Info.Printf("Instruction: %v", instuction)

		switch instuction.Type {
		case PUSH:
			if len(instuction.Params) < 0 {
				panic("Not enough params for PUSH")
			}
			stack = append(stack, instuction.Params[0])
			break
		case PRINT:
			if len(stack) < 1 {
				panic("Not enough values on stack for PRINT")
			}
			value := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            Output.Printf("%d", value)
            break
        default:
            panic(fmt.Sprintf("Unsupported instruction %v", instuction))
		}
	}
}
