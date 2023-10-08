package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	TorF bool = false
	MaxInt uint64 = 1<<64 -1 
	z complex128 = cmplx.Sqrt(-5 + 12i)

)

func main () {
	var x,y float64 = 6,2
	h := math.Pow(x, y)
	p := math.Pow(h,x)


	fmt.Printf("Type : %T, Value: %v\n",TorF,TorF)
	fmt.Printf("Type : %T, Value: %v\n",MaxInt,MaxInt)
	fmt.Printf("Type : %T, Value: %v\n",z,z)
	fmt.Println(x,y,h,p)
	fmt.Printf("p is type of %T\n",p)	

}

