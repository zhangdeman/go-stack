package go_curl

import (
	"strings"
	"net/url"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
)

/**
 * curl请求类库
 */

type Curl struct {
	url            string                 //请求地址
	headers        map[string]string      //header头
	cookie         map[string]string      //cookie信息
	method         string                 //请求方式 get/post等
	queryParams    map[string]string //请求参数
	postData       map[string]interface{} //post参数
	connectTimeout int                    //连接超时时间,单位 ： ms
	executeTimeout int                    //执行超时时间
	retryTimes     int                    //重试次数
	response		*http.Response				//响应结果
	responseData string	//响应数据
}

/**
 * 创建curl实例
 */
func MakeCurl(url string, method string) Curl {
	defaultMap := make(map[string]string)
	defaultConnectTimeout := 10
	defaultExecuteTimeout := 100
	defaultRetryTimes := 1
	curl := Curl{
		url:            url,
		headers:        defaultMap,
		cookie:         defaultMap,
		method:         strings.ToUpper(method),
		queryParams:defaultMap,
		postData:make(map[string]interface{}),
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
	curl.connectTimeout = connectTimeout
	return curl
}

/**
 * 设置执行超时时间
 */
func (curl *Curl) SetExecuteTimeout(executeTimeout int) *Curl {
	curl.executeTimeout = executeTimeout
	return curl
}

/**
 * 设置重试次数
 */
func (curl *Curl) SetRetryTimes(retryTimes int) *Curl {
	curl.retryTimes = retryTimes
	return curl
}

/**
 * 设置查询参数
 */
func (curl *Curl) SetQueryParam(queryMap map[string]string) *Curl {
	curl.queryParams = queryMap
	return curl
}

/**
 * 设置post参数
 */
func (curl *Curl) SetPostData(postData map[string]interface{}) *Curl {
	curl.postData = postData
	return curl
}

/**
 * 设置请求header头
 */
func (curl *Curl) SetRequestHeader(headerName string, headerVal string) *Curl {
	curl.headers[headerName] = headerVal
	return curl
}

/**
 * 设置请求cookie信息
 */
func (curl *Curl) SetRequestCookie(cookieName string, cookieVal string) *Curl {
	curl.headers[cookieName] = cookieVal
	return curl
}

/**
 * 发送请求
 */

func (curl *Curl) Request() *Curl {
	switch {
	case curl.method == RequestPost:
		return curl.Post()
	default:

	}
	return curl.DefaultCurl()
}

func (curl *Curl) DefaultCurl() *Curl {
	return curl
}

func (curl *Curl) Post() *Curl {
	return curl
}

/**
 * 解析get请求查询字符串,并组装成url
 */
func (curl *Curl) ParseQueryUrl() string {
	u, _ := url.Parse(curl.url)
	uriStr := "&"+u.RequestURI()	//附在url末尾的查询串
	for k,v := range curl.queryParams{
		uriStr = uriStr + "&" + k + "=" + v
	}
	scheme := u.Scheme
	host := u.Host
	return scheme+"://"+host+strings.Trim(uriStr, "&")
}

/**
 * 设置响应数据
 */
func (curl *Curl) setResponseData(closer io.ReadCloser) *Curl {
	responseData, _ := ioutil.ReadAll(closer)
	curl.responseData = string(responseData)
	return curl
}

/**
 * 发送get请求
 */
func (curl *Curl) Get() *Curl {
	requestUrl := curl.ParseQueryUrl()
	resp, err := http.Get(requestUrl)
	if err != nil {
		return curl
	}
	curl.response = resp
	curl.setResponseData(resp.Body)
	defer resp.Body.Close()
	fmt.Println(requestUrl)
	return curl
}

func (curl *Curl) Head() *Curl {
	return curl
}

/**
 * 获取响应http状态码
 */
func (curl *Curl) GetResponseCode() int {
	return curl.response.StatusCode
}

/**
 * 获取http响应状态
 */
func (curl *Curl) GetResponseStatus() string {
	return curl.response.Status
}
