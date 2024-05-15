package main

import (
	"brainfxxk/brainfxxk"
	"brainfxxk/util"
	"fmt"
	"log"
	"os"
)

func main()  {
	args := os.Args
	// ファイルパス取得
	var filepath string
	if len(args) > 1 {
		filepath = args[1]
	} else {
		fmt.Scanf("%s", &filepath)
	}
	// テキストファイルの読み込み
	code, err := util.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// brainfxxkインタプリタを作成
	brainfxxk := brainfxxk.NewParser(code ,128)
	// 実行
	brainfxxk.Exec()
	//brainfxxk.ShowMemory(25)
}