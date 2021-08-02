package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)


//45개의 숫자 배열 생성 > 크기가 6인 빈 배열 생성 > 빈 배열에 45개의 배열중 랜덤한 숫자 삽입 >
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
