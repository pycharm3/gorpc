package main

import (
    "net/rpc"
    "net"
	"net/http"
    "time"
    "errors"
    "gorpc/grpc/message"
)

type OrderServiceImpl struct{

}

func (os *OrderServiceImpl)GetOrderInfo(request message.OrderReq,response *message.OrderInfo)error{
    orderMap := map[string]message.OrderInfo{
        "20191218001" : {OrderId: "20191218001",OrderName : "mysuit",OrderStatus : "已付款"},
        "20191218002" : {OrderId:"20191218002",OrderName : "mycar",OrderStatus : "已付款"},
        "20191218003" : {OrderId:"20191218003",OrderName : "myhouse",OrderStatus : "未付款"},
    }

    current := time.Now().Unix()
    if (request.TimeStamp > current){
        *response = message.OrderInfo{OrderId: "0",OrderName : "",OrderStatus : "订单异常"}
    }else{
        result := orderMap[request.OrderId]
        if result.OrderId != ""{
            *response = orderMap[request.OrderId]
        }else{
            return errors.New("server error")
        }
    }
    return nil
}

// main 方法
func main(){
    // 创建 gRPC 服务器的一个实例
    server := grpc.NewServer()
    // 将RegisterOrderServiceServer实例化出一个OrderServiceImpl
    message.RegisterOrderServiceServer(server,new(OrderServiceImpl))
    // 返回一个监听器
    listen,err := net.Listen("tcp",":8080")
    if err != nil{
        panic(err.Error())
    }
    // 监听器监听到连接请求交给server处理
    server.Serve(listen)
}