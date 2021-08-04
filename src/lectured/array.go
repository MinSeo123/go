package main

import "fmt"

func main ()  {

	a := [2][5]int{
		{1,2,3,4,5},
		{5,6,7,8,9},
	}
	for _, arr := range a {
		for _, v := range  arr{
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

 	//a := [5]int{1,2,3,4,5}
 	//b := [5]int{500, 400, 300, 200, 100}
	//
 	//for i, v := range a {
 	//	fmt.Printf("a[%d] = %d\n", i, v)
	//}
	//
	//fmt.Println()
 	//for i, v:= range b {
 	//	fmt.Printf("b[%d] = %d\n", i, v)
 	//}
	//b = a
	//
	//fmt.Println()
	//
	//for i, v := range b {
	//	fmt.Printf("b[%d] = %d\n", i, v)
	//}

	//var a int
	//
	//_, err := fmt.Scan(&a)
	//if err == nil {
	//	fmt.Println(a)
	//} else {
	//	fmt.Println(err)
	//}


	// _ 는 안쓰는 변수
	//for i, v := range t {
	//	fmt.Println(i, v)
	//}



	//t := [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	//days := [5]string{"mon", "tue", "wed", "tur", "fri"}
	////요소 개수 생략
	//x := [...] int{10, 20, 30}
	////for i:=0; i<5; i++ {
	////	fmt.Println(t[i])
	////	fmt.Println(days[i])
	////}
	//fmt.Println(t)
	//fmt.Println(days)
	//fmt.Println(x)
}
