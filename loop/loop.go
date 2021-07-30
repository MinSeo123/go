package main

import "fmt"

func main (){
		for i:=2; i<=9; i++ {
			for j:=1; j<=9; j++ {
				fmt.Println(i, "X" ,j, "=", i*j )
			}
		}

	star()
}

func star () {
	for i:=1; i<=6; i++ {
		for j:=1; j<=i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}