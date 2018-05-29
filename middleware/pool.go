// middleware
package middleware

//-----------------------------------------------------
//实体池
//实体池的接口类型
type Pool interface {
	Take() (Entity, error)      //从池中取出一个实体
	Return(entity Entity) error // 把一个实体放入池中
	Total() uint32              // 获得池的总量
	Used() uint32               // 获得实体池正在使用的实体的数量
}

type Entity interface {
	Id() uint32 //id的获取方法
}
