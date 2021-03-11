package main

import (
	"fmt"
)

// 行列の回転
// N＊Nの行列に描かれた、１ピクセルが4バイト四方の画像があります
// これを90度回転させるメソッドを、追加の領域なしで書いてください

func main() {
	count := 4
	input := [][]int{
		[]int{1,2,3,4},
		[]int{5,6,7,8},
		[]int{9,10,11,12},
		[]int{13,14,15,16},
	}

	for i:=0; i<count; i++ {
		fmt.Println(input[i])
	}

	fmt.Println("")

	ans(input, count)
}

// 追加の領域なし、、とは
func ans(input [][]int, count int) {
	// 各レイヤーの端っこを移動させていく
	// (temp的なのは追加の領域とは言わないっぽい)

	// for i:=0; i < count-1; i++ {

	// 	// 左上を保管
	// 	temp := input[0][i]

	// 	// 左下を左上に
	// 	input[0][i] = input[count-1][i]

	// 	// 右下を左下に
	// 	input[count-1][i] = input[count-1][count-1]

	// 	// 右上を右下に
	// 	input[count-1][count-1] = input[0][count-1]

	// 	// 右上に左上を
	// 	input[0][count-1] = temp 
	// }

	// for i:=0; i<count; i++ {
	// 	fmt.Println(input[i])
	// }

	// fmt.Println("")

	// 頭が悪くて嫌になる
}