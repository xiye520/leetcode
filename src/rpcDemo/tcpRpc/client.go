package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

/*
E:\xiyegit\leetcode\src\rpcDemo>go run tcpRpc\client.go  127.0.0.1:1234 20 1
客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go
长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip= 127.0.0.1:1234 嘿嘿,加上后面的被除数os.Args[2]= 20 和除数os.Args[3]= 1
我进行正常结果如下
Num : 20*1=20
Num: 20/1=0 余数 0

*/
func main() {
	fmt.Println("客户端 其他端 去调用的地方  对应的例子是 GoTCPRPC9.go")

	if len(os.Args) == 4 {
		fmt.Println("长度必须等于4,因为呢，你输入的肯定是一个ip的地址ip=", os.Args[1], "嘿嘿,加上后面的被除数os.Args[2]=", os.Args[2], "和除数os.Args[3]=", os.Args[3])
		//os.Exit(1)
	}
	// 获取 ip 地址
	service := os.Args[1]
	//连接 拨号连接到指定的网络地址的RPC服务器。
	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("老子在连接Dial的发生了错误，我要退出了", err)
	}
	num1 := os.Args[2]
	i1, error1 := strconv.Atoi(num1)
	if error1 != nil {
		fmt.Println("自己不知道 自己输入错误了啊 请看error ：", error1)
		os.Exit(1)
	}
	num2 := os.Args[3]
	i2, error2 := strconv.Atoi(num2)
	if error2 != nil {
		fmt.Println("自己不知道 自己输入错误了啊 请看error ：", error2)
		os.Exit(1)
	}
	aa := AAA{i1, i2}
	var reply int
	err1 := client.Call("Num.M", aa, &reply)

	if err1 != nil {
		log.Fatal("我要退出了，因为我在Call的时候发生了 错误", err1)
	}
	fmt.Println("我进行正常结果如下")
	fmt.Printf("Num : %d*%d=%d\n", aa.A, aa.B, reply)

	var bb BDemo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Num.F", aa, &bb)
	if err != nil {
		log.Fatal("我对这个方法发生了过敏的反应 哈哈哈哈  err=====", err)
	}
	fmt.Printf("Num: %d/%d=%d 余数 %d\n", aa.A, aa.B, bb.DD, bb.CC)

}

// 定义两个类，那边需要操作的类
type AAA struct {
	A, B int
}

//记住这里不能够大写 两个连着一起大写 有点意思
//reading body gob: type mismatch: no fields matched compiling decoder for  DDDD
//  todo 为啥 第二个参数  只要是两个连在一起的DDDD   就会报错   reading body gob: type mismatch: no fields matched compiling decoder for
type BDemo struct {
	DD, CC int
}
