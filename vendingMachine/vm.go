package main

import "fmt"

//음료수 목록 출력 > 사용자가 돈 입력 > 입력받은 돈 출력 > 원하는 음료수 선택 > 거스름돈 반환 > 돈이 부족할 경우 사용종료



func main () {
	var userMoney int
	var selectBever string
	//음료수 목록
	beverage := map[string] int {
		"콜라" : 1000,
		"사이다" : 900,
		"환타" : 700,
	}
	//목록 출력
	for key, val := range beverage {
		fmt.Println("목록", key, val)
	}
	//돈입력
	fmt.Println("돈을 넣어주세요")
	fmt.Scan(&userMoney)
	//입력받은 돈 출력
	fmt.Println(userMoney,"원이 입력되었습니다 !")
	//음료수 선택하기, 잔돈 반환
	fmt.Println("이제 음료수를 선택해주세요 !")
	fmt.Scan(&selectBever)
	for key, val := range beverage {
		if key == selectBever {
			if val <= userMoney {
				println("선택하신", selectBever, "가 나왔습니다! 잔돈은 ", userMoney-val, "원 입니다 !")
			} else {
				println("돈이 부족합니다." )
			}
		}
	}
}
