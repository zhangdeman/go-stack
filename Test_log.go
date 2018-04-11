package main

import (
	"fmt"
	"self_log"
)
func main()  {
	fmt.Println("测试日志记录")
	log, _ := self_log.MakeLog("/share/go-lib/stack/log/", self_log.LogLevelDebug)
		log.WriteFatalStr("我的 ", "名字 ", " [张德满] ")

}
