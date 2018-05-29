// itempipeline
package itempipeline

import (
	"Netant/base"
)

//条目处理的管道的接口类型
type ItemPipeLine interface {
	//发送条目
	Send(item base.Item) []error
	//Failfast 方法返回一个布尔值。改值表示当前条目处理管道是否是快速失败的
	//这里的快速失败是指：只要对某个条目的处理流程在某个步骤上出错
	//那么条目的处理管道就会忽略掉后续所有的处理步骤并返回错误
	FailFase() bool
	//设置是否快速失败
	SetFailFast(failfast bool)

	//获得已发送 已接收 和已处理的条目计数值
	ProcessingNumber() uint64

	//获取摘要信息
	Summary() string
}

//被用来处理条目的函数类型
type ProcessItem func(item base.Item) (result base.Item, err error)
