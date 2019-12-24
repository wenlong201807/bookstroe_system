package base

import (
	"regexp"
	"fmt"
)
// Golang中的正则表达式
 // https://my.oschina.net/kuerant/blog/199146
 // func Match(pattern string ,b []byte) (matched bool,err error)

func Regexp()  {
	isok,_ := regexp.Match("[a-zA-Z]{3}",[]byte("zhu"))
	fmt.Printf("%v",isok)
	fmt.Println("\n regexp.Match:",isok) //true

	isok2,_ := regexp.MatchString("[a-zA-Z]{3}","zh8u")
	fmt.Printf("%v",isok2)
	fmt.Println("\n regexp.MatchString:",isok2) //false

	// 正则表达式中的\ 符号需要转义  \\  => \
	reg := regexp.MustCompile(`^z.*1$`)
	fmt.Println("测试正则规则为：",reg)

	result := reg.FindAllString("zhuwenlong1",-1)
	fmt.Println("result的获取匹配结果：",result)

	reg1 := regexp.MustCompile(`^z(.*)1$`)
	result1 := reg1.FindAllString("zhuwenlong",-1)
	fmt.Println("result1的结果是：",result1)

	reg2 := regexp.MustCompile(`^z(.{1})(.{1})(.*)1$`)
	result2 := reg2.FindAllStringSubmatch("zhuwenlong",-1)
	fmt.Println("result2的结果是：",result2)
}
