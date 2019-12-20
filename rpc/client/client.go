package main

import (
    "net/rpc"
	"fmt"
    "time"
    "gorpc/rpc/message"
)

func main(){
    // 拨号连接server服务
    client,err := rpc.DialHTTP("tcp","localhost:8080")
    if err != nil{
        panic(err.Error())
    }
    // 定义请求参数req和接收响应参数res
    timeStamp := time.Now().Unix()
    request := message.OrderRes{OrderId : "20191218001",TimeStamp :timeStamp}
    var response *message.OrderInfo
    err = client.Call("OrderService.GetOrderInfo",request,&response)
    if err != nil{
        panic(err.Error())
    }
    fmt.Println(*response)
}