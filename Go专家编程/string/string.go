package string

import (
	"fmt"
	"strings"
)

func ByteToString() {
	b := []byte{'H', 'e', 'l', 'l', 'o'}
	s := string(b)
	fmt.Println(s)
}

func StringToByte() {
	s := "Hello"
	b := []byte(s)
	fmt.Println(b)
}

func StringIteration() {
	s := "中国"
	for index, value := range s {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}
}

func Function() {
	s := "Hello World."

	fmt.Println(strings.Contains(s, "lo"))

	sSlice := strings.Split(s, "")
	fmt.Println(sSlice)

	s1 := strings.Join(sSlice, "")
	fmt.Println(s1)

	fmt.Println(strings.HasPrefix(s, "He"))

	fmt.Println(strings.HasSuffix(s, "lo"))

	fmt.Println(strings.ToUpper(s))

	fmt.Println(strings.ToLower(s))

	fmt.Println(strings.Trim(s, "."))

	fmt.Println(strings.TrimLeft(s, "H"))

	fmt.Println(strings.TrimRight(s, "."))

	fmt.Println(strings.TrimSpace(s))

	fmt.Println(strings.TrimPrefix(s, "He"))

	fmt.Println(strings.TrimSuffix(s, "."))

	fmt.Println(strings.Replace(s, "o", "i", 1))

	fmt.Println(strings.ReplaceAll(s, "o", "i"))

	s2 := "hello world."
	fmt.Println(strings.EqualFold(s1, s2))

}
