// middleware
package middleware

//Id生成器
type IdGenerator interface {
	GetUint32() uint32
}
