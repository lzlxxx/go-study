package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User //在线用户列表
	mapLock   sync.RWMutex
	Message   chan string //消息广播的channel
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}

}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg

}
func (this *Server) Handler(conn net.Conn) {
	//...当前连接的业务
	//fmt.Println("连接建立成功")
	user := NewUser(conn)
	//用户上线。将用户加入到onlineMap中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	//广播当前用户上线
	this.BroadCast(user, "已上线")
	//当前handler阻塞
	select {}
}

// 启动服务器的接口
func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err：", err)
		return
	}
	defer listen.Close()

	go this.ListenMessage()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
			continue
		}

		go this.Handler(conn)
	}

}
