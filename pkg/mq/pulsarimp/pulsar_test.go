package pulsarimp

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestPulsar(t *testing.T) {
	mq, _ := New(Configuration{
		Url:          "",
		Token:        "",
		ListenerName: "",
	}, log.Default())
	go sendMsgT(mq)
	go recMsgT2(mq)
	recMsgT(mq)
}

func sendMsgT(ins *PulsarService) {
	fmt.Println("begin produce msg ====================")
	for {
		time.Sleep(1 * time.Second)
		ins.SendMsg("testTopic", []byte("hello!!!"+time.Now().String()))
	}
}

func recMsgT(ins *PulsarService) {
	fmt.Println("begin receive msg ====================")
	for {
		msgID, msg, _ := ins.ConsumeMsgByPull("testTopic", "subsName")
		fmt.Printf("receive msg: %s, %s, %s", "subsName", msgID, msg)
		ins.AckMsgId("testTopic", "subsName", msgID)
		time.Sleep(2 * time.Second)
	}
}

func recMsgT2(ins *PulsarService) {
	fmt.Println("begin receive msg ====================")
	for i := 0; i < 100; i++ {
		msgID, msg, _ := ins.ConsumeMsgByPull("testTopic", "subsName2")
		fmt.Printf("receive msg: %s, %s, %s", "subsName2", msgID, msg)
		ins.AckMsgId("testTopic", "subsName", msgID)
		time.Sleep(3 * time.Second)
	}
}
