package mq

import (
	"fmt"
	"sync"

	"github.com/geoffomen/go-app/internal/pkg/config"
	"github.com/geoffomen/go-app/internal/pkg/mq/pulsarimp"
)

type Iface interface {
	SendMsg(topic string, msg []byte) (msgID string, err error)
	ConsumeMsgByPull(topic, subsName string) (msgID string, content string, err error)
	ConsumeMsgByPush(topic string, handler func(content []byte) error) error
	AckMsgId(topic, subsName, msgId string) error
}

var (
	ins      Iface
	once     sync.Once
	isInited bool
)

func New(cf config.Iface) Iface {
	once.Do(func() {
		if isInited {
			return
		}
		service, err := pulsarimp.New(pulsarimp.Configuration{
			Url:          cf.GetStringOrDefault("", ""),
			Token:        cf.GetStringOrDefault("", ""),
			ListenerName: cf.GetStringOrDefault("", ""),
		})
		if err != nil {
			panic(fmt.Sprintf("failed to initrialize config component, err: %v", err))
		}
		ins = service
		isInited = true
	})
	return ins
}

func GetInstance() Iface {
	return ins
}

func RegisterConsumer(topicAndHandler map[string]func(content []byte) error) {
	for topic, handler := range topicAndHandler {
		ConsumeMsgByPush(topic, handler)
	}
}

func SendMsg(topic string, msg []byte) (string, error) {
	return ins.SendMsg(topic, msg)
}

func ConsumeMsgByPull(topic, subsName string) (string, string, error) {
	return ins.ConsumeMsgByPull(topic, subsName)
}
func ConsumeMsgByPush(topic string, handler func(content []byte) error) error {
	return ins.ConsumeMsgByPush(topic, handler)
}

func AckMsgId(topic, subsName, msgId string) error {
	return ins.AckMsgId(topic, subsName, msgId)
}
