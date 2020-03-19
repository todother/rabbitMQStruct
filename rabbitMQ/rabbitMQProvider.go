package rabbitMQ

import "github.com/streadway/amqp"

type MessageType string

const (
	TextMessage = "text/plain"
	JsonMessage = "application/json"
)

type Provider struct {
	conn       *amqp.Connection
	ch         *amqp.Channel
	queue      string
	exchange   string
	exType     string
	routingKey string
	host       string
	connType   string
}

func (mq *Provider) Publish(messageType string, message string) error {
	err := mq.ch.Publish(mq.exchange, mq.routingKey, false, false, amqp.Publishing{
		ContentType: messageType,
		Body:        []byte(message),
	})
	return err
}

func (mq *Provider) PublishWithParam(messageType string, message string,exchangeName string,routingKey string) error {
	err := mq.ch.Publish(exchangeName, routingKey, false, false, amqp.Publishing{
		ContentType: messageType,
		Body:        []byte(message),
	})
	return err
}

func InitProviderConn(hostName string) (*Provider, error) {
	c, err := amqp.Dial(hostName)
	if err != nil {
		return nil, err
	}
	var mq Provider
	mq.conn = c
	mq.host = hostName
	mq.ch, err = c.Channel()
	return &mq, err
}

func (mq *Provider) DeclareExchange(exchangeName string, exchangeType string) error {
	err := mq.ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	mq.exchange = exchangeName
	return err
}

func (mq *Provider) DeclareQueue(queueName string) error {
	_, err := mq.ch.QueueDeclare(queueName, true, false, false, false, nil)
	mq.queue = queueName
	return err
}

func (mq *Provider) BindQueue(queueName string, exchangeName string, routingKey string) error {
	err := mq.ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
	mq.routingKey = routingKey
	return err
}

func (mq *Provider) Dispose() error {
	if mq.conn.IsClosed() {
		return nil
	} else {
		err := mq.conn.Close()
		return err
	}
}

