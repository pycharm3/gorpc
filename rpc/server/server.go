package main

import (
    "net/rpc"
    "net"
	"net/http"
    "time"
    "errors"
    "gorpc/rpc/message"
)

type OrderService struct{

}

func (os *OrderService)GetOrderInfo(request message.OrderRes,response *message.OrderInfo)error{
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
    tr := new(OrderService)
    err := rpc.Register(tr)
    if err != nil{
        panic(err.Error())
    }

    rpc.HandleHTTP()

    listen,err := net.Listen("tcp",":8080")
    if err != nil{
        panic(err.Error())
    }
    http.Serve(listen,nil)
}