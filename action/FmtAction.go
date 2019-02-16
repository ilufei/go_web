package action 

import (
	"fmt"
	"net/http"
)

func FmtExample(writer http.ResponseWriter, request *http.Request) {
    /*
	func Printf(format string, a ...interface{}) (n int, err error)
	Printf根据format参数生成格式化的字符串并写入标准输出。返回写入的字节数和遇到的任何错误。
	
	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	Fprintf根据format参数生成格式化的字符串并写入w。返回写入的字节数和遇到的任何错误。
	
	func Sprintf(format string, a ...interface{}) string
	Sprintf根据format参数生成格式化的字符串并返回该字符串。
	
	func Print(a ...interface{}) (n int, err error)
	Print采用默认格式将其参数格式化并写入标准输出。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。返回写入的字节数和遇到的任何错误。

	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	Fprint采用默认格式将其参数格式化并写入w。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。返回写入的字节数和遇到的任何错误。

	func Sprint(a ...interface{}) string
	Sprint采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。如果两个相邻的参数都不是字符串，会在它们的输出之间添加空格。

	func Println(a ...interface{}) (n int, err error)
	Println采用默认格式将其参数格式化并写入标准输出。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。返回写入的字节数和遇到的任何错误。

	func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	Fprintln采用默认格式将其参数格式化并写入w。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。返回写入的字节数和遇到的任何错误。

	func Sprintln(a ...interface{}) string
	Sprintln采用默认格式将其参数格式化，串联所有输出生成并返回一个字符串。总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。

	var err = fmt.Errorf("%s happened", thing)

	通用
		%v    值的默认格式表示
		%+v    类似%v，但输出结构体时会添加字段名
		%#v    值的Go语法表示
		%T    值的类型的Go语法表示
		%%    百分号
		
	布尔
		%t    单词true或false
		
	整数
		%b    表示为二进制
		%c    该值对应的unicode码值
		%d    表示为十进制
		%o    表示为八进制
		%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
		%x    表示为十六进制，使用a-f
		%X    表示为十六进制，使用A-F
		%U    表示为Unicode格式：U+1234，等价于"U+%04X"
		
	浮点数与复数的两个组分
		%b    无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
		%e    科学计数法，如-1234.456e+78
		%E    科学计数法，如-1234.456E+78
		%f    有小数部分但无指数部分，如123.456
		%F    等价于%f
		%g    根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
		%G    根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
		
		%f:    默认宽度，默认精度
		%9f    宽度9，默认精度
		%.2f   默认宽度，精度2
		%9.2f  宽度9，精度2
		%9.f   宽度9，精度0
		
	字符串和[]byte：
		%s    直接输出字符串或者[]byte
		%q    该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
		%x    每个字节用两字符十六进制数表示（使用a-f）
		%X    每个字节用两字符十六进制数表示（使用A-F）
		
	指针
		%p    表示为十六进制，并加上前导的0x
	*/
	fmt.Fprintf(writer, "fmt example")
}