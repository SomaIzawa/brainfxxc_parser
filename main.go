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
	brainfxxk := brainfxxk.NewParser(util.ExtractSpecificCharacters(code, []string{"+","-",">","<","[","]",".",","}) ,128)
	brainfxxk.Exec()
}