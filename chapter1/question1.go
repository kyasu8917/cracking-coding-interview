package main

import(
	"fmt"
	"errors"
)

func main() {
	runes := []rune("asdfghjkoiuwqqjakkds")

	ans2(runes)
}

// 全探索
// O(N^2)
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


// バブルソート*二分探索
// O(N^2) + O(logN) = O(N^2 + logN)
func ans2(runes []rune) {
	runes = bubbleSort(runes)

	result := false
	for _, r := range runes {
		_, err := binarySearch(r, runes, 0, len(runes))
		if(err == nil) {
			result = true
		}
	}

	printResult(result)
}

// // クイックソート*二分探索
// // O(N logN) * O(logN) = O(N logN)
// func ans3(runes []rune) {

// }

// ハッシュテーブルを使用
// O(N)　早い。
func ans4(runes []rune) {
	result := false
	hashmap := map[rune]int{}
	for _, r := range(runes) {
		hashmap[r] ++
		if(hashmap[r] > 1) {
			result = true
		}
	}
	
	printResult(result)
}

// // クイックソート
// func quickSort(input []rune) []rune] {
	
// }

// バブルソート
// O(N^2)
func bubbleSort(input []rune) []rune {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if(input[i] < input[j]) {
				temp := input[j]
				input[j] = input[i]
				input[i] = temp
			}
		}
	}
	fmt.Print(string(input))
	return input
}

// 二分探索
func binarySearch(c rune, sorted []rune, imin int, imax int)(int, error){
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