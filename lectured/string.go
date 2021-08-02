package main

import "fmt"

func main() {

	str := "Hello 월드!"
	for _, v := range str {
		fmt.Printf(" 타입: %T 값: %d 문자:%c\n", v, v, v)
	}

//	a := `안녕하세요
//저는 고언어 입니다.`
//	b := []rune(a)
//	fmt.Println([]rune(b))
//	fmt.Println(a)
}
