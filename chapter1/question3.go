package main

import (
	"fmt"
)

// 文字列に登場するすべての空白文字を"%20"で置き換えるメソッド
// 文字列の後ろには追加用スペースが十分にある。バッファのサイズは気にしなくて良い
// 追加用スペースを除いた文字列の真の長さは与えられる

func main() {
	input := "Mr John Smith    "
	count := 13
	fmt.Println(ans(input,count))
	fmt.Println(ans2(input,count))
	fmt.Println(ans3(input,count))
}

func ans(input string, count int) string {
	result := ""
	runes := []rune(input)

	for i := 0; i < count; i++ {
		if runes[i] == ' ' {
			result = result + "%20"
		} else {
			result = result + string(runes[i])
		}
	}

	return result

}

// 追加の領域を使わない
func ans2(input string, count int) string {
	runes := []rune(input)

	for i := 0; i < count; i++ {
		if runes[i] != ' ' {
			continue
		}

		for j := count - 1; j > i; j-- {
			runes[j+2] = runes[j]
		}

		runes[i] = '%'
		runes[i+1] = '2'
		runes[i+2] = '0'

		count = count + 2
	}

	return string(runes)

}