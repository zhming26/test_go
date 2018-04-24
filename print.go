//当前程序的包名
package test_go

// 导入其他的包

import "fmt"
// import "os"
// import "time"
// import "strings"


/*
import (
    "fmt";
    "os";
    "time";
    "strings";
    "io"
)
*/

// 常量的定义
const PI = 3.14

// 全局变量的声明与赋值
var name = "gopher"

// 一般类型的声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由main 函数作为程序入口点启动
func main() {
    fmt.Println("Hello world!你好， 世界！")
}
