package brainfxxk

import (
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

func (p *Parser) ShowMemory(){
	fmt.Println(p.Memory)
}

func (p *Parser) Output(){
	s := fmt.Sprintf("%c", p.Memory[p.MIndex])
	fmt.Printf("%s", s)
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