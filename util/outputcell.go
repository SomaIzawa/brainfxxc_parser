package util

import (
	"fmt"
	"strconv"
	"strings"
)

func OutPutLine(width int){
	for i:=0; i<width; i++ {
		fmt.Print("+---")
	}
	fmt.Print("+\n")
}

func OutPutEmptyLine(){
	fmt.Print("\n")
}

func OutPutValues(array []int){
	for _, item := range array {
		fmt.Print("|")
		fmt.Print(OutPutIntOnXByte(item, 3))
	}
	fmt.Print("|\n")
}

func OutPutIntOnXByte(input interface{}, byte int) string{
	var inputStr string

	switch v := input.(type) {
	case int:
		inputStr = strconv.Itoa(v)
	case string:
		inputStr = v
	default:
		return ""
	}

	numSpaces := byte - len(inputStr)
	if numSpaces < 0 {
		numSpaces = 0
	}

	outputStr := strings.Repeat(" ", numSpaces) + inputStr
	return outputStr
}