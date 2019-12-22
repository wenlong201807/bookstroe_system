package base

import "fmt"

func Range()  {
	// 遍历对象
	x := make(map[string]int)
	x["zhu"] = 66
	x["wen"] = 86
	x["long"] = 96

	for i,v := range x{
		fmt.Println(i,v)
	}

	// 遍历字符串
	y := "zhuwenlong"
	for _,v := range y{
		//fmt.Println("ascidd输出",v)
		fmt.Printf("%c",v)
	}
}