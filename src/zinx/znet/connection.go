/*
   @Author: Jason Chen
   @Date: 2020/11/12 14:46
   @Code: UTF-8
*/
package znet

import (
	"net"
	"zinx/ziface"
)

//链接模块
type Connection struct {
	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接的ID
	ConnID uint32

	//当前链接的状态
	isClosed bool

	//当前链接所绑定的处理业务方法的Api
	handleAPI ziface.HandleFunc

	//告知当前链接已经退出或者停止 channel
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection  {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: callback_api,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}