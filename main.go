package main

import (
	"brainfxxk/brainfxxk"
	"brainfxxk/util"
	"fmt"
	"log"
)

func main()  {
	var filepath string
	fmt.Scanf("%s", &filepath)
	code, err := util.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	brainfxxk := brainfxxk.Parser{
		Code: util.ExtractSpecificCharacters(code, []string{"+","-",">","<","[","]",".",","}),
		CIndex: 0,
		Memory: make([]int, 16),
		MIndex: 7,
		LoopStart: make([]int, 16),
		LSIndex: -1,
	}
	brainfxxk.Exec()
}