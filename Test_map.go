package main

import (
	"hashmap"
	"fmt"
)
func main()  {
	hashMap := hashmap.MakeHashMap()
	hashMap.Add("a", "zhang")
	hashMap.Add("b", "zhang")
	hashMap.Add("c", "zhang")
	fmt.Println(hashMap.GetAllData())
	hashMap.Add("a", "man")
	hashMap.Add("b", "de")
	fmt.Println(hashMap.GetAllData())

	fmt.Println(hashMap.GetMapValue("b"))
	fmt.Println(hashMap.BatchDelete("a","b","c"))
}