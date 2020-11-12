/*
   @Author: Jason Chen
   @Date: 2020/11/12 14:46
   @Code: UTF-8
*/
package znet

import (
	"fmt"
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

func (c *Connection) StartReader()  {
	fmt.Println("Reader Goroutine is running..")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit ,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for true {
		//读客户端的数据
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		//调用当前链接所绑定的HandleApi
		if err := c.handleAPI(c.Conn, buf,cnt);err != nil {
			fmt.Println("ConnID", c.ConnID, "handle err", err)
		}
	}

}

//启动连接 让当前的链接准备工作
func (c *Connection) Start() {
	fmt.Println("Conn start ... ConnID = ", c.ConnID)
	//启动从当前链接的读数据业务

	go c.StartReader()
	//TODO 启动当前链接写数据的业务

}

//停止链接 结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn stop() .. connID = ", c.ConnID)

	//如果当前链接以及关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//关闭socket链接
	c.Conn.Close()

	//回收资源
	close(c.ExitChan)


}

//获取当前链接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}


// 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}


// 获取远程客户端的TCP状态 IP Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}


//发送数据， 将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}