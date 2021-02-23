package main

import (
	"fmt"
)

func main() {
	input := "Mr John Smith    "
	count := 13
	fmt.Println(ans(input,count))
}

func ans(input string, count int) string {
	runes := []rune(input)

	// 空白の数を数える
	spaceCount := 0
	for i := 0; i < count; i++ {
		if runes[i] == ' ' {
			spaceCount ++
		} 
	}

	// 空白の数から、返還後の文字配列のケツを決める
	index := count-1 + spaceCount*2

	// 後ろから走査する
	for i := count-1; i > -1; i-- {
		if runes[i] == ' ' {
			// 空白だったら3つ使って文字置き換える
			runes[index] = '0'
			runes[index-1] = '2'
			runes[index-2] = '%'
			index = index -3
		} else {
			// 空白じゃなかったら後ろから置いていく
			runes[index] = runes[i]
			index = index - 1
		}
	}
	// 返還後の文字配列の要素数はさっき決めてるから、これで上手くハマる

	return string(runes)
}