package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	PUSH = iota
	POP
	ADD
	SUB
	PRINT
	PUTS
	JUMP
)

var InstructionTypeNames = []string{
	"PUSH",
	"POP",
	"ADD",
	"SUB",
	"PRINT",
	"PUTS",
	"JUMP",
}

const MIN_SUPPORTED_CLASS_VERSION = 1

type InstructionType int8

type Instruction struct {
	Type   InstructionType
	Params []int64
}

type Program struct {
	Instructions []Instruction
}

func (p *Program) saveToFile() {
	Info.Printf("Program : %#v", p)
	fileName := "program.class"
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		Error.Printf("%v", err)
	}
	var bin_buf bytes.Buffer
	magicByte := []byte("TLANG")
	magicByte = append([]byte{0}, magicByte...)
	binary.Write(&bin_buf, binary.BigEndian, magicByte)
	version := int16(1)
	binary.Write(&bin_buf, binary.BigEndian, version)
	binary.Write(&bin_buf, binary.BigEndian, int64(len(p.Instructions)))
	for _, instruction := range p.Instructions {
		binary.Write(&bin_buf, binary.BigEndian, int8(instruction.Type))
		if instruction.Params != nil {
			binary.Write(&bin_buf, binary.BigEndian, int8(len(instruction.Params)))
			for _, param := range instruction.Params {
				binary.Write(&bin_buf, binary.BigEndian, int64(param))
			}
		} else {
			binary.Write(&bin_buf, binary.BigEndian, int8(0))
		}
	}
	_, err = file.Write(bin_buf.Bytes())

	if err != nil {
		Error.Printf("%v", err)
	}
}

func readFromFile() Program {

	program := Program{}

	fileName := "program.class"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		Error.Printf("%v", err)
	}

	magicByte := readNextBytes(file, 6)
	if magicByte[0] != 0 && string(magicByte[1:]) != "TLANG" {
		Error.Printf("The file is not valid TLANG class file")
		return program
	}

	buffer := bytes.NewBuffer(readNextBytes(file, 2))
	var version int16
	binary.Read(buffer, binary.BigEndian, &version)
	if version < MIN_SUPPORTED_CLASS_VERSION {
		Error.Printf("Version %d is not supported. Minimum supported version %d", version, MIN_SUPPORTED_CLASS_VERSION)
		return program
	}
	buffer = bytes.NewBuffer(readNextBytes(file, 8))
	var instructionsCount int64
	binary.Read(buffer, binary.BigEndian, &instructionsCount)
	Info.Printf("Expecting %d instructions", instructionsCount)
	var instructions = make([]Instruction, instructionsCount)

	for i := int64(0); i < instructionsCount; i++ {
		buffer = bytes.NewBuffer(readNextBytes(file, 1))
		instruction := Instruction{}
		binary.Read(buffer, binary.BigEndian, &instruction.Type)
		buffer = bytes.NewBuffer(readNextBytes(file, 1))
		var paramCount int8
		binary.Read(buffer, binary.BigEndian, &paramCount)
		Info.Printf("Instruction %d Expecting %d params", i, paramCount)
		params := make([]int64, paramCount)
		for j := 0; j < int(paramCount); j++ {
			buffer = bytes.NewBuffer(readNextBytes(file, 8))
			var param int64
			binary.Read(buffer, binary.BigEndian, &param)
            params[j] = param
		}
        instruction.Params = params
        instructions[i] = instruction
	}
	program.Instructions = instructions
	return program
}

func readNextBytes(file *os.File, count int) []byte {

	buf := make([]byte, count)

	_, err := file.Read(buf)

	if err != nil {
		Error.Printf("%v", err)
	}
	return buf
}

type Parser struct {
	tokens []string
}

func (p *Parser) Parse() Program {
	instructions := make([]Instruction, 0)

	for _, token := range p.tokens {
		if strings.Trim(token, " ") == "" {
			continue
		}
		Info.Printf("Token: %s", token)

		token_b := []byte(token)
		if token_b[0] >= '0' && token_b[0] <= '9' {
			number, err := strconv.ParseInt(token, 10, 64)
			if err != nil {
				panic("Couldn't parse number")
			}
			params := make([]int64, 1)
			params[0] = number
			instructions = append(instructions, Instruction{
				Type:   PUSH,
				Params: params,
			})
		} else if token == "print" {
			instructions = append(instructions, Instruction{
				Type:   PRINT,
				Params: make([]int64, 0),
			})
		} else {
			panic(fmt.Sprintf("Unrecognized token %v", token))
		}
	}
	return Program{
		Instructions: instructions,
	}
}
