package main

import (
	"fmt"
)

// ブルートフォースのアルゴリズムで解くことができるらしい


func main() {
	input := "pale"
	compare := "bale"
	fmt.Print(ans([]rune(input),[]rune(compare)))
}

// 悲しき。。
func ans(input []rune, compare []rune) bool {

	// ここはほとんど一緒
	if len(input) == len(compare) {
		foundDefference := false
		for i:=0;i<len(input);i++{
			if input[i] != compare[i] {
				if foundDefference {
					return false
				}
			}
			foundDefference = true
		}
		return true
	}

	// 文字数が違う １つ違い
	var longer []rune
	var shorter []rune
	if len(input) >= len(compare) {
		longer = input[0:len(input)]
		shorter = compare[0:len(compare)]
	} else {
		longer = compare[0:len(compare)]
		shorter = input[0:len(input)]
	}

	for i:=0; i<len(longer);i++{
		//ここはまた後日
		
		// ラスト文字まで残ったらそれはtrue
		// if i == len(longer)-1 {
		// 	return true
		// }

		// if longer[i] != shorter[i] {
		// 	other ++
		// 	if other > 1 {
		// 		return false
		// 	}

		// 	if longer[i+1] == shorter[i] {
		// 		shorter = append(shorter[:i+1], shorter[i:]...)
		// 		shorter[i] = longer[i]
		// 	} else if longer[i] == shorter[i+1] {
		// 		longer = append(longer[:i], longer[i+1:]...)
		// 	} else {
		// 		return false
		// 	}
		// }
	}

	return true

}