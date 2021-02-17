package main

import(
	"fmt"
	"errors"
)

// Question 1.1
// ある文字列が、全て固有である（重複する文字がない）かどうかを判定するアルゴリズムを実装してください。
// また、それを実装するのに新たなデータ構造が使えない場合、どのようにすればよいですか？

func main() {
	const input = "abcdefghijklss"
	ans1(input)
}

// 単純なループ処理
// N*1/2N => O(N^2)
func ans1(input string) {
	var result = false

	for i := 0; i < len(input); i++ {
		for j := i+1; j < len(input); j++ {
			if(input[i] == input[j]) {
				result = true
			}
		}
	}

	printResult(result)
}

// 二分探索を使用？
func ans2(input string) {

}

// ハッシュテーブルを使用？
func ans3(input string) {

}

// クイックソート
func quickSort(input string) string {	
}

// 二分探索
func binarySearch(c byte, sorted string, imin int, imax int)(int, error){
	if(imin > imax) {
		return 0, errors.New("Not Found!")
	}

	imid := imin + (imax - imin)/2

	if(sorted[imid] > c) {
		return binarySearch(c, sorted, imin, imid-1)
	} else if (sorted[imid] < c) {
		return binarySearch(c, sorted, imid+1, imax)
	} else {
		return imid, nil
	}
}

func printResult(result bool) {
	if(result){
		fmt.Print("Duplicate!\n")
	} else {
		fmt.Print("No Duplicate!\n")
	}
}