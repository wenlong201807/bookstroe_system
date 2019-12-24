package base


import (
	json "encoding/json"
	"fmt"
)

/*
 JSON编码
    func Marshal(v interface{} ) ([]byte,error)

 JSON解码
    func Unmarshal(data []byte,v interface{} ) error
*/

type Student3 struct {
	Name string `json:"student_name"` // 转成json格式的数据之后，字段变为student_name  而不是Name
	Age int                           // {"student_name":"wen","Age":23}
}


func Json()  {

	// 对数组类型的json编码
	x  := [5]int{1,2,3,4,5}
	s,err := json.Marshal(x)
	if err != nil{
		panic(err)
	}
	fmt.Println("ascidd编码：",s)
	fmt.Println("字符串内容：",string(s)) //[1,2,3,4,5]

	// 对map类型的json编码
	m2  := make(map[string]float64)
	m2["zhu"] = 24.6
	s2,err2 := json.Marshal(m2)
	if err2 != nil{
		panic(err2)
	}
	fmt.Println("ascidd编码：",s2)
	fmt.Println("字符串内容：",string(s2)) //{"zhu":24.6}

	// 对对象类型的json编码
	student := Student3{"wen",23}
	s3,err3 := json.Marshal(student)
	if err3 != nil{
		panic(err3)
	}
	fmt.Println("ascidd编码：",s3)
	fmt.Println("字符串内容：",string(s3)) //{"Name":"wen","Age":23}
	//  可以自定义字段写法 {"student_name":"wen","Age":23}

	// 对s3进行解码
	var s4 interface{}
	json.Unmarshal([]byte(s3),&s4)
	fmt.Printf("%v",s4) // map[Age:23 student_name:wen]

}
