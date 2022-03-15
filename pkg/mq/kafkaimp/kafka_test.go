package kafkaimp

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestKafka(t *testing.T) {
	mq, _ := New(KafkaConf{
		Brokers:          []string{""},
		ProducerRetryMax: 3,
		Version:          "",
		GroupName:        "",
		Assignor:         "",
		Oldest:           true,
		SaslUserName:     "",
		SaslUserPass:     "",
	}, log.Default())
	go sendMsgT(mq)
	time.Sleep(3 * time.Second)
	rcvMsg1(mq)
	rcvMsg2(mq)
	time.Sleep(1000 * time.Second)
}

func sendMsgT(ins *KafkaService) {
	fmt.Println("begin produce msg ====================")
	for {
		time.Sleep(1 * time.Second)
		ins.SendMsg("testTopic", []byte("hello!!!"+time.Now().String()))
	}
}
func rcvMsg1(ins *KafkaService) {
	ins.ConsumeMsgByPush("testTopic", func(content []byte) error {
		fmt.Println(string(content))
		return nil
	})
}

func rcvMsg2(ins *KafkaService) {
	ins.ConsumeMsgByPush("testTopic2", func(content []byte) error {
		fmt.Print(string(content))
		return nil
	})
}
