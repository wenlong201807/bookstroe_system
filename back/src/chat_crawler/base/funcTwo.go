package base

import "fmt"

func getSum(num []int)int  {
	sum := 0
	for i:=0;i<len(num);i++{
		sum += num[i] // 注意写法是[]
	}
	return sum
}
func FuncTwo()  {
	num := []int{1,2,3,4,5,6,7,8}
	fmt.Println(getSum(num))
}
