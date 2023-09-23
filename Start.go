package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main(){

	dateTime := time.Date(2023, 9, 22, 0, 0, 0, 0, time.Local)

	fmt.Println("Hello world!!")
	fmt.Println("The time is today is ",dateTime)	
	fmt.Println("Print rand number \n", rand.Intn(50))
	fmt.Println("I love you ", math.Sqrt(9.3))
	fmt.Println(math.Pi)
}