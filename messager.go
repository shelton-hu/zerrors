package zerrors

var msger Messager = new(defaultMessage)

type Messager interface {
	// 通过code获取message
	ErrorMessage(code int) string
	// 默认的code和对应的message
	DefaultMessage() (code int, message string)
}

func SetMessager(m Messager) bool {
	msger = m
	return true
}

type defaultMessage struct{}

func (d *defaultMessage) ErrorMessage(_ int) string {
	return "System exception"
}

func (d *defaultMessage) DefaultMessage() (int, string) {
	return 1000000, "System exception"
}
