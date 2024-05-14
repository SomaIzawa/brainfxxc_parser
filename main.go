package main

import (
	"brainfxxk/brainfxxk"
	"brainfxxk/util"
	"log"
)

func main()  {
	code, err := util.ReadFile("source/hello.bf")
	if err != nil {
		log.Fatal(err)
	}
	brainfxxk := brainfxxk.Parser{
		Code: util.ExtractSpecificCharacters(code, []string{"+","-",">","<","[","]","."}),
		CIndex: 0,
		Memory: make([]int, 16),
		MIndex: 7,
		LoopStart: make([]int, 16),
		LSIndex: -1,
	}
	brainfxxk.Exec()
}