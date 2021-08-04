package main

import (
	"fmt"
	"math/rand"
	"time"
)
//100 이하의 랜덤한 초기값(int 타입 숫자)하나 생성 > 입력값 받기 > 입력값과 초기값 비교 > 크다면 down > 작다면 up > 같다면 반복문 탈출 후 종료>
//처음 정해진 숫자에 따라 입력할 수 있는 횟수가 정해진다
func main() {
	//초기값
	rand.Seed(time.Now().UnixNano())
	firstNum := rand.Intn(100)
	fmt.Println(firstNum)
	//입력값 받기
	var putNum int
	//비교
	for i:= 1; i <= firstNum; i++ {
		_, err := fmt.Scan(&putNum)
		if err == nil {
			if firstNum > putNum {
				fmt.Println("입력하신 숫자가 더 작습니다.","남은횟수 :", firstNum-i)
			} else if firstNum < putNum {
				fmt.Println("입력하신 숫자가 더 큽니다.", "남은횟수 :", firstNum-i)
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한횟수 :" ,i)
				break
			}
		} else {
			fmt.Println(err)
		}
	}
}
