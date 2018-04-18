package server

import (
	"fmt"
	"net/http"
)

/**
 * 创建服务器，监听http请求
 */

var (
	NewServerInstance       NewServer
	NewServerConfigInstance NewServerConfig
)

func init() {

}

//config
type NewServerConfig struct {
	scheme      string            //请求协议
	allowIpList []string          //允许请求的ip列表
	allowMethod []string          //允许的请求方法
	listenPort  string            //监听的端口
	funcMap     map[string]func(w http.ResponseWriter, r *http.Request) //监听的方法
}

/**
 * 接口
 */
type NewServerInterface interface {
	MakeServer(scheme string, allowIpList []string, allowMethod []string, listenPort string)	//创建一个服务器
	RunServer()	//运行服务器
	AddUriMap(uri string, dealFunc func(w http.ResponseWriter, r *http.Request))	//增加一个请求map
}

/**
 * 实现 NewServerInterface 接口的结构体
 */
type NewServer struct {
}

/**
 * 创建sever
 */
func (newServer *NewServer) MakeServer(scheme string, allowIpList []string, allowMethod []string, listenPort string) {
	NewServerConfigInstance.scheme = scheme
	NewServerConfigInstance.allowIpList = allowIpList
	NewServerConfigInstance.allowMethod = allowMethod
	NewServerConfigInstance.listenPort = ":"+listenPort
	mapFunc := make(map[string]func(w http.ResponseWriter, r *http.Request))
	NewServerConfigInstance.funcMap = mapFunc
}

/**
 * 增加新的请求map
 */
func (newServer *NewServer) AddUriMap(uri string, dealFunc func(w http.ResponseWriter, r *http.Request))  {
	NewServerConfigInstance.funcMap[uri] = dealFunc
}

/**
 * 运行server
 */
func (newServer *NewServer) RunServer() {
	fmt.Println("服务器监听端口 " + NewServerConfigInstance.listenPort)
	for uri, method := range NewServerConfigInstance.funcMap{
		fmt.Println("注册请求 : " + uri )
		http.HandleFunc(uri, method)
	}
	err := http.ListenAndServe(NewServerConfigInstance.listenPort, nil)
	if err != nil {
		fmt.Println(err)
	}

}

/**
 * 创建服务器
 */
func MakeServer(scheme string, allowIpList []string, allowMethod []string, listenPort string) {
	NewServerInstance.MakeServer(scheme, allowIpList, allowMethod, listenPort)
}

/**
 * 运行服务器
 */
func RunServer()  {
	NewServerInstance.RunServer()
}

/**
 * 增加请求map
 */
func AddUriMap(uri string, dealFunc func(w http.ResponseWriter, r *http.Request)) {
	NewServerInstance.AddUriMap(uri, dealFunc)
}
