/*
   @Author: Jason Chen
   @Date: 2020/11/12 10:37
   @Code: UTF-8
*/
package ziface



//定义一个服务器接口

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()


}