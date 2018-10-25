package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//使用Go提供的json-rpc 标准包
func init() {
	fmt.Println("JSON RPC 采用了JSON，而不是 gob编码，和RPC概念一模一样，")
}

type Work struct {
	Who, DoWhat string
}

type DemoM string

func (m *DemoM) DoWork(w *Work, whoT *string) error {
	*whoT = "是谁：" + w.Who + "，在做什么---" + w.DoWhat
	return nil
}

/*
E:\xiyegit\leetcode\src\rpcDemo>go run jsonRpc\client.go 127.0.0.1:8080  shiming gongzuo
这是客户端，用来启动，通过命令行来启动
客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go
长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip= 127.0.0.1:8080 嘿嘿,加上后面的被除数os.Args[2]= shiming 和除数os.Args[3]= gongzuo
收到信息了 是谁：shiming，在做什么---gongzuo
*/
func main() {
	str := new(DemoM)
	rpc.Register(str)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("大哥发生错误了啊，请看错误 ResolveTCPAddr err=", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("发生错误了--》err=", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)

	}

}
