package main

import (
	"runtime"
	thelper "loadgen/testhelper"
	"fmt"
	"time"
)

func main() {
	// 设置P最大数量
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化服务器
	server := thelper.NewTcpServer()
	//defer server.Close()
	serverAddr := "10.10.80.113:8080"
	fmt.Printf("Startup TCP server(%s)...\n", serverAddr)
	err := server.Listen(serverAddr)
	if err != nil {
		println("TCP Server startup failing! (addr=%s)!\n", serverAddr)
	}
	time.Sleep(time.Second * 100)
}
