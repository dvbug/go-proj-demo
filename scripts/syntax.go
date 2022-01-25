package scripts

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

// 包初始化使用 init(){}
// 如果要在函数内部修改结构体成员的话，用指针传入是必须的；在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。

// 全局变量声明 var
var (
	goalA int
	goalB bool
)

// 全局常量声明 const
const (
	Unknown = iota
	Female
	Male
)

const (
	Tag     = "abc"
	TagLen  = len(Tag)
	TagSize = unsafe.Sizeof(Tag)
)

// 打印
func prt() {
	fmt.Println("Hello world")
}

// 打印全局, fmt格式化
func prtG() {
	fmt.Println(fmt.Sprintf("goal vars: goalA=%d, goalB=%v", goalA, goalB))
	fmt.Println(fmt.Sprintf("goal consts: Unknown=%d, Female=%d, Male=%d", Unknown, Female, Male))
	fmt.Println(fmt.Sprintf("goal consts: Tag=%s, TagLen=%d, TagSize=%d", Tag, TagLen, TagSize))
}

//变量声明
func funcvar() {
	var m int = 2
	var n int = 2
	fmt.Println("res:", m+n)
}

//格式化
func funcfmt() {
	var succ = true
	var statuscode int = 200
	var lon = 119.232322
	var lat = 39.232322

	// prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	// bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	strLon := strconv.FormatFloat(lon, 'f', -1, 32)
	strLat := strconv.FormatFloat(lat, 'f', -1, 32)
	var date = "2021-10-01"
	var url = "succ=%t&code=%d&date=%s&lon=%s&lat=%s"
	var targetUrl = fmt.Sprintf(url, succ, statuscode, date, strLon, strLat)
	fmt.Println(targetUrl)
}

// 引用
func funcref() {
	var n = 2
	// &var 获取地址, 引用类型
	// *m 获取地址的表示的值
	m := &n
	var m2 *int = &n
	fmt.Println(n, m, m2)
	n = 4
	fmt.Println(n, m, m2, *m)
}

// 指针
func funcptr() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	fmt.Println("ip == nil:", ip == nil)
	var ptr *int
	fmt.Println("ptr == nil:", ptr == nil)
}

//多返回值
func funnumbers() (int, float64, string) {
	a, b, c := 20, 64.2231, "ok"
	return a, b, c
}

// 多返回值调用
func funccallf() {
	_, f, s := funnumbers() // _ 忽略第一个返回值. _ 是一个只写变量，不能得到它的值
	fmt.Println(f, s)
}

// iota使用
func funcconstiota() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	const (
		x = "hx" // hx
		y        // hx
	)
	const (
		i2 = 1 << iota // 1<<0
		j2 = 3 << iota // 3<<1
		k2             // 3<<2
		l2             // 3<<3
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(x, y)
	fmt.Println(i2, j2, k2, l2)
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func funcselect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

// 条件语句 if switch select使用
func funccondition(i int) {
	if i < 0 {
		fmt.Println("if end", fmt.Sprintf("%d < 0", i))
	} else {
		fmt.Println("else end")
	}

	switch i {
	case 0:
		fmt.Println("case 0")
	case 1:
		fmt.Println("case 1")
	default:
		fmt.Println("case default")
	}
}

//多函数参数
func funcparams(name, nick string, age, work int) {
	fmt.Println(name, nick, age, work)

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
}

//闭包函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//闭包函数使用
func funcclosepk() {
	nextNumber := getSequence()
	fmt.Println(nextNumber(), nextNumber(), nextNumber())
}

//递归函数使用
func fibonacci3(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci3(n-2) + fibonacci3(n-1)
}

func funcrecursion() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d,", fibonacci3(i))
	}
	fmt.Print("\n")
}

// Circle 结构体定义和方法定义
type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

// c 需要使用指针 才能修改
func (c *Circle) setRadius(r float64) {
	c.radius = r
}

// Celsius Fahrenheit 温度类型定义和转换
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 冰点温度
	BoilingC      Celsius = 100     //沸水温度
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}


