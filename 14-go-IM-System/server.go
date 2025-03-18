package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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
	user := NewUser(conn, this)
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接受客户端发送的数据
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn read err:", err)
				return
			}
			//提取用户的消息(去除‘\n’)
			msg := string(buf[:n-1])
			//将得到的信息进行广播
			user.DoMessage(msg)

			//用户的任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	// 4. 监听用户是否超时
	timer := time.NewTimer(time.Second * 10) // 创建定时器
	defer timer.Stop()                       // 退出时释放资源

	for {
		select {
		case <-isLive:
			// 用户活跃，重置定时器
			if !timer.Stop() {
				<-timer.C // 清空通道，防止旧数据影响
			}
			timer.Reset(time.Second * 10)

		case <-timer.C:
			// 超时处理
			user.SendMsg("你被踢了")

			// 关闭用户 channel（避免重复关闭导致 panic）
			select {
			case <-user.C:
			default:
				close(user.C)
			}

			// 关闭连接
			conn.Close()
			return
		}
	}
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
