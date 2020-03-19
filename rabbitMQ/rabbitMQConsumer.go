package rabbitMQ

import "github.com/streadway/amqp"

type Consumer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	msgs chan amqp.Delivery
}

func InitConsumerConn(hostName string) (*Consumer,error){
	c,err:=amqp.Dial(hostName)
	if err!=nil{
		return nil,err
	}
	var mq Consumer
	mq.conn=c
	mq.ch,err=c.Channel()
	return &mq,err
}

func(mq *Consumer) Consume(queue string,consumer string) (<-chan amqp.Delivery,error){
	msgs,err:=mq.ch.Consume(queue,consumer,false,false,false,false,nil)
	return msgs,err
}

func(mq *Consumer) Dispose() error{
	if mq.conn.IsClosed(){
		return nil
	}else {
		err:=mq.conn.Close()
		return err
	}
}

