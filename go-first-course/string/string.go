package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	encodeRune()

	decodeRune()

	runeCount()
}

func encodeRune() {
	var r rune = 0x4E2D
	fmt.Printf("the unicode charactor is %c\n", r)
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r) // 对 rune 进行 utf-8 编码
	fmt.Printf("utf-8 representation is 0x%x\n", buf)
}

func decodeRune() {
	var buf = []byte{0xE4, 0xB8, 0xAD}
	r, _ := utf8.DecodeRune(buf) // 对 buf 进行 utf-8 解码
	fmt.Printf("the unicode charactor after decoding [0xE4, 0xB8, 0xAD] is %s\n", string(r))
}

func runeCount() {
	var v string = "中国人"
	fmt.Println(utf8.RuneCountInString(v))
}
