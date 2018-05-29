// Analyzer
package analyzer

import (
	"Netant/base"
	"net/http"
)

//分析器的接口类型
type Analyzer interface {
	Id()
	Analyze(
		respParsers []ParseResponse,
		resp base.Response) ([]base.Data, []error) //根据规则分析响应并返回请求和条目
}

type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)

//分析器池的接口类型
type AnalyzerPool interface {
	Take() (Analyzer, error)        //从池中取出一个分析器
	Return(analyzer Analyzer) error // 把一个分析器放入池中
	Total() uint32                  // 获得池的总量
	Used() uint32                   // 获得正在使用的分析器的数量
}