func functempconv() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingF-CtoF(FreezingC)) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var c Celsius       // 0
	var f Fahrenheit    // 0
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f)          // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"

	c = FtoC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}

// 结构体的使用
func funcmethod() {
	var c Circle = Circle{2}
	fmt.Println(c.radius, "getArea", c.getArea())
	c.setRadius(3)
	fmt.Println(c.radius, "getArea", c.getArea())
}

//数组的使用
func funcarrays() {
	// 声明数组
	var l1 [5]int
	fmt.Println(l1)

	//初始化数组
	l2 := [3]int{2, 1, 3}
	//初始化不定长数组
	l3 := [...]int{2, 1, 3, 0}
	fmt.Println(l2)
	fmt.Println(l3)
	//将 1,3索引初始化
	l4 := [5]int{1: 1, 3: 9}
	fmt.Println(l4)

	l1[2] = 50
	fmt.Println(l1)

	ml1 := [2][3]int{}
	fmt.Println(ml1)
	ml2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(ml2)
	fmt.Println(ml2[1][2])
}

//切片的使用
func funcslice() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(s1[:3])
	fmt.Println(s1[1:3])
	fmt.Println(s1[1:])
	fmt.Println(cap(s1)) // 返回的是数组切片分配的空间大小
	//fmt.Println(s1 != nil)

}

func funcmap() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

// range的使用
func funcrange() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	for i, c := range "aBcd" {
		fmt.Println(i, c) // 输出的是字符  ascii码 97 66 99 100
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k := range kvs {
		fmt.Printf("%s \n", k)
	}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println(kvs["a"])
}

// Phone 接口定义
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

//接口的使用
func funcinterface() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

// 针对入参不应该有问题的函数，使用panic设计
func sqrtx(f float64) float64 {
	f2, err := sqrt(f)
	if err != nil {
		strF := strconv.FormatFloat(f, 'f', -1, 64)
		panic("error: sqrt(" + strF + "): " + err.Error())
	}
	return f2
}

// DivideError 定义结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `Error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// Divide 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

//错误处理的使用
func funcerror() {

	v1, err1 := sqrt(64)
	fmt.Println("sqrt(64) = ", v1, err1)
	v2, err2 := sqrt(-64)
	fmt.Println("sqrt(-64) = ", v2, err2)
	v3 := sqrtx(64)
	fmt.Println("sqrtx(64) = ", v3)

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

//定义耗时处理函数
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// go 并发
func funcgo() {
	go say("world")
	say("hello")
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// channel 遍历的使用
func funcforchan() {
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

// Just can read ch in func
func chread(ch <-chan int) {
	fmt.Printf("Read: %d\n", <- ch)
}

// Just can write ch in func
func chwrite(ch chan <- int) {
	ch <- 22
}

func funccallc() {
	cstr := C.CString("Hello, world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}

// ch这个channel的buffer大小是1，所以会交替的为空或为满，
// 所以只有一个case可以进行下去，无论i是奇数或者偶数，它都会打印0 2 4 6 8
// 如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。
// 增加前一个例子的buffer大小会使其输出变得不确定，因为当buffer既不为满也不为空时，select语句的执行情况就像是抛硬币的行为一样是随机的。
func funcchselect() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x:= <- ch:
			fmt.Println(x)
		case ch <- i:

		}
	}
}

func SyntaxMain() {
	//prt()
	//prtG()
	//funcvar()
	//functempconv()
	//funcfmt()
	//funcref()
	//funcptr()
	//funccallf()
	//funcconstiota()
	//funccondition(-22)
	//funcselect()
	//funcparams("jim", "kk", 22, 2)
	//funcclosepk()
	//funcrecursion()
	//funcmethod()
	//funcarrays()
	//funcslice()
	//funcmap()
	//funcrange()
	//funcinterface()
	//funcerror()
	//funcgo()
	//funcforchan()
	//funccallc()
	funcchselect()
}
