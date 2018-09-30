package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	// Scanf scans text read from standard input
	// 格納すべきアドレスの場所を伝えるためにアドレス演算子が必要
	fmt.Scanf("%s\n", &input)
	answer := 0

	// Go言語では、文字列は実質的に読み取り専用のバイトのスライス
	// b := []byte(input)
	// fmt.Println(b)

	// nihongo := "日本語"
	// for index, runeValue := range nihongo {
	// 	fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	// }
	for _, rune := range input {

		str := string(rune)
		if strings.ToUpper(str) == str {
			answer++
		}
	}
	fmt.Println(answer)
}
