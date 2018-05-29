// datatype
package base

import (
	"bytes"
	"fmt"
	"net/http"
)

//请求
type Request struct {
	httpReq *http.Request
	depth   uint32
}

//创建新请求
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

//获取http请求
func (req *Request) HttpReq() *http.Request {
	return req.httpReq
}

//获取深度值
func (req *Request) Depth() uint32 {
	return req.depth
}

//响应
type Response struct {
	httpResp *http.Response
	depth    uint32
}

//创建新的响应
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

//获取http响应
func (resp *Response) HttpResp() *http.Response {
	return resp.httpResp
}

//获取深度值
func (resp *Response) Depth() uint32 {
	return resp.depth
}

//条目
type Item map[string]interface{}

//数据的接口
type Data interface {
	Vaild() bool //数据是否有效
}

//响应是否有效
func (resp *Response) Vaild() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}

// 请求是否有效
func (req *Request) Vaild() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}

//条目是否有效
func (item Item) Vaild() bool {
	return item != nil
}

type ErrorType string

type CrawlerError interface {
	Type() ErrorType //获得错误类型
	Error() string   //获得错误提示信息

}

type myCrawError struct {
	errType    ErrorType //错误类型
	errMsg     string    //错误提示信息
	fullErrMsg string    //完整的错误提示信息
}

//错误类型常量
const (
	DOWNLOADER_ERROR     ErrorType = "Download Error"
	ANALYZER_ERROR       ErrorType = "Analyzer Error"
	ITEM_PROCESSER_ERROR ErrorType = "Item Process Error"
)

// 创建一个新的爬虫错误
func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawError{errType: errType, errMsg: errMsg}
}

//获得错误类型
func (ce *myCrawError) Type() ErrorType {
	return ce.errType
}

//获得错误提示信息
func (ce *myCrawError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

//生成错误提示信息，并给响应的字段赋值
func (ce *myCrawError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("Crawler Error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")

	}

	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s\n ", buffer.String())
	return
}
