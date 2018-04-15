package main

/**
 * 测试curl请求
 */

import (
	"go_curl"
	"fmt"
	"reflect"
)

func main()  {
	params := make(map[string]string)
	params["age"] = "22"
	params["high"] = "180"
	url := "https://baike.baidu.com/item/PHP/9337?fr=aladdin"
	go_curl.MakeCurl(url, "get")
	go_curl.SetQueryParam(params)
	fmt.Println(go_curl.Request())
	fmt.Println(params, reflect.TypeOf(go_curl.CurlInstance))
}