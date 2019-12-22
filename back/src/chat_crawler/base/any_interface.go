package base

import "fmt"

func AnyInterface() {
	var v1 interface{}
	v1 = "zhu"
	fmt.Println("字符串", v1)

	v1 = 666
	fmt.Println("数字v1", v1)
	if v, ok := v1.(int); ok {
		fmt.Println("来查看v1的数据类型", v, ok)
	}

	v1 = 1.23
	fmt.Println("浮点类型v1", v1)
	if v, ok := v1.(float64); ok {
		fmt.Println("来查看v1的数据类型", v, ok)
	}

	v1 = []int{1, 2, 3}
	fmt.Println("数组v1", v1)

	v1 = make(map[string]string)
	fmt.Println("对象v1", v1)

	// 因此需要判断interface的数据类型
	switch v := v1.(type) {
	case int:
		fmt.Println("int ： ", v)
	case float64:
		fmt.Println("float64 ： ", v)
	case float32:
		fmt.Println("float32 ： ", v)
	default:
		fmt.Println("default:我是一定会有内容输出的", v)
	}
}
