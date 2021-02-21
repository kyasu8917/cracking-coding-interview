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
	input := "abcdefghijklmnopqrstu"
	ans2(input)
}

// rune -> Int32のエイリアス
// 'a' -> 99
// 'z' -> 112
// 'z' - 'a' = 25
func runeTest() {
	runes := []rune{'a','b','c','z'}
	fmt.Print(runes)
	fmt.Print(runes[3] - runes[0])
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
	// // 文字を小文字26種類と仮定する時、Int32型のビット演算でメモリ効率化ができる

	var result = false
	runes := []rune(input)
	var checker = 0
	for _, r := range(runes) {
		charAt := r - 'a'

		// << -> 左シフト. 左辺の数字を右辺ビット左へシフトする
		// ex) 00000010 (3) << 3 => 00010000 になる
		// この時、「文字コードでn番目の文字」をビットの1で表現している
		// そして、checkerと論理和をとって、
		// 1を超える　=　同じビット番目が1-1になっている = 重複していると判断する 
		if(checker & (1 << charAt) > 0) {
			result = true
			break
		}
		checker |= (1 << charAt)
	}

	printResult(result)
}

func printResult(result bool) {
	if(result) {
		fmt.Print("Duplicate")
	} else {
		fmt.Print("No Duplicate")
	}
}