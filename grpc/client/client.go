package client

import (
	"net/rpc"
    "net"
	"net/http"
    "time"
    "errors"
    "gorpc/grpc/message"
)

func main(){
	conn, err := grpc.Dial("localhost:8080",grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)
	timeStamp := time.Now().Unix()
	request := &message.OrderReq{OrderId:"20191218001",TimeStamp:timeStamp}
	orderInfo,err := orderServiceClient.GetOrderInfo(context.Background(),request)
	
	if orderInfo != nil{
		fmt.Println(orderInfo)
	}
}
