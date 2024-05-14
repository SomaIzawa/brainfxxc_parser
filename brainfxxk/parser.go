package brainfxxk

import (
	"brainfxxk/util"
	"fmt"
	"log"
)

type Parser struct {
	Code []string
	CIndex int
	Memory []int
	MIndex int
	LoopStart []int
	LSIndex int
}

func NewParser(code string, memorySize int) Parser{
	return Parser{
		Code: util.ExtractSpecificCharacters(code, []string{"+","-",">","<","[","]",".",","}),
		CIndex: 0,
		Memory: make([]int, memorySize),
		MIndex: memorySize / 2,
		LoopStart: make([]int, len(code)/2),
		LSIndex: -1,
	}
}

func (p *Parser) Exec() {
	for p.CIndex < len(p.Code) {
		//評価
		p.EvaluateCode(p.Code[p.CIndex])
		//p.ShowMemory()
	}
}

func (p *Parser) EvaluateCode(str string) {
	switch str {
	case "+":
		p.MCountUp()
	case "-":
		p.MCountDown()
	case ">":
		p.MPInc()
	case "<":
		p.MPDec()
	case "[":
		p.ProcessLoopStart()
	case "]":
		p.ProcessLoopEnd()
	case ".":
		p.Output()
	case ",":
		p.Input()
	}
}

func (p *Parser) MPInc(){
	if p.MIndex != (len(p.Memory) - 1) {
		p.MIndex++
	} else {
		log.Fatalf("out of index")
	}
	p.next()
}

func (p *Parser) MPDec(){
	if p.MIndex != 0 {
		p.MIndex--
	} else {
		log.Fatalf("out of index")
	}
	p.next()
}

func (p *Parser) MCountUp(){
	p.Memory[p.MIndex]++
	p.next()
}

func (p *Parser) MCountDown(){
	p.Memory[p.MIndex]--
	p.next()
}

func (p *Parser) Output(){
	s := fmt.Sprintf("%c", p.Memory[p.MIndex])
	fmt.Printf("%s", s)
	p.next()
}

func (p *Parser) Input(){
	var str string
  fmt.Scan(&str)
	firstChar := str[0]
	p.Memory[p.MIndex] = int(firstChar)
	p.next()
}

func (p *Parser) ProcessLoopStart()  {
	p.LSIndex++
	p.LoopStart[p.LSIndex] = p.CIndex
	p.next()
}

func (p *Parser) ProcessLoopEnd()  {
	if p.Memory[p.MIndex] > 0 {
		p.CIndex = p.LoopStart[p.LSIndex]
		p.next()
	} else {
		p.LoopStart[p.LSIndex] = 0
		p.LSIndex--
		p.next()
	}
}

func (p *Parser) next() {
	p.CIndex++
}

func (p *Parser) ShowMemory(width int){
	var head []int
	var output []int
	outputWidth := 0
	height := len(p.Memory) / width + 1
	fmt.Print("\n\n")
	println("Memory:")
	println("")
	for i, item := range p.Memory {
		if i % width == 0 && i / width + 1 < height {
			outputWidth = width
			util.OutPutLine(outputWidth)
		} else if i % width == 0 && i / width + 1 >= height {
			outputWidth = len(p.Memory) % width
			util.OutPutLine(outputWidth)
		}

		head = append(head, i)
		output = append(output, item)

		if (i % width == width - 1 && i / width + 1 < height) ||
		((i  % width) == (len(p.Memory) % width - 1) && i / width + 1 >= height) {
			util.OutPutValues(head)
			util.OutPutLine(outputWidth)
			util.OutPutValues(output)
			util.OutPutLine(outputWidth)
			head = []int{}
			output = []int{}
			util.OutPutEmptyLine()
		}
	}
}