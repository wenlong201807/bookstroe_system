package base

import (
	"net/http"
	"fmt"
)

func HttpServe()  {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("打开浏览器，在地址栏输入http://www.localhost:8089/hello,就可以看见我啦！"))
	})

	http.ListenAndServe("127.0.0.1:8089",nil)
	fmt.Println("启动服务啦 -->> http://www.localhost:8089/hello")
}
