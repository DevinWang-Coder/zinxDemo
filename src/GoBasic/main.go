package main

import (
	"fmt"
	"go_basic/arrayDemo"
	"go_basic/ptrDemo"
	"go_basic/sliceDemo"
	"go_basic/stringsDemo"
	"math"
	"math/cmplx"
)

var (
	name = "jason"
	age = 18
	gender string
)

const (
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func factorization()  {
	fmt.Println(name)
	fmt.Println(age)
	gender = "男"
	fmt.Println(gender)
	// 在局部使用
	var (
		x1 = 123
		x2 = 456
	)
	fmt.Println(x1)
	fmt.Println(x2)
}
/*全局变量:
	如果全局变量名首字母小写，则只能被当前包中的go文件使用
	如果全局变量名首字母大写，则任意文件都能使用
  局部变量：
	go变量有作用域之分，每一个大括号就是一个作用域，每个作用域都可以定义相关的局部变量，不能跨域操作

*/
func iotaModel()  {
	//const 为go语言定义常量的标识关键字
	// iota
	//const (
	//	v1 = 1
	//	v2 = 2
	//	v3 = 3
	//	v4 = 4
	//	v5 = 5
	//)
	//fmt.Println(v1, v2, v3, v4, v5)

	//2
	const (
		v1 = iota + 2
		v2
		v3
		v4
		v5
	)
	fmt.Println(v1, v2, v3, v4, v5)

}

//输入
func inputModel()  {

	/*
		var name string
		fmt.Println("please input name:")
		fmt.Scan(&name)
		fmt.Println(name)
	*/

	var name string
	var age int

	fmt.Println("please input name:")
	_, err := fmt.Scan(&name, &age)
	if err == nil {
		fmt.Println(name, age)
	}else {
		fmt.Println("input data err", err)
	}
}

//func readerModel()  {
//	reader := bufio.NewReader(os.Stdin)
//	// line，从stdin中读取一行的数据（字节集合 -> 转化成为字符串）
//	// reader默认一次能4096个字节（4096/3)
//	//    1. 一次性读完，isPrefix=false
//	//       2. 先读一部分，isPrefix=true，再去读取isPrefix=false
//	line, _, _ = reader.ReadLine()
//	data := string(line)
//	fmt.Println(data)
//}

func ecler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(
		cmplx.Pow(math.E, 1i* math.Pi) + 1)
}

func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main()  {
	//this is my first go
	//fmt.Println("start")
	//var name string = "jason"
	//fmt.Println(name)

	//if true {
	//	var age = 18
	//	fmt.Println(age)
	//}
	//fmt.Println(age)
	//factorization()
	//fmt.Println("##########")
	//iotaModel()
	//inputModel()
	//readerModel()
	ecler()
	triangle()

	//字符串
	stringsDemo.StringsDemo()
	//指针
	ptrDemo.PtrDemo()
	//数组
	arrayDemo.ArrayDemo()
	//切片
	sliceDemo.SliceDemo()


}


