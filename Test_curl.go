package main

/**
 * 测试curl请求
 */

import (
	"go_curl"
)

func main()  {
	params := make(map[string]string)
	params["age"] = "22"
	params["high"] = "180"
	url := "https://baike.baidu.com/item/PHP/9337?fr=aladdin"
	curl := go_curl.MakeCurl(url, "post")
	curl.SetQueryParam(params).Get()
}