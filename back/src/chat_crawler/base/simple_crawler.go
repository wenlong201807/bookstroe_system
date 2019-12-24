package base

import (
	"fmt"
	"regexp"
	"net/http"
	"io/ioutil"
)

func SimpleCrawler() {
	fmt.Println("666")
	url := "https://movie.douban.com/subject/24751763/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	sHtml3, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(sHtml3)) // 转换成字符串的方法

	// 必须是字符串
	sHtml := `<span property="v:itemreviewed">罗曼提克消亡史</span>`

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	result := reg.FindAllStringSubmatch(sHtml,-1)
	fmt.Println("result:",result) //[[<span property="v:itemreviewed">罗曼提克消亡史</span> 罗曼提克消亡史]]
	for _,v := range result{
		fmt.Println("获取需要的内容：",v[1]) // 罗曼提克消亡史
	}
	fmt.Println("一次性获取内容：",result[0][1])

	// 必须是字符串
	sHtml2 := `<strong class="kk  bb" property="v:average">66.7</strong>`

	reg2 := regexp.MustCompile(`<strong\s*class="kk\s*bb"\s*property="v:average">(.*)</strong>`)
	result2 := reg2.FindAllStringSubmatch(sHtml2,-1)
	fmt.Println("result2:",result2) //[[<span property="v:itemreviewed">罗曼提克消亡史</span> 罗曼提克消亡史]]
	for _,v := range result2{
		fmt.Println("获取需要的内容：",v[1]) // 罗曼提克消亡史
	}
	fmt.Println("一次性获取内容：",result2[0][1])

	////resp,err := http.Get("http://www.baidu.com")
	//resp,err := http.Get("https://movie.douban.com/subject/24751763/")
	//if err != nil{
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body,err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// https://movie.douban.com/subject/24751763/
}
