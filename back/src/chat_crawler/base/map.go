package base

import "fmt"

func Map()  {
	// 对象
	var student map[string]float32
	// 使用前必须要make这一步骤
	student = make(map[string]float32)
	student["zhangsan"] = 18.2
	fmt.Printf("%v",student)

	// 对象定义与实例化  ***简写版
	// 使用前必须要make这一步骤
	student2 := make(map[string]float32)
	student2["zhangsan"] = 19.2
	fmt.Printf("%v",student2)
}
