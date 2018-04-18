package main

import (
	"server"
	"test"
)
func main()  {
	defaultMap := []string{}
	server.MakeServer("http", defaultMap,defaultMap, "9009")
	server.AddUriMap("/test", test.TestMethod)
	server.RunServer()
}