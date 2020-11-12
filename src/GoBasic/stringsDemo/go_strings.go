package stringsDemo

import (
	"fmt"
	"strconv"
	"strings"
)

//遍历字符串
func bianliString()  {
	myname := "hello go"
	for _, ret := range myname{
		fmt.Printf("ret=%c\n", ret)
	}
}

func xiugaiString()  {
	myname := "hello jason"
	m1 := []rune(myname) //转化为[]int32的切片，rune是int32的别名
	fmt.Println(m1)
	//m1[4] = "" //修改索引对应的值
	myname = string(m1) //类型强制转换，rune强制转换为string
	fmt.Println(myname)
}

func stringsPkg()  {
	str := "hello world"

	//判断是不是以某个字符串开头
	res0 := strings.HasPrefix(str, "http://")
	res1 := strings.HasPrefix(str, "hello")
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("ret1 is %v\n", res1)

	//判断是不是以某个字符串结尾
	res3 := strings.HasSuffix(str, "http")
	res4 := strings.HasSuffix(str, "world")
	fmt.Printf("res3 is %v\n", res3)
	fmt.Printf("res4 is %v\n", res4)

	//判断字符在字符串中首次出现的索引位置， 没有则返回1
	res5 := strings.Index(str, "o")
	res6 := strings.Index(str, "x")
	fmt.Printf("res5 is %v\n", res5)
	fmt.Printf("res6 is %v\n", res6)

	//求字符在字符串中出现的次数，不存在返回0次
	countTime0 := strings.Count(str, "h")
	countTime1 := strings.Count(str, "x")
	fmt.Printf("countTime0 is %v\n", countTime0)
	fmt.Printf("countTime1 is %v\n", countTime1)

	//重复几次字符串
	res11 := strings.Repeat(str, 0)
	res12 := strings.Repeat(str, 1)
	res13 := strings.Repeat(str, 2)
	// strings.Repeat("原字符串", 重复次数)
	fmt.Printf("res11 is %v\n", res11)
	fmt.Printf("res12 is %v\n", res12)
	fmt.Printf("res13 is %v\n", res13)

	//字符串改大写
	res14 := strings.ToUpper(str)
	fmt.Printf("res14 is %v\n", res14)

	//字符串改小写
	res15 := strings.ToLower(str)
	fmt.Printf("res15 is %v\n", res15)

	//去除首尾的空格
	res16 := strings.TrimSpace(str)
	fmt.Printf("res16 is %v\n", res16)

	//去除首尾指定的字符,遍历l、d、e然后去除
	res17 := strings.Trim(str, "ld")
	fmt.Printf("res17 is %v\n", res17)

	//去除开头指定的字符
	res18 := strings.TrimLeft(str, "he")
	fmt.Printf("res18 is %v\n", res18)

	//去除结尾指定的字符,遍历d、l、r
	res19 := strings.TrimRight(str, "dlr")
	fmt.Printf("res19 is %v\n", res19)

	//用指定的字符串将string类型的切片元素结合
	str1 := []string{"hello", "world", "hello", "golang"}
	res20 := strings.Join(str1, "+")
	fmt.Printf("res20 is %v\n", res20)

}

//基本数据类型转化
func baseDataTypeChange()  {
	/*
	go在不同的数据类型之间赋值时需要显示转化，无法自动转换。
	注意：
		1.go中数据类型的转换可以从  数值范围较大的 -> 数值范围较小，
		2.被转换的变量存储的数值，变量本身的数据类型没有变化
	*/
	var num1 int32 = 100
	var num2 float32 = float32(num1) //num2强制转换成浮点型
	fmt.Printf("num1=%v num2=%v\n", num1, num2) //%v 值的默认格式
	fmt.Printf("num1类型：%T, num2的类型：%T\n", num1, num2) //本身的类型没有变化

	//不同变量之间的计算
	var n1 int32 = 12
	var n2 int64 = 11

	//n3 := n1 + n2 //不同类型之间无法计算，需要强制转换
	n3 := int64(n1) + n2
	fmt.Println(n3)
}

func baseDataToString1()  {
	var num1 int = 66
	var num2 float64 = 25.25
	var b bool = true
	var myChar byte = 'c'

	//%q 单引号
	//%d 十进制表示
	str1 := fmt.Sprintf("%d", num1)
	fmt.Printf("str1 type %T str1=%q\n", str1, str1)

	//%f 有小数点
	str2 := fmt.Sprintf("%f", num2)
	fmt.Printf("str2 type %T str2=%q\n", str2, str2)

	//%t 布尔值
	str3 := fmt.Sprintf("%t", b)
	fmt.Printf("str3 type %T str3=%q\n", str3, str3)

	//%c Unicode码对应的字符
	str4 := fmt.Sprintf("%c", myChar)
	fmt.Printf("str4 type %T str4=%q\n", str4, str4)
}

func baseDataToString2()  {
	var num1 int = 99
	var num2 float64 = 66.66
	var b1 bool = true
	str1 := strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str1类型是%T str1=%q\n", str1, str1)
	//参数解释
	// f 格式
	// 10 小数位保留10位
	// 64  表示float64
	str2 := strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str2类型是%T str2=%q\n", str2, str2)
	str3 := strconv.FormatBool(b1)
	fmt.Printf("str3类型是%T str3=%q\n", str3, str3)
	//Itoa，将int转为string
	var num3 int64 = 1123
	str4 := strconv.Itoa(int(num3))//必须强转int()
	fmt.Printf("str4类型是%T str4=%q\n", str4, str4)
}

func StringsDemo() {
	fmt.Println("stringDemo")

	// 拼接字符串
	// go的字符串是不可变的
	var city string = "I love beijing"
	fmt.Println(city)

	// 字符串用双引号识别，识别转义字符“\n  \t”
	var story string = "您好\n老板"
	fmt.Println(story)

	//反引号，以原生的形式输出，包含特殊字符。防止注入攻击
	story2 := `
	反引号，
以原生的形式\n输出
`
	fmt.Println(story2)

	//字符串拼接
	//注意：换行拼接的时候需要+加号写在上一行
	str1 := "hello" + "boss"
	fmt.Println(str1)
	myname := "wo" + "shi" + "bei" + "jing" +
		"sha" + "he" + "yu"
	fmt.Println(myname)

	//使用索引取出某字节
	str3 := " hello goland"
	//索引取出的码值，格式化输出
	fmt.Printf("str3=%c\n", str3[2])
	//输出str3的长度
	fmt.Println(len(str3))

	//遍历字符串
	bianliString()

	//修改字符串
	xiugaiString()

	//strings包
	stringsPkg()

	//基础数据类型转换
	baseDataTypeChange()

	//其他数据类型转换成字符串
	baseDataToString1()
	baseDataToString2()
}