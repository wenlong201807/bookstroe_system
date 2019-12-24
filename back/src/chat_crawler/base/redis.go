package base

// 源码 https://github.com/astaxie/goredis/blob/master/redis.go
import (
	"github.com/monnand/goredis"
	"fmt"
)

// 使用之前，本地的redis需要先启动，类似mysql
func Redis()  {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("test",[]byte("hello zhu"))
	if err != nil{
		panic(err)
	}
	res,err := client.Get("test")
	if err != nil{
		panic(err)
	}
	fmt.Println("获取test的内容：",string(res))

	// test hmset
	f := make(map[string]interface{})
	f["name"] = "zhangsan"
	f["age"] = 12
	f["sex"] = "male"
	err = client.Hmset("test_hash",f)
	if err != nil{
		panic(err)
	}

	// test zset
	_,err = client.Zadd("test_zset",[]byte("zhu"),100)
	if err !=nil{
		panic(err)
	}


}

 // client ===> get test  =>hello zhu