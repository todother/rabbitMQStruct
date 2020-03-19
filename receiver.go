package main

import (
	"com.redis.guojio/redistest/rabbitMQ"
	"fmt"
)

func main() {
	mq,err:=rabbitMQ.InitConsumerConn("amqp://guest:guest@127.0.0.1:5672/")
	defer mq.Dispose()
	if err!=nil{
		fmt.Println(err)
	}
	forever:=make(chan bool)
	consumer:="consumer3"
	msgs,err:=mq.Consume("QueueStruct",consumer)

	go func() {
		for item:=range msgs{
			fmt.Println("recerve msg in ",consumer," value is ",string( item.Body))
			item.Ack(false)
		}
	}()
	<-forever
	//conn,err:=amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	//defer conn.Close()
	//if err!=nil{
	//	fmt.Println("fail to connect to rabbitmq ,",err)
	//	panic(conn)
	//}
	//ch,err:=conn.Channel()
	//defer ch.Close()
	//if err!=nil{
	//	fmt.Println("fail to generate channel ,",err)
	//	panic(ch)
	//}
	//err=ch.ExchangeDeclare("ex1",amqp.ExchangeTopic,true,false,false,false,nil)
	//if err!=nil{
	//	fmt.Println("fail to declare exchange ,",err)
	//	//panic(q)
	//}
	//err=ch.QueueBind("q1","good.*.*","ex1",false,nil)
	//err=ch.QueueBind("q2","#.a","ex1",false,nil)
	//msgs1,err:=ch.Consume("q1","consumer1",false,false,false,false,nil)
	//msgs2,err:=ch.Consume("q2","consumer2",false,false,false,false,nil)
	//
	//if err!=nil{
	//	fmt.Println("failed to consumer record :" ,err)
	//}
	//forever:=make(chan bool)
	//go func() {
	//	for item:=range msgs1{
	//		fmt.Println("receive a message from q1, ",string(item.Body))
	//		item.Ack(false)
	//	}
	//}()
	//
	//go func() {
	//	for item:=range msgs2{
	//		fmt.Println("receive a message from q2, ",string(item.Body))
	//		item.Ack(false)
	//	}
	//}()
	//<-forever
}

