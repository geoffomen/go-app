package mq

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestPulsar(t *testing.T) {
	NewPulsar(Configuration{}, log.Default())
	go sendMsgT()
	go recMsgT2()
	recMsgT()
}

func TestKafka(t *testing.T) {
	NewKafka(Configuration{}, log.Default())
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
		GetInstance().SendMsg("testTopic", []byte("hello!!!"+time.Now().String()))
	}
}

func rcvMsgPushT() {
	GetInstance().ConsumeMsgByPush("testTopic", func(content []byte) error {
		fmt.Println(string(content))
		return nil
	})
}

func rcvMsgPushT2() {
	GetInstance().ConsumeMsgByPush("testTopic2", func(content []byte) error {
		fmt.Print(string(content))
		return nil
	})
}

func recMsgT() {
	fmt.Println("begin receive msg ====================")
	for {
		msgID, msg, _ := GetInstance().ConsumeMsgByPull("testTopic", "subsName")
		fmt.Printf("receive msg: %s, %s, %s", "subsName", msgID, msg)
		GetInstance().AckMsgId("testTopic", "subsName", msgID)
		time.Sleep(2 * time.Second)
	}
}

func recMsgT2() {
	fmt.Println("begin receive msg ====================")
	for i := 0; i < 100; i++ {
		msgID, msg, _ := GetInstance().ConsumeMsgByPull("testTopic", "subsName2")
		fmt.Printf("receive msg: %s, %s, %s", "subsName2", msgID, msg)
		GetInstance().AckMsgId("testTopic", "subsName", msgID)
		time.Sleep(3 * time.Second)
	}
}
