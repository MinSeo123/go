package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)


//크기가 6인 빈 배열 생성 > 크기가 6인 빈 배열에 랜던숫자 삽입 > 중복값 제거
func main() {
	lottoNumb := make([]int, 6)
	keys := make(map[int]bool)
	sameNumb := []int{}

	for i := 0; i<6; i++ {
		rand.Seed(time.Now().UnixNano())
		lottoNumb[i] = rand.Intn(45) + 1
	}

	for _, value := range lottoNumb {
		if _,saveValue := keys[value]; !saveValue {
			keys[value] = true
			sameNumb = append(sameNumb, value)
		}
	}
	if len(sameNumb) <= 5 {
		rand.Seed(time.Now().UnixNano())
		sameNumb = append(sameNumb, rand.Intn(45))
	}

	sort.Ints(sameNumb)

	fmt.Println(sameNumb)

}

//문제점 ㅡㅡ 중복값을 제거하게되면 갯수가 6개가 아닌 5개가 나오는 문제 발생 하여 6개가 정확하게 맞아 떨어지는 날
//나는 로또를 살 것이다.