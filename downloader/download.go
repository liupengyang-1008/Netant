// crawinterface
package downloader

//Id生成器的接口类型
type PageDownLoader interface {
	Id() uint32 //获得id
	Download(req base.Request) (*base.Reponse, error)
}



//网页下载器的接口类型
type PageDownLoaderPool interface {
	Take() (PageDownLoader, error)  //从池中获取一个下载器
	Return(dl PageDownLoader) error // 把一个网页下载器归还给池
	Rotal() uint32                  //获取池的总量
	Used() uint32                   //正在被使用的下载器总量
}
