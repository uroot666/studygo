package session

import "fmt"

var (
	sessionMgr SessionMgr
)

// 中间件让客户选择使用哪个版本
func Init(provider string, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()

	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持")
		return
	}
	err = sessionMgr.Init(addr, options...)
	return
}
