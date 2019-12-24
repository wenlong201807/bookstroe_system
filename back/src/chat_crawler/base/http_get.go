package base

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Http_get()  {

	resp,err := http.Get("http://www.baidu.com")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))


}
