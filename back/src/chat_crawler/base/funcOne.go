package base

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}
// 带指针的***难点***传值与传指针
func add(a *int) *int  {
	*a = *a +5
	return a
}

func Func1() {
	c := 1
	b := 2
	c, b = swap(c, b)
	fmt.Printf("%d,%d", c, b)

	fmt.Println("\n-----------")
	a := 1
	add(&a)
	fmt.Printf("%d",a)

}
