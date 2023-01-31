## Go 基础语法

第一个简单的Go程序：

```go
// 当前程序的包名
package main

// 导入其他包
import . "fmt"

// 常量定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由main函数作为程序入口点启动
func main() {
	Println("Hello World!")
}
```

运行方法;

- 执行`go run [filename].go`

- `go build [filename].go`, `./[filename]`生成二进制格式编译

> *package main* 定义了包名。

>*import "fmt"* 告诉 Go 编译器这个程序需要使用 fmt 包

>下一行 *func main()* 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

>`/*...*/` 是注释

> *fmt.Println(...)* 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。

>当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

- 需要注意的是 **{** 不能单独放在一行，所以以下代码在运行时会产生错误



###### 变量的声明

1.`var age int = xxx`

2.`v_name := value`

3.`var v_name = value`

```go
package main
import "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)
}
```

**多变量的声明**

```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误，而且只能在函数体内出现


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

**语言常量**

```go
const identifier [type] = value
```

###### iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量。iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

```go
const (
    a = iota
    b = iota
    c = iota
)
```

简写为

```go
const (
    a = iota
    b
    c
)
```



###### 输出

- **Sprintf** 根据格式化参数生成格式化的字符串并返回该字符串。
- **Printf** 根据格式化参数生成格式化的字符串并写入标准输出。



##### 判断/循环语句

`if`

```go
if a >= 0 {
    fmt.Println("Hello, World!")
} else {
    fmt.Println("Nice to meet you")
}
```

`for`循环

```go
for i := 1; i <= 2; i++ {
    fmt.Println(i)
}
```

`switch`

```go
b := 2
switch b {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3, 4:
    fmt.Println("big")
default:
    fmt.Println("None")
}
```



##### 数组

```go
var a [5]int
a[4] = 100
fmt.Println(a[4])

b := [5]int{1, 2, 3, 4, 5}
fmt.Println(b[2])

var g [5][5]int
g[2][3] = 21
fmt.Println(g[2][3])
```

> 真实的业务代码内并不使用数组，而是使用切片(即可变长数组)

##### 切片

```go
func main() {
   fmt.Println("Henry")

   s := make([]string, 3)
   s[0] = "a"
   s[1] = "b"
   s[2] = "c"
   fmt.Println(s[2], len(s))

   s = append(s, "d")
   s = append(s, "e", "f")
   fmt.Println(s)

   t := make([]string, len(s))
   copy(t, s)
   fmt.Println(t[:2], s[3:5], t[2:])

   good := []string{"a", "b", "c", "d"}
   fmt.Println(good)
}
```

##### Map

```go
func main() {
	fmt.Println("Henry")

	mp := make(map[string]int)
	mp["one"] = 1
	mp["two"] = 2
	fmt.Println(mp, len(mp), mp["one"], mp["two"])
	delete(mp, "one")

	mp2 := map[string]int{"one": 1, "two": 2}
	var mp3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(mp2, mp3)

	for key, val := range mp {
		fmt.Println(key, val)
	}
    // 单纯遍历就是key
	for key := range mp {
		fmt.Println(key)
	}
}
```



##### 函数

```go
func sum(a int, b int) int {
   return a + b
}

func mswap(a int, b int) (c int, d int) {
   c, d = b, a
   return c, d
}

func exists(m map[string]string, k string) (v string, ok bool) {
   v, ok = m[k]
   return v, ok
}

func main() {
   fmt.Println("Henry")
   a, b := 1, 2
   fmt.Println(a, b)
   c,d := mswap(a, b)
   fmt.Println(c, d)
}
```

> 函数可以返回多个值

##### 指针

指针类似于C++中的引用

```go
func mswap(a *int, b *int) {
   *b = *b ^ *a
   *a = *a ^ *b
   *b = *b ^ *a
}

func main() {
   fmt.Println("Henry")
   a, b := 1, 2
   fmt.Println(a, b)
   mswap(&a, &b)
   fmt.Println(a, b)
}
```



##### 字符串

- 字符串的连接使用``+``连接
- 以下为字符串常用函数

```go
type point struct {
	x, y int
}

func main() {
	fmt.Println("Henry")
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))   //true
	fmt.Println(strings.Count(a, "l"))       //2
	fmt.Println(strings.HasPrefix(a, "he"))  //true
	fmt.Println(strings.HasSuffix(a, "llo")) //true
	fmt.Println(strings.Index(a, "ell"))     //1

	fmt.Println(strings.Join([]string{"he", "ll"}, "[insert]")) //he[insert]ll
	fmt.Println(strings.Repeat(a, 2))                           //hellohello
	fmt.Println(strings.Replace(a, "e", "E", -1))               //hEllo  n<0表示替换所有old子串,n>0则将前n个不重叠的old子串替换
	fmt.Println(strings.Split("a-b-c", "-"))                    //[a b c]
	fmt.Println(strings.ToLower(a))                             //[a b c]
	fmt.Println(strings.ToUpper(a))                             //[a b c]
	b := "你好"
	fmt.Println(len(b))
	// 字符串格式化
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Printf("s=%v\n", s)
	fmt.Printf("n=%v\n", n)
	fmt.Printf("p=%+v\n", p) //更加详细
	fmt.Printf("p=%#v\n", p) //最详细
	f := 3.12312
	fmt.Println(f)
	fmt.Printf("%.2lf", f)
}
```



##### 结构体

```go
type user struct {
   name     string
   password string
}

