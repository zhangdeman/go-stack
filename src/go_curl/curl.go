package go_curl

import (
	"strings"
	"net/url"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"errors"
)

/**
 * curl接口
 */
type CurlInterface interface {
	MakeCurl(url string, method string) *Curl
	SetConnectTimeout(connectTimeout int) *Curl
	SetExecuteTimeout(executeTimeout int) *Curl
	SetRetryTimes(retryTimes int) *Curl
	SetQueryParam(data map[string]string) *Curl
	SetPostData(data map[string]interface{}) *Curl
	SetRequestHeader(headerName string, headerVal string) *Curl
	SetRequestCookie(cookieName string, cookieVal string) *Curl
	Request() (CurlResponse, error)
	DefaultCurl() (CurlResponse, error)
	Post() (CurlResponse, error)
	ParseQueryUrl() string
	SetResponseData(closer io.ReadCloser) *Curl
	Get() (CurlResponse, error)
	Head() (CurlResponse, error)
	GetResponseCode() int
	GetResponseStatus() string
}

/**
 * curl请求类库
 */

type CurlConfig struct {
	url            string                 //请求地址
	headers        map[string]string      //header头
	cookie         map[string]string      //cookie信息
	method         string                 //请求方式 get/post等
	queryParams    map[string]string      //请求参数
	postData       map[string]interface{} //post参数
	connectTimeout int                    //连接超时时间,单位 ： ms
	executeTimeout int                    //执行超时时间
	retryTimes     int                    //重试次数
	response       *http.Response         //响应结果
	responseData   string                 //响应数据
}

type CurlResponse struct {
	resp *http.Response
	data string
	err error
}

type Curl struct {
}

var CurlInstance Curl
var CurlConfigInstance CurlConfig

/**
 * 构造函数
 */
func init() {

}

/**
 * 创建curl实例
 */
func (curl *Curl) MakeCurl(url string, method string) *Curl {
	defaultMap := make(map[string]string)
	defaultConnectTimeout := 10
	defaultExecuteTimeout := 100
	defaultRetryTimes := 1
	CurlConfigInstance = CurlConfig{
		url:            url,
		headers:        defaultMap,
		cookie:         defaultMap,
		method:         strings.ToUpper(method),
		queryParams:    defaultMap,
		postData:       make(map[string]interface{}),
		connectTimeout: defaultConnectTimeout,
		executeTimeout: defaultExecuteTimeout,
		retryTimes:     defaultRetryTimes,
	}
	return curl
}

/**
 * 设置连接超时时间
 */
func (curl *Curl) SetConnectTimeout(connectTimeout int) *Curl {
	CurlConfigInstance.connectTimeout = connectTimeout
	return curl
}

/**
 * 设置执行超时时间
 */
func (curl *Curl) SetExecuteTimeout(executeTimeout int) *Curl {
	CurlConfigInstance.executeTimeout = executeTimeout
	return curl
}

/**
 * 设置重试次数
 */
func (curl *Curl) SetRetryTimes(retryTimes int) *Curl {
	CurlConfigInstance.retryTimes = retryTimes
	return curl
}

/**
 * 设置查询参数
 */
func (curl *Curl) SetQueryParam(queryMap map[string]string) *Curl {
	CurlConfigInstance.queryParams = queryMap
	return curl
}

/**
 * 设置post参数
 */
func (curl *Curl) SetPostData(postData map[string]interface{}) *Curl {
	CurlConfigInstance.postData = postData
	return curl
}

/**
 * 设置请求header头
 */
func (curl *Curl) SetRequestHeader(headerName string, headerVal string) *Curl {
	CurlConfigInstance.headers[headerName] = headerVal
	return curl
}

/**
 * 设置请求cookie信息
 */
func (curl *Curl) SetRequestCookie(cookieName string, cookieVal string) *Curl {
	CurlConfigInstance.headers[cookieName] = cookieVal
	return curl
}

/**
 * 发送请求
 */

