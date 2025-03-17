package main

import "net"

type User struct {
	Name string
	Addr string      //当前客户端地址
	C    chan string //每个用户的channel
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String() //当前客户端地址
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	//启动监听当前user，channel消息的goroutine
	go user.ListenMessage()
	return user
}

// 监听当前User channel的方法，一旦有消息就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}

}
