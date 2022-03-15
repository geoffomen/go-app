package pulsarimp

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarService struct {
	client pulsar.Client
	logger LogIface
}

type Configuration struct {
	Url          string
	Token        string
	ListenerName string
}

type LogIface interface {
	Printf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

func New(cf Configuration, logger LogIface) (*PulsarService, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://" + cf.Url,
		ListenerName:      cf.ListenerName,
		Authentication:    pulsar.NewAuthenticationToken(cf.Token),
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("could not instantiate Pulsar client: %v", err)
	}

	return &PulsarService{client: client, logger: logger}, nil
}

func (p *PulsarService) getProducer(topic string) (pulsar.Producer, error) {
	producer, err := p.client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func (p *PulsarService) getConsumer(topic, subNm string) (pulsar.Consumer, error) {
	consumer, err := p.client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: subNm,
		Type:             pulsar.Shared,
	})
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (p *PulsarService) SendMsg(topic string, msg []byte) (string, error) {
	producer, err := p.getProducer(topic)
	if err != nil {
		return "", err
	}
	defer producer.Close()

	msgId, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: msg,
	})
	if err != nil {
		return "", err
	}
	return string(msgId.Serialize()), nil
}

func (p *PulsarService) ConsumeMsgByPull(topic, subsName string) (string, string, error) {
	consumer, err := p.getConsumer(topic, subsName)
	if err != nil {
		return "", "", err
	}
	defer consumer.Close()

	msg, err := consumer.Receive(context.Background())
	if err != nil {
		return "", "", err
	}
	return string(msg.ID().Serialize()), string(msg.Payload()), nil
}

func (p *PulsarService) ConsumeMsgByPush(topic string, handler func(content []byte) error) error {
	return fmt.Errorf("not implemented")
}

func (p *PulsarService) AckMsgId(topic, subsName, msgId string) error {
	consumer, err := p.getConsumer(topic, subsName)
	if err != nil {
		return err
	}
	defer consumer.Close()

	pmid, err := pulsar.DeserializeMessageID([]byte(msgId))
	if err != nil {
		return err
	}
	consumer.AckID(pmid)
	return nil
}
