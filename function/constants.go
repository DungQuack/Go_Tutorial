package main

import("fmt"
)

const d = "DepTroai"
const(
	big = 1<<10
	small = big >> 9
)
func needInt(x int) int {
	return x*10 + 2
}
func needFloat(x float64)float64{
	return x *0.1
}

func main (){
	fmt.Println("DungQuack is ", d)
	fmt.Println(needInt(small))
	fmt.Println(needFloat(big))
}