package main

import "fmt"

func main() {

	//pop기능
	slice := []int{1,2,3,4,5,6}
	t, slice := slice[len(slice)-1], slice[:len(slice)-1]
	fmt.Println(t, slice)

	//slice := []int{1,2,3,4,5,6,7,8,9,10}
	//result := slice[len(slice)-2:]
	//fmt.Println(result)

	////요소 추가 방법
	//slice := []int{1,2,3,4,5,6}
	////맨뒤에 요소 추가
	//slice = append(slice, 8)
	////추가하려는 위치
	//idx := 2
	//for i := len(slice)-2; i>=idx; i-- {
	//	slice[i+1] = slice[i]
	//}
	//
	//slice[idx] =100
	//
	//fmt.Println(slice)

	//요소 삭제 방법
	//slice := []int{1,2,3,4,5,6}
	//idx := 4
	//slice = append(slice[:idx], slice[idx+1:]...)
	////
	////for i := idx+1; i < len(slice); i++ {
	////	slice[i-1] = slice[i]
	////}
	////
	////slice = slice[:len(slice)-1]
	//fmt.Println(slice)

	//slice1 := []int{1,2,3,4,5}
	//slice2 := make([]int, 3, 10)
	//slice3 := make([]int, 10)
	//
	//cnt1 := copy(slice2, slice1)
	//cnt2 := copy(slice3, slice1)
	//
	//fmt.Println(cnt1, slice2)
	//fmt.Println(cnt2, slice3)




}
