package main

import (
	"fmt"
)

// 文字列に対して、文字の挿入、文字の削除、文字の置き換えを行うことができる。
// 2つの文字列が与えられたとき、一方の文字に対して１操作（もしくは操作無し）でもう一方の文字列にできるかどうか

// 例
// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false

// 文字の入れ替えは無し

// 要するに、違いがひとつしかない（もしくはひとつもない）ことを確認すれば良い
// 順番も大事

func main() {
	input := "pale"
	compare := "bale"
	fmt.Print(ans([]rune(input),[]rune(compare)))
}

// 悲しき。。
func ans(input []rune, compare []rune) bool {

	// 文字数が2つ以上違ったら論外
	if len(input) - len(compare) > 1 {
		return false
	}

	// 異なる文字のカウント
	other := 0

	// 文字数が同じ = 違う文字が1つだけか？
	if len(input) == len(compare) {
		for i:=0;i<len(input);i++{
			
			if input[i] != compare[i] {
				other ++
			}

			if other > 1 {
				return false
			}
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

		// ラスト文字まで残ったらそれはtrue
		if i == len(longer)-1 {
			return true
		}

		if longer[i] != shorter[i] {
			other ++
			if other > 1 {
				return false
			}

			if longer[i+1] == shorter[i] {
				shorter = append(shorter[:i+1], shorter[i:]...)
				shorter[i] = longer[i]
			} else if longer[i] == shorter[i+1] {
				longer = append(longer[:i], longer[i+1:]...)
			} else {
				return false
			}
		}
	}

	return true

}