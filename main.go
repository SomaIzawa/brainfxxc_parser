package main

import (
	"brainfxxk/brainfxxk"
	"brainfxxk/util"
	"fmt"
	"log"
)

func main()  {
	// ファイルパス取得
	var filepath string
	fmt.Scanf("%s", &filepath)
	// テキストファイルの読み込み
	code, err := util.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// brainfxxkインタプリタを作成
	brainfxxk := brainfxxk.NewParser(code ,128)
	// 実行
	brainfxxk.Exec()
}