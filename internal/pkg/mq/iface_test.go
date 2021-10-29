package mq

import (
	"fmt"
	"testing"
	"time"

	"github.com/geoffomen/go-app/internal/pkg/config"
)

func TestPulsar(t *testing.T) {
	cf := config.NewEmpty()
	cf.Set("url", "localhost:6650")
	New(cf)
	go sendMsgT()
	go recMsgT2()
	recMsgT()
}

func TestKafka(t *testing.T) {
	cf := config.NewEmpty()
	cf.Set("kafka.brokers", "localhost:9092")
	cf.Set("kafka.producerRetryMax", 1)
	cf.Set("kafka.version", "1.1.1")
	cf.Set("kafka.group", "localhost")
	cf.Set("kafka.assignor", "range")
	cf.Set("kafka.oldest", true)
	cf.Set("kafka.verbose", true)
	New(cf)
	go sendMsgT()
	time.Sleep(3 * time.Second)
	rcvMsgPushT()
	rcvMsgPushT2()
	time.Sleep(1000 * time.Second)
}

func sendMsgT() {
	fmt.Println("begin produce msg ====================")
	for {
		time.Sleep(1 * time.Second)
		SendMsg("testTopic", []byte("hello!!!"+time.Now().String()))
	}
}

func rcvMsgPushT() {
	ConsumeMsgByPush("testTopic", func(content []byte) error {
		fmt.Println(string(content))
		return nil
	})
}

func rcvMsgPushT2() {
	ConsumeMsgByPush("testTopic2", func(content []byte) error {
		fmt.Print(string(content))
		return nil
	})
}

func recMsgT() {
	fmt.Println("begin receive msg ====================")
	for {
		msgID, msg, _ := ConsumeMsgByPull("testTopic", "subsName")
		fmt.Printf("receive msg: %s, %s, %s", "subsName", msgID, msg)
		AckMsgId("testTopic", "subsName", msgID)
		time.Sleep(2 * time.Second)
	}
}

func recMsgT2() {
	fmt.Println("begin receive msg ====================")
	for i := 0; i < 100; i++ {
		msgID, msg, _ := ConsumeMsgByPull("testTopic", "subsName2")
		fmt.Printf("receive msg: %s, %s, %s", "subsName2", msgID, msg)
		AckMsgId("testTopic", "subsName", msgID)
		time.Sleep(3 * time.Second)
	}
}
