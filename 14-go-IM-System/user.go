package main

import "net"

type User struct {
	Name   string
	Addr   string      //当前客户端地址
	C      chan string //每个用户的channel
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() //当前客户端地址
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动监听当前user，channel消息的goroutine
	go user.ListenMessage()
	return user
}

// 用户上线业务
func (this *User) Online() {
	//用户上线。将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	//广播当前用户上线
	this.server.BroadCast(this, "已上线")
}

// 用户下线业务
func (this *User) Offline() {
	//用户下线。将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	//广播当前用户上线
	this.server.BroadCast(this, "下线")

}
func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)

}

// 监听当前User channel的方法，一旦有消息就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}

}
