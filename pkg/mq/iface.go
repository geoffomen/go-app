package mq

import (
	"fmt"
	"sync"

	"github.com/geoffomen/go-app/pkg/mq/kafkaimp"
	"github.com/geoffomen/go-app/pkg/mq/pulsarimp"
)

type Iface interface {
	SendMsg(topic string, msg []byte) (msgID string, err error)
	ConsumeMsgByPull(topic, subsName string) (msgID string, content string, err error)
	ConsumeMsgByPush(topic string, handler func(content []byte) error) error
	AckMsgId(topic, subsName, msgId string) error
}

type Configuration struct {
	Url              string
	Token            string
	ListenerName     string
	Brokers          []string
	ProducerRetryMax int
	Version          string
	GroupName        string
	Assignor         string
	Oldest           bool
	Verbose          bool
	SaslUserName     string
	SaslUserPass     string
}

var (
	ins      Iface
	once     sync.Once
	isInited bool
)

func New(conf Configuration, logger pulsarimp.LogIface) Iface {
	return NewPulsar(conf, logger)
}

func NewPulsar(conf Configuration, logger pulsarimp.LogIface) *pulsarimp.PulsarService {
	var srv *pulsarimp.PulsarService
	once.Do(func() {
		if isInited {
			return
		}
		service, err := pulsarimp.New(pulsarimp.Configuration{
			Url:          conf.Url,
			Token:        conf.Token,
			ListenerName: conf.ListenerName,
		}, logger)
		if err != nil {
			panic(fmt.Sprintf("failed to initrialize config component, err: %v", err))
		}
		ins = service
		isInited = true
		srv = service
	})
	return srv
}

func NewKafka(conf Configuration, logger kafkaimp.LogIface) *kafkaimp.KafkaService {
	var srv *kafkaimp.KafkaService
	once.Do(func() {
		if isInited {
			return
		}
		service, err := kafkaimp.New(kafkaimp.KafkaConf{
			Brokers:          conf.Brokers,
			ProducerRetryMax: conf.ProducerRetryMax,
			Version:          conf.Version,
			GroupName:        conf.GroupName,
			Assignor:         conf.Assignor,
			Oldest:           conf.Oldest,
			SaslUserName:     conf.SaslUserName,
			SaslUserPass:     conf.SaslUserPass,
		}, logger)
		if err != nil {
			panic(fmt.Sprintf("failed to initrialize config component, err: %v", err))
		}
		ins = service
		isInited = true
		srv = service
	})
	return srv
}

func GetInstance() Iface {
	return ins
}
