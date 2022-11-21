package session

// Session 结构
type Session interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
}
