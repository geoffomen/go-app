package kafkaimp

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/geoffomen/go-app/internal/pkg/mylog"
)

type KafkaService struct {
	producer producer
	consumer Consumer
	conf     kafkaConf
}

type kafkaConf struct {
	brokers          []string
	producerRetryMax int
	version          string
	groupName        string
	assignor         string
	oldest           bool
	saslUserName     string
	saslUserPass     string
}

func New(cf kafkaConf) (*KafkaService, error) {
	ins := KafkaService{}
	ins.conf = cf
	err := ins.initSyncProducer()
	if err != nil {
		return nil, fmt.Errorf("初始化消息队列失败, err:%s", err)
	}

	err = ins.initPushConsumer()
	if err != nil {
		return nil, fmt.Errorf("初始化消息队列失败, err:%s", err)
	}

	return &ins, nil
}

// producer ..
type producer struct {
	syncProducer sarama.SyncProducer
}

func (p *producer) Close() {
	if err := p.syncProducer.Close(); err != nil {
		mylog.Infof(err.Error())
	}
}

// initSyncProducer 初始化一个同步发送消息的生产者
func (srv *KafkaService) initSyncProducer() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = srv.conf.producerRetryMax
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Net.SASL.Enable = true
	config.Net.SASL.Handshake = true
	config.Net.SASL.User = srv.conf.saslUserName
	config.Net.SASL.Password = srv.conf.saslUserPass

	sp, err := sarama.NewSyncProducer(srv.conf.brokers, config)
	if err != nil {
		return err
	}
	srv.producer = producer{
		syncProducer: sp,
	}

	return nil
}

// initPushConsumer 初始化一个接收消息推送的消费者
func (srv *KafkaService) initPushConsumer() error {
	version, err := sarama.ParseKafkaVersion(srv.conf.version)
	if err != nil {
		return err
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Net.SASL.Enable = true
	config.Net.SASL.Handshake = true
	config.Net.SASL.User = srv.conf.saslUserName
	config.Net.SASL.Password = srv.conf.saslUserPass

	switch srv.conf.assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	}

	if srv.conf.oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	} else {
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
	}

	consumer := Consumer{
		groupName:      srv.conf.groupName,
		handlers:       make(map[string]func(content []byte) error),
		consumerGroups: make(map[string]sarama.ConsumerGroup),
		config:         *config,
	}
	srv.consumer = consumer

	return nil
}

// Consumer ..
type Consumer struct {
	handlers       map[string]func(content []byte) error
	groupName      string
	consumerGroups map[string]sarama.ConsumerGroup
	config         sarama.Config
}

func (consumer *Consumer) Close() {
	for _, cg := range consumer.consumerGroups {
		cg.Close()
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h, ok := consumer.handlers[message.Topic]
		if !ok {
			mylog.Errorf("没有找到消息处理器，topic: %s", message.Topic)
		}
		mylog.Debugf("收到消息， topic: %s, offset: %d, content: %s", message.Topic, message.Offset, message.Value)
		err := h(message.Value)
		if err != nil {
			mylog.Warnf("消息处理失败")
			mylog.Debugf("消息处理失败, %s", string(message.Value))
		} else {
			session.MarkMessage(message, "")
		}
		session.MarkMessage(message, "")
	}

	return nil
}

func (srv *KafkaService) SendMsg(topic string, msg []byte) (string, error) {
	payload := &sarama.ProducerMessage{Topic: topic, Value: sarama.ByteEncoder(msg)}
	partition, offset, err := srv.producer.syncProducer.SendMessage(payload)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d_%d", partition, offset), nil
}

var lock = &sync.Mutex{}

func (srv *KafkaService) ConsumeMsgByPush(topic string, handler func(content []byte) error) error {
	if _, ok := srv.consumer.consumerGroups[topic]; ok {
		return nil
	}
	lock.Lock()
	defer lock.Unlock()
	if _, ok := srv.consumer.consumerGroups[topic]; ok {
		return nil
	}

	srv.consumer.handlers[topic] = handler
	cg, err := sarama.NewConsumerGroup(srv.conf.brokers, srv.consumer.groupName, &srv.consumer.config)
	if err != nil {
		return err
	}
	srv.consumer.consumerGroups[topic] = cg

	go func() {
		ctx := context.Background()
		for {
			err := cg.Consume(ctx, []string{topic}, &srv.consumer)
			if err != nil {
				// fmt.Errorf("消费者加入topic时发生错误: %s", err)
				time.Sleep(5 * time.Second)
			}
		}
	}()

	return nil
}

func (srv *KafkaService) ConsumeMsgByPull(topic, subsName string) (string, string, error) {
	return "", "", fmt.Errorf("not implemented")
}

func (srv *KafkaService) AckMsgId(topic, subsName, msgId string) error {
	return fmt.Errorf("not implemented")
}
