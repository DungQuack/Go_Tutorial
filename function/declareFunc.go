package main

import ("fmt"
)

func Add(num1 int, num2 int) int{
	return num1 + num2
}
func Mul(num1 int, num2 int) int{
	return num1 * num2
}
func Div(num1, num2 float64) float64{
	return num1 / num2
}
func Min(num1 int, num2 int) int{
	return num1 - num2
}


func swap(a ,b ,c string) (string,string,string){
	return b, c, a
}


func split(sum float64) (v,z float64){
	z = sum / 9 
	v = sum - z
	return
}


func main(){
	st,nd,rd := swap("hello", "Go","!")

	fmt.Println(Add(1, 2))
	fmt.Println(Div(19, 2))
	fmt.Println(Mul(18, 2))
	fmt.Println(Min(100, 2))

	fmt.Println(st,nd,rd)

	fmt.Println(split(20))
}