package main

import (
	"fmt"
	"math/rand"
	"time"
)

//가진돈은 1000원 > 1 ~5 사이의 값을 입력받는다 > 1~5사이 랜덤한 값 선택 > 입력값과 랜덤값이 같으면 가진 돈에 500원 추가 ( 메시지와 잔액 표시) > 다른경우 가진돈에서 100원 뺌 (메시지와 잔액) > 가진돈이 0원 이하나 5000원 이상이 되면 게임 종료

func  main() {
 //가진돈
	haveMoney := 1000

//입력값 받기
	var putNum int
	for true {
		//랜덤값
		rand.Seed(time.Now().UnixNano())
		randomNum := rand.Intn(5)
		fmt.Println("값을 입력해주세요 1~5")
		fmt.Scan(&putNum)
		//돈 비교
		if haveMoney <= 0 || haveMoney >=5000 {
			fmt.Println("게임이 종료됩니다.")
			break
		} else {
			if putNum == randomNum {
				haveMoney = haveMoney + 500
				fmt.Println("축하합니다! +500 잔액 :", haveMoney)
			} else {
				haveMoney = haveMoney - 100
				fmt.Println("아쉽습니다 ㅠ -100 잔액 :", haveMoney)
			}
		}
	}



}
