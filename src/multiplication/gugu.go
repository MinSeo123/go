package main

import "fmt"




	func Gugu(a, b int) int {
		//계산로직
		return a*b
	}


	func main() {
		var a int
		var b int
		//입력값 받기
		n, err := fmt.Scan(&a, &b)
		//계산
		if err == nil {
			fmt.Println(Gugu(a, b))
		} else {
			fmt.Println(err, n)
		}

	}