func main() {
   fmt.Println("Henry")
   stu1 := user{name: "wyh", password: "12345"}
   fmt.Println(stu1.name, stu1.password)
}
```

###### 结构体函数

```go
type user struct {
	name     string
	password string
}

func (u user) checkname(name string) string {
	if u.name == name {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	fmt.Println("Henry")
	stu1 := user{name: "wyh", password: "12345"}
	fmt.Println(stu1.checkname("wyh"))
}
```

> 带指针可以对结构体进行修改，即
>
> `func (u *user) checkname(name string) string{}`



##### 异常处理

两种写法

```go
package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	fmt.Println("Henry")
	u, err := findUser([]user{{"wyh", "1024"}}, "wyh")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	if u, err := findUser([]user{{"wyh", "name"}}, "wyh"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u.name)
	}
}
```



##### Json

```go
type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
} //只要保证每一个字段的第一个字母大写，就可以进行Json序列化

func main() {
	a := userInfo{Name: "wyh", Age: 19, Hobby: []string{"Python", "C++"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf)

	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b) //反序列化
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)

	fmt.Println("Henry")
}
```

注意！

- 如果需要序列化成Json格式的结构体中，显示定义的首字母应该大写

```go
type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
} //只要保证每一个字段的第一个字母大写，就可以进行Json序列化
func main(){
	a := userInfo{Name: "wyh", Age: 19, Hobby: []string{"Python", "C++"}}
}	
```

输出:

```
{                                                                 
        "Name": "wyh",                                            
        "age": 19,                                                
        "Hobby": [                                                
                "Python",                                         
                "C++"                                             
        ]                                                         
}    
```

若改成

```go
type userInfo struct {
	name  string
	Age   int `json:"age"`
	Hobby []string
} //只要保证每一个字段的第一个字母大写，就可以进行Json序列化

func main() {
	a := userInfo{name: "wyh", Age: 19, Hobby: []string{"Python", "C++"}}
}
```

则输出为

```
{                                                              
        "age": 19,                                             
        "Hobby": [                                             
                "Python",                                      
                "C++"                                          
        ]                                                      
}    
```

##### 时间处理

```go
import (
	"errors"
	"fmt"
	"time"
)

func main() {
   now := time.Now()
   fmt.Println(now)
   t := time.Date(2022, 3, 27, 1, 25, 36, 10, time.UTC)
   fmt.Println(t)
   fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.YearDay())
   //2023-01-29 02:11:21.1448466 +0800 CST m=+0.005494101
   //2022-03-27 01:25:36.00000001 +0000 UTC
   //2022 March 27 1 25 86
   fmt.Println(t.Format("2002-03-27 17:06:06"))
   fmt.Println("Henry")
}
```





##### 数字解析

```go
func main() {
   n, err := strconv.ParseInt("072313", 0, 64) //数字 进制(0为自动推断 精度)
   fmt.Println(n, err)
   m, _ := strconv.Atoi("123")
   fmt.Println(m)
   fmt.Println("Henry")
}
```





#### 项目一：猜数字

简单，略

#### 项目二：在线词典

1. 复制`https://fanyi.caiyunapp.com/#/`控制台抓取的dict文件的cURL(bash)

2. 粘贴进`https://curlconverter.com/go/`自动转换成Go文件

3. 复制dict响应头在`https://oktools.net/json2go`中获得对应结构体
4. 利用Go结构体的特性分别将要查的单词序列化发送后反序列化输出即可

运行：`go run OnlineDic/OnlineDic.go [word] `

#### 项目三：SOCKET5代理服务器

运行`go run SCK5.go `

并在windows环境下截至netcat使用nc命令

![001120](E:\MyProject\ByteDance_Winter_Camp\img\001120.jpg)

<hr>

![010139](E:\MyProject\ByteDance_Winter_Camp\img\010139.jpg)

<hr>

![image-20230201014522554](E:\MyProject\ByteDance_Winter_Camp\img\014532.jpg)

<hr>

![image-20230201015716251](C:\Users\Henry\AppData\Roaming\Typora\typora-user-images\image-20230201015716251.png)

最终完工
