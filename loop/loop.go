package main

import "fmt"

func Gugudan () {
	//구구단
	for i:=2; i<=9; i++ {
		for j:=1; j<=9; j++ {
			fmt.Println(i, "X" ,j, "=", i*j )
		}
	}
}

func Star () {
	//별찍기
	for i:=1; i<=6; i++ {
		for j:=1; j<=i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main (){
	Gugudan()
	Star()
}

