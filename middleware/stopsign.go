// stopsign
package middleware

//-----------------------------------------------------
//停止信号的接口类型
type StopSign interface {
	//设置停止信号
	//如果已经停止返回false
	Sign() bool

	//判定信号是否已经发出
	Signed() bool
	//reset
	Reset()

	//处理停止信息 code表示处理停止信号的方法
	Deal(code string)
	//获得某一停止信号的处理方的计数
	DealCount(code string) uint32
	//获得停止信号的被处理的总计数
	DealTotal() uint32

	//获取摘要消息
	Summary() string
}
