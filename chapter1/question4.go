package main

import (
	"fmt"
	"strings"
)

// 大文字と小文字を区別しない
// 空白文字は無視する

// 文字列が与えられたとき、文字列が回文の順列であるかを調べる関数を書いてください
// 順列とは文字を並べ替えたもの

func main() {
	input := "abcdefwwsfedcba"
	fmt.Print(ans2(input))
}

func ans(input string) bool{
	// 文字列の長さが偶数の時；登場回数がすべて偶数
	// 文字列の長さが奇数の時；登場回数が奇数の文字が1個
	// であれば、それは回文の順列である（回文になり得る）といえそう

	runes := []rune(strings.ToLower(input))
	length := len(runes)
	count := map[rune]int{}

	spaceCount := 0
	for _,v := range(runes) {
		if v == ' ' {
			spaceCount ++
		} else {
			count[v] ++
		}
	}

	oddCount := 0
	for _, v := range(count) {
		if v % 2 != 0 {
			oddCount ++
		}
	}

	if length-spaceCount % 2 == 0 {
		return oddCount == 0
	} else {
		return oddCount == 1
	}

}

// ビット演算を使ってみる
func ans2(input string) bool {
	runes := []rune(strings.ToLower(input))
	var checker int32 = 0

	spaceCount := 0
	for _, v := range(runes) {
		if v == ' ' {
			spaceCount ++
		} else {
			charAt := v - 'a'
			// XOR どちらか片方だけが1のときに1を返す（奇数なら1、偶数すなわち1:1になると0に戻る）
			checker ^= (1 << charAt)
		}
	}

	if (len(input)-spaceCount) % 2 == 0 {
		return checker == 0
	} else {
		oddCount := 0
		for i := 0; i < 32; i++ {
			// i個右シフトして1との論理和を取る
			// →i桁目が1だったらcountに1が加算される！
			oddCount += int((checker >> i) & 1)
		}
		return oddCount == 1
	}
}