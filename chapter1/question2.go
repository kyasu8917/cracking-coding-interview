package main


import (
		"fmt"
)

// 前提：完全一致は「並び替えである」とする
// 文字コードはUnicode(日本語も含むことができるように）

func main() {
	input := "fabcdeffg"
	compare := "gfedfcbfaz"

	printResult(isSort1(input, compare))
}

// O(3N) = O(N)
func isSort(input string, compare string) bool {
	// 全てが重複していればtrue
	// なおかつ個数も同じ
	var result = true

	// 文字数が違う場合は論外
	if len(input) != len(compare) {
		return false
	}

	inputRunes := []rune(input)
	compareRunes := []rune(compare)
	
	// golangは実行時に値が計算されるものを配列の要素数に指定できない
	// 定数配列はコンパイル時に生成されるため。
	// そのためSliceを使用する
	// input,compareの長さのSliceを定義　初期値はすべて0になる
	var inputRuneCount = make([]int, len(input))
	var compareRuneCount = make([]int, len(compare))
	
	for _, r := range(inputRunes) {
		// これ、結局Unicodeを考慮してないな。。。
		charAt := r - 'a'
		inputRuneCount[charAt] ++
	}

	for _, r := range(compareRunes) {
		// これ、結局Unicodeを考慮してないな。。。
		charAt := r - 'a'
		compareRuneCount[charAt] ++
	}

	for i, c := range(inputRuneCount) {
		if c != compareRuneCount[i] {
			result = false
		}
	}

	return result
}


func printResult(result bool) {
	if(result){
		fmt.Print("Sorted!!\n")
	} else {
		fmt.Print("No Sorted!\n")
	}
}