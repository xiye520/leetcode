package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

func init() {
	fmt.Println("基于TCP协议实现的RPC，服务端的代码如下")
}

type Me struct {
	A, B int
}
type You struct {
	CC, D int
}
type Num int

/*
Go RPC的函数只有符合下面的条件才能够被远程访问，不然会被忽略
1 函数必须是导出的（首字母大写）
2 必须有两个导出类型的参数
3 第一个参数是接受的参数，第二个参数是返回给客户端的参数，第二个参数必须是指正类型的
4 函数还必须有一个返回值error
*/
func (n *Num) M(args *Me, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (n *Num) F(args *Me, u *You) error {
	if args.B == 0 {
		return errors.New("输入不能够为0 被除数")
	}
	u.D = args.A / args.B
	u.CC = args.A % args.B
	return nil
}

func main() {
	//内建函数new本质上说跟其它语言中的同名函数功能一样：new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：
	//new返回指针。
	num := new(Num)
	rpc.Register(num)
	//ResolveTCPAddr返回TCP端点的地址。
	//网络必须是TCP网络名称。
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")

	if err != nil {
		fmt.Println("错误了哦")
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		// todo   需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc 来处理
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
