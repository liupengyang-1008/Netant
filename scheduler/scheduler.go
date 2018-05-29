// scheduler
package scheduler

import (
	"Netant/analyzer"
	"Netant/itempipeline"
	"net/http"
)

//调度器的接口类型
type Scheduler interface {
	//启动调度器
	//调用该方法会是调度器创建和初始化各个组件。再次之后 调度器会激活爬取的流程
	//参数channellen 被用来制定数据传输通道的长度
	//参数 poolsize 用来设定网页下载池和分析池的容量
	//参数 crawldepth 代表了需要被爬取网页的最大深度。大于此深度的网页会被忽略
	//参数 httpClientGenerator 代表的是被用来生成http客户端的函数
	//参数 respParsers 的值应为分析器所需的被用来解析http响应函数的序列
	//参数 itemProcessors 的值应为需要被置入条目处理管道中的条目处理器的序列
	//参数 firstHttpReq 即代表首次请求 调度器会以此为起始点开始执行爬取流程
	Start(channelLen uint,
		poolSize uint32,
		crawlDepth uint32,
		httpClientGenerator GenHttpClient,
		respParsers []analyzer.ParseResponse,
		itemProcessors []itempipeline.ProcessItem,
		firstHttpReq *http.Request) (err error)

	//调用该方法会停止调度器的执行，所有处理模块的执行流程都会被终止
	Stop() bool

	//判定调度器是否正在运行
	Running() bool

	//获得错误通道，调度器以及各个处理模块运行过程中出现的所有错误都会被发送到错误通道
	//过该方法返回nil 则说明错误通道是不可用的或者调度器已被停止
	ErrorChan() <-chan error
	//判定所有处理模块是否都处于空闲状态
	Idle() bool
	//获得摘要信息
	Summary(prefix string) SchedSummary
}

//被用来生成http客户端的函数
type GenHttpClient func() *http.Client

//调度器摘要信息的接口类型
type SchedSummary interface {
	String() string               //获得摘要信息的一般表示
	Detail() string               // 获取摘要信息的详细表示
	Same(other SchedSummary) bool // 判定是否与另外一分摘要信息相同
}
