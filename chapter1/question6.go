package main

import (
	"fmt"
	"strconv"
)

// 文字列圧縮

// aabcccccaaa →a2b1c5a3
// 文字列は小文字と大文字のみ
// 圧縮したのが元の文字列より短くならない時は元の文字列を返す

func main(){
	input := "abcdefghijklmn"
	fmt.Println(ans([]rune(input)))

}

// ヒント2 繰り返し文字列を連結するのは非効率だよ。　→これは繰り返し文字列を連結している、、？していないよ。
func ans(input []rune) string {

	var conv []rune

	conv = append(conv, input[0])
	current := input[0]
	count := 1
	for i:=1; i<len(input); i++ {
		if(input[i]==current){
			count ++
		} else {
			conv = append(conv, []rune(strconv.Itoa(count))...)

			count = 1
			conv = append(conv, input[i])
			current = input[i]
		}
	}
	conv = append(conv, []rune(strconv.Itoa(count))...)

	if len(conv) > len(input) {
		return string(input)
	} else {
		return string(conv)
	}
}