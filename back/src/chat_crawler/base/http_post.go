package base

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
)

func Http_post()  {
	resp,err := http.Post("http://www.baidu.com","application/x-www-from-urlencoded",strings.NewReader("id=1"))
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
