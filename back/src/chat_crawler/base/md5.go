package base

import (
	"crypto/md5"
	"fmt"
)

func Md5()  {
	// 不可逆转的方式加密
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("zhuWenLong"))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%v",string(Result)) // 加密之后的内容： �hڝ��.�H�͈jz
	fmt.Println("\n 常规加密方式：")
	fmt.Printf("%x\n\n",Result)  // 加密之后的内容：  e168da9d86a42e14de48ac06cd886a7a


	// 可以相互转换的加密与解密方式
	// 需自行补充...
}