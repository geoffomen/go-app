package mq

type Iface interface {
	SendMsg(topic string, msg []byte) (msgID string, err error)
	ConsumeMsgByPull(topic, subsName string) (msgID string, content string, err error)
	ConsumeMsgByPush(topic string, handler func(content []byte) error) error
	AckMsgId(topic, subsName, msgId string) error
}

var (
	ins Iface
)

func SetInstance(i Iface) {
	ins = i
}

// GetInstance ..
func GetInstance() Iface {
	return ins
}
