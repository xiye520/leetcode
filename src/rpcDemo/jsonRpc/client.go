package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	fmt.Println("这是客户端，用来启动，通过命令行来启动")

	fmt.Println("客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go")

	if len(os.Args) == 4 {
		fmt.Println("长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip=", os.Args[1], "嘿嘿,加上后面的被除数os.Args[2]=", os.Args[2], "和除数os.Args[3]=", os.Args[3])
		//os.Exit(1)
	}

	service := os.Args[1]
	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("Dial 发生了错误了哦 错误的信息为   err=", err)
	}
	send := Send{os.Args[2], os.Args[3]}
	var resive string
	err1 := client.Call("DemoM.DoWork", send, &resive)
	if err1 != nil {
		fmt.Println("shiming call error    ")
		fmt.Println("Call 的时候发生了错误了哦  err=", err1)
	}
	fmt.Println("收到信息了", resive)

}

// 类可以不一样 但是 Who 和DoWhat 要必须一样  要不然接收到不到值，等我在详细的了解了 才去分析下原因  感觉有点蒙蔽啊
type Send struct {
	Who, DoWhat string
}
