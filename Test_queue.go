package main

import (
	"fmt"
	"stack"
)

func main()  {
	queue := stack.MakeQueue(32, true)
	queue.Push("zhangdeman")
	queue.Push("qwe")
	queue.Push("wer")
	for  {
		data, err := queue.Pop()
		if err != nil {
			fmt.Println("队列数据为空")
			break
		}
		fmt.Println("data ",data)
	}

	queue.Push("123")
	queue.Push("234")
	queue.Push("345")
	for  {
		data, err := queue.Pop()
		if err != nil {
			fmt.Println("队列数据为空")
			break
		}
		fmt.Println("data ",data)
	}
	queue.Push("q")
	queue.Push("w")
	queue.Push("e")
	data, _ := queue.Pop()
	fmt.Println("data ",data)
	queue.Push("a")
	queue.Push("s")
	queue.Push("d")
	for  {
		data, err := queue.Pop()
		if err != nil {
			fmt.Println("队列数据为空")
			break
		}
		fmt.Println("data ",data)
	}
}