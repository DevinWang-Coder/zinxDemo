/*
   @Author: Jason Chen
   @Date: 2020/11/12 14:00
   @Code: UTF-8
*/
package main

import "zinx/znet"

func main()  {
	//zinx
	//创建一个句柄，使用zinx的api
	s := znet.NewServer("[zinx 1.0]")
	//启动server
	s.Serve()
}
