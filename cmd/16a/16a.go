package day16a

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day16a", "Day 16 - First Part", run)
}

type opcode interface {
	execute(code [4]int, registers [4]int) [4]int
}

type addr struct {
}

type addi struct {
}

type mulr struct {
}

type muli struct {
}

type banr struct {
}

type bani struct {
}

type borr struct {
}

type bori struct {
}

type setr struct {
}

type seti struct {
}

type gtir struct {
}

type gtri struct {
}

type gtrr struct {
}

type eqir struct {
}

type eqri struct {
}

type eqrr struct {
}

func equal(r1, r2 [4]int) bool {
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

func (a addr) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] + registers[code[2]]
	return registers
}

func (a addi) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] + code[2]
	return registers
}

func (m mulr) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] * registers[code[2]]
	return registers
}

func (m muli) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] * code[2]
	return registers
}

func (b banr) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] & registers[code[2]]
	return registers
}

func (b bani) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] & code[2]
	return registers
}

func (b borr) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] | registers[code[2]]
	return registers
}

func (b bori) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]] | code[2]
	return registers
}

func (s setr) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = registers[code[1]]
	return registers
}

func (s seti) execute(code [4]int, registers [4]int) [4]int {
	registers[code[3]] = code[1]
	return registers
}

func (g gtir) execute(code [4]int, registers [4]int) [4]int {
	if code[1] > registers[code[2]] {
		registers[code[3]] = 1
	} else {
		registers[code[3]] = 0
	}
	return registers
}

func (g gtri) execute(code [4]int, registers [4]int) [4]int {
	if registers[code[1]] > code[2] {
		registers[code[3]] = 1
	} else {
		registers[code[3]] = 0
	}
	return registers
}

func (g gtrr) execute(code [4]int, registers [4]int) [4]int {
	if registers[code[1]] > registers[code[2]] {
		registers[code[3]] = 1
	} else {
		registers[code[3]] = 0
	}
	return registers
}

func (e eqir) execute(code [4]int, registers [4]int) [4]int {
	if code[1] == registers[code[2]] {
		registers[code[3]] = 1
	} else {
		registers[code[3]] = 0
	}
	return registers
}

func (e eqri) execute(code [4]int, registers [4]int) [4]int {
	if registers[code[1]] == code[2] {
		registers[code[3]] = 1
	} else {
		registers[code[3]] = 0
	}
	return registers
}

func (e eqrr) execute(code [4]int, registers [4]int) [4]int {
	if registers[code[1]] == registers[code[2]] {
		registers[code[3]] = 1
	} else {
		registers[code[3]] = 0
	}
	return registers
}

var result int
var opcodes []opcode

func testOpCode(registersBefore, op, registersAfter string) {
	var bArr, code, aArr [4]int

	_, err := fmt.Sscanf(registersBefore, "Before: [%d, %d, %d, %d]", &bArr[0], &bArr[1], &bArr[2], &bArr[3])
	if err != nil {
		panic(err)
	}
	_, err = fmt.Sscanf(op, "%d %d %d %d", &code[0], &code[1], &code[2], &code[3])
	if err != nil {
		panic(err)
	}

	_, err = fmt.Sscanf(registersAfter, "After: [%d, %d, %d, %d]", &aArr[0], &aArr[1], &aArr[2], &aArr[3])
	if err != nil {
		panic(err)
	}

	successfulOp := 0

	for _, op := range opcodes {
		var in [4]int
		copy(in[:], bArr[:])
		result := op.execute(code, in)
		if equal(result, aArr) {
			successfulOp++
		}
	}

	if successfulOp >= 3 {
		result++
	}
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	opcodes = []opcode{addi{}, addr{}, mulr{}, muli{}, banr{}, bani{}, borr{}, bori{}, setr{}, seti{}, gtir{}, gtri{}, gtrr{}, eqir{}, eqri{}, eqrr{}}

	for i := 0; i < len(input); i += 4 {
		if !strings.HasPrefix(string(input[i]), "Before") {
			break
		}
		testOpCode(input[i], input[i+1], input[i+2])
	}

	fmt.Println(result)
}