func (curl *Curl) Request() (CurlResponse,error) {
	switch {
	case CurlConfigInstance.method == RequestPost:
		return curl.Post()
	case CurlConfigInstance.method == RequestGet:
		return curl.Get()
	default:

	}
	return curl.DefaultCurl()
}

func (curl *Curl) DefaultCurl() (CurlResponse, error) {
	return CurlResponse{}, nil
}

func (curl *Curl) Post() (CurlResponse, error) {
	return CurlResponse{}, nil
}

/**
 * 解析get请求查询字符串,并组装成url
 */
func (curl *Curl) ParseQueryUrl() string {
	u, _ := url.Parse(CurlConfigInstance.url)
	uriStr := "&" + u.RequestURI() //附在url末尾的查询串
	for k, v := range CurlConfigInstance.queryParams {
		uriStr = uriStr + "&" + k + "=" + v
	}
	scheme := u.Scheme
	host := u.Host
	return scheme + "://" + host + strings.Trim(uriStr, "&")
}

/**
 * 设置响应数据
 */
func (curl *Curl) SetResponseData(closer io.ReadCloser) *Curl {
	responseData, _ := ioutil.ReadAll(closer)
	CurlConfigInstance.responseData = string(responseData)
	return curl
}

/**
 * 发送get请求
 */
func (curl *Curl) Get() (CurlResponse, error) {
	requestUrl := curl.ParseQueryUrl()
	resp, err := http.Get(requestUrl)
	if err != nil {
		return CurlResponse{}, errors.New("获取请求结果失败")
	}
	CurlConfigInstance.response = resp
	responseData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(requestUrl,resp)
	response := CurlResponse{
		resp:resp,
		data:string(responseData),
		err:nil,
	}
	return response, nil
}

func (curl *Curl) Head() (CurlResponse, error) {
	return CurlResponse{}, nil
}

/**
 * 获取响应http状态码
 */
func (curl *Curl) GetResponseCode() int {
	return CurlConfigInstance.response.StatusCode
}

/**
 * 获取http响应状态
 */
func (curl *Curl) GetResponseStatus() string {
	return CurlConfigInstance.response.Status
}

func MakeCurl(url string, method string) *Curl {
	return CurlInstance.MakeCurl(url, method)
}

func SetConnectTimeout(connectTimeout int) *Curl {
	return CurlInstance.SetConnectTimeout(connectTimeout)
}

func SetExecuteTimeout(executeTimeout int) *Curl {
	return CurlInstance.SetExecuteTimeout(executeTimeout)
}

func SetRetryTimes(retryTimes int) *Curl {
	return CurlInstance.SetRetryTimes(retryTimes)
}

func SetQueryParam(data map[string]string) *Curl {
	return CurlInstance.SetQueryParam(data)
}

func SetPostData(data map[string]interface{}) *Curl {
	return CurlInstance.SetPostData(data)
}

func SetRequestHeader(headerName string, headerVal string) *Curl {
	return SetRequestHeader(headerName, headerVal)
}

func SetRequestCookie(cookieName string, cookieVal string) *Curl {
	return CurlInstance.SetRequestCookie(cookieName, cookieVal)
}

func Request() (CurlResponse, error) {
	return CurlInstance.Request()
}

func DefaultCurl() (CurlResponse, error) {
	return CurlInstance.DefaultCurl()
}

func Post() (CurlResponse, error) {
	return CurlInstance.Post()
}

func ParseQueryUrl() string {
	return CurlInstance.ParseQueryUrl()
}

func SetResponseData(closer io.ReadCloser) *Curl {
	return CurlInstance.SetResponseData(closer)
}

func Get() (CurlResponse, error) {
	return CurlInstance.Get()
}

func Head() (CurlResponse, error) {
	return CurlInstance.Head()
}

func GetResponseCode() int {
	return GetResponseCode()
}

func GetResponseStatus() string {
	return GetResponseStatus()
}
