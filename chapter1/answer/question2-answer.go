package main


import (
		"fmt"
)

// ＜前提＞
// 大文字と小文字は区別するのか？？
// 空白文字はどのように扱うのか？？

func main() {
	// 可読性の観点では、文字をソートして比較するのもアリ。（オーダーはソートの計算量）
}

func isSort(input string, compare string) bool {
	if len(input) != len(compare) {
		return false
	}

	inputRunes := []rune(input)
	compareRunes := []rune(compare)
	
	var runeCount = make([]int, len(input))
	
	// inputの文字をインクリメントする
	for _, r := range(inputRunes) {
		charAt := r - 'a'
		runeCount[charAt] ++
	}

	// compareの文字をデクリメントする
	// この時、数が負の数になったら終了。（一致しない。）
	// 正の有無を見る必要がないのは、文字数が一緒だから。（負の値が存在する時のみ、正の値が存在する条件。）
	for _, r := range(compareRunes) {
		charAt := r - 'a'
		runeCount[charAt] --
		if(runeCount[charAt] < 0) {
			return false
		}
	}

	return true
}

func printResult(result bool) {
	if(result){
		fmt.Print("Sorted!!\n")
	} else {
		fmt.Print("No Sorted!\n")
	}
}