package utils

import "strconv"

// HasError 错误断言
// 当 err 不为 nil 时触发 panic
// 对于当前请求不会继续执行之后的代码，并通过 recover 机制，返回指定格式的信息
func HasError(err error, msg string, code ...int) {
	if err != nil {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		if msg != "" {
			msg = err.Error()
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// Assert 条件断言
// 当 断言条件 为 false 时触发 panic
// 对于当前请求不会继续执行之后的代码，并通过 recover 机制，返回指定格式的信息
func Assert(confition bool, msg string, code ...int) {
	if !confition {
		statusCode := 200
		if len(code) > 0 {
			statusCode = code[0]
		}
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}
