//package main
//
//import (
//	"fmt"
//	"github.com/gomodule/redigo/redis"
//	"time"
//)
//
//func main() {
//
//	c,err:=redis.Dial("tcp","192.168.0.4:6379")
//	if err!=nil{
//		fmt.Println(err)
//	}
//	defer c.Close()
//	now:=time.Now()
//	i:=0
//	for i=0;i<10000;i++{
//		_,err=c.Do("set","mykey","todother","ex","100")
//		if err!=nil {
//			fmt.Println(err)
//		}
//	}
//	//value,err:=redis.String( c.Do("get","mykey"))
//	//if err!=nil{
//	//	fmt.Println(err)
//	//}
//	diff:=time.Now().Sub(now)
//	fmt.Println("this func runs ",diff)
//
//	now=time.Now()
//	for i=0;i<10000;i++{
//		_,err=c.Do("get","mykey")
//		if err!=nil {
//			fmt.Println(err)
//		}
//	}
//	diff=time.Now().Sub(now)
//	fmt.Println("this func runs ",diff)
//	//fmt.Println("valur is ",value)
//}

package main

import (
	"com.redis.guojio/redistest/rabbitMQ"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {

	forever :=make(chan bool)
	mq,err:=rabbitMQ.InitProviderConn("amqp://guest:guest@127.0.0.1:5672/")
	defer mq.Dispose()
	if err!=nil{
		fmt.Println(err)
	}

	if err!=nil{
		fmt.Println(err)
	}
	err=mq.DeclareQueue("QueueStruct")
	err=mq.DeclareExchange("ExStruct",amqp.ExchangeTopic)
	err=mq.BindQueue("QueueStruct","ExStruct","*.*.todother")

	for i:=0;i<10;i++{
		err=mq.PublishWithParam(rabbitMQ.TextMessage,"hello","ExStruct","c.com.todother")
	}
	fmt.Println("finished")
	<-forever
	//conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	//defer conn.Close()
	//forever:=make(chan bool)
	//if err != nil {
	//	fmt.Println("fail to connect to rabbitmq ,", err)
	//	panic(conn)
	//}
	//ch, err := conn.Channel()
	//defer ch.Close()
	//if err != nil {
	//	fmt.Println("fail to generate channel ,", err)
	//	panic(ch)
	//}
	//q1, err := ch.QueueDeclare("q1", true, false, false, false, nil)
	//if err != nil {
	//	fmt.Println("fail to declare queue ,", err)
	//	panic(q1)
	//}
	//
	//q2, err := ch.QueueDeclare("q2", true, false, false, false, nil)
	//if err != nil {
	//	fmt.Println("fail to declare queue ,", err)
	//	panic(q2)
	//}
	//err=ch.ExchangeDeclare("ex1",amqp.ExchangeTopic,true,false,false,false,nil)
	//if err!=nil{
	//	fmt.Println("fail to declare exchange ,",err)
	//	//panic(q)
	//}
	//err = ch.Publish("ex1", "good.evening.a", false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte("hello world2")})

	//fmt.Println("finished")
	//<-forever
}
