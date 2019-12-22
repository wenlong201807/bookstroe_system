package base

import "fmt"

func NilFunc()  {
	f := func(x,y int)int{
		return x+y
	}
	fmt.Println(f(2,3))
}
