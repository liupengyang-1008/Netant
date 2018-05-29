// middleware
package middleware

import (
	"Netant/base"
	"errors"
	"fmt"
	"sync"
)

var defaultChanlen uint = 10

//-----------------------------------------------------
//通道管理器的接口类型
type ChannelManager interface {
	//初始化
	//channelLen 通道管理器中各类通道的初始长度
	//reset是否初始化通道管理器
	Init(channelLen uint, reset bool) bool

	//关闭通道管理器
	Close() bool
	//获取请求传输通道
	ReqChan() (chan base.Request, error)

	//获取响应传输通道
	RespChan() (chan base.Response, error)
	//获取条目传输通道
	ItemChan() (chan base.Item, error)
	//获取错误传输通道
	ErrorChan() (chan error, error)

	//获取通道长度
	ChannelLen() uint
	//获取通道管理器的状态
	Status() ChannelManagerStatus
	//获取摘要信息
	Summary() string
}

//通道管理器的状态
type ChannelManagerStatus uint8

const (
	CHANNEL_MANAGER_STATUS_UNINITIALIZED ChannelManagerStatus = 0
	CHANNEL_MANAGER_STATUS_INITIALIZED   ChannelManagerStatus = 1
	CHANNEL_MANAGER_STATUS_CLOSED        ChannelManagerStatus = 2
)

type myChannelManager struct {
	channelLen uint
	reqCh      chan base.Request
	respCh     chan base.Response
	itemCh     chan base.Item
	errorCh    chan error
	status     ChannelManagerStatus
	rwmutex    sync.RWMutex
}

//创建通道管理器
func NewChannelManager(channelLen uint) ChannelManager {
	if channelLen == 0 {
		channelLen = defaultChanlen
	}
	chanman := &myChannelManager{}
	chanman.Init(channelLen, true)
	return chanman
}

func (chanman *myChannelManager) Init(channelLen uint, reset bool) bool {
	if channelLen == 0 {
		panic(errors.New("The channel length is Invalid"))
	}
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if chanman.status == CHANNEL_MANAGER_STATUS_INITIALIZED && !reset {
		return false
	}
	chanman.channelLen = channelLen
	chanman.reqCh = make(chan base.Request, channelLen)

	chanman.respCh = make(chan base.Response, channelLen)

	chanman.itemCh = make(chan base.Item, channelLen)

	chanman.errorCh = make(chan error, channelLen)

	chanman.status = CHANNEL_MANAGER_STATUS_INITIALIZED
	return true
}

func (chanman *myChannelManager) Close() bool {
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if chanman.status != CHANNEL_MANAGER_STATUS_INITIALIZED {
		return false
	}

	close(chanman.reqCh)
	close(chanman.respCh)
	close(chanman.itemCh)
	close(chanman.errorCh)
	chanman.status = CHANNEL_MANAGER_STATUS_CLOSED
	return true
}

//检查状态
func (chanman *myChannelManager) checkStatus() error {
	if chanman.status == CHANNEL_MANAGER_STATUS_INITIALIZED {
		return nil
	}

	statusName, ok := statusNameMap[chanman.status]
	if !ok {
		statusName = fmt.Sprintf("%d", chanman.status)
	}
	errMsg := fmt.Sprintf("The undesirable status of channel manager:%s! ", statusName)
	return errors.New(errMsg)
}

//字典
var statusNameMap = map[ChannelManagerStatus]string{
	CHANNEL_MANAGER_STATUS_CLOSED:        "closed",
	CHANNEL_MANAGER_STATUS_INITIALIZED:   "initialized",
	CHANNEL_MANAGER_STATUS_UNINITIALIZED: "uninitialized",
}

func (chanman *myChannelManager) ReqChan() (chan base.Request, error) {
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.reqCh, nil
}

func (chanman *myChannelManager) RespChan() (chan base.Response, error) {
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.respCh, nil
}

func (chanman *myChannelManager) ErrorChan() (chan error, error) {
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.errorCh, nil
}

func (chanman *myChannelManager) ItemChan() (chan base.Item, error) {
	chanman.rwmutex.Lock()
	defer chanman.rwmutex.Unlock()
	if err := chanman.checkStatus(); err != nil {
		return nil, err
	}
	return chanman.itemCh, nil
}

var chanmanSummaryTemplate = "status:%s, " +
	"requestChannel:%d/%d, " +
	"responseChannel:%d/%d, " +
	"itemChannel:%d/%d, " +
	"errorChannel:%d/%d, "

func (chanman *myChannelManager) Summary() string {
	summary := fmt.Sprintf(chanmanSummaryTemplate,
		statusNameMap[chanman.status],
		len(chanman.reqCh), cap(chanman.reqCh),
		len(chanman.respCh), cap(chanman.respCh),
		len(chanman.itemCh), cap(chanman.itemCh),
		len(chanman.errorCh), cap(chanman.errorCh))
	return summary
}

func (chanman *myChannelManager) Status() ChannelManagerStatus {
	return chanman.status
}

func (chanman *myChannelManager) ChannelLen() uint {
	return chanman.channelLen
}
