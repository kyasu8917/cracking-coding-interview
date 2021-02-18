package main

import (
	"fmt"
	"unicode/utf8"
)

// 前提を確認する。
// 例えばASCIIなら128文字しかないとか、Unicodeならめちゃたくさんあるとか。
// 文字列を操作することができるか
// 追加のデータ構造を用いることができるか　など。

func main() {
	input := "qwertyuiahjk"
	ans(input)
}

// O(N)
func ans(input string) {
	// 文字コードASCIIと仮定する
	if !(utf8.ValidString(input) && utf8.RuneCountInString(input) == len(input)) {
		fmt.Print("error! out of ASCII")
		return
	}

	var result = false

	var rune_set[128] bool
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if rune_set[runes[i]] {
			result = true
		}
		rune_set[runes[i]] = true
	}

	printResult(result)
}

// ビットベクトルを用いることで消費メモリを削減する
func ans2(input string) {

}

func printResult(result bool) {
	if(result) {
		fmt.Print("Duplicate")
	} else {
		fmt.Print("No Duplicate")
	}
}