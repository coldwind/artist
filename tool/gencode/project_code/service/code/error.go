package code

var (
	ErrorNetwork OutputCode = &logicCode{
		Code: 1000,
		Msg:  "网络错误",
	}
	ErrorServer OutputCode = &logicCode{
		Code: 1001,
		Msg:  "服务器故障",
	}
	ErrorRateLimit OutputCode = &logicCode{
		Code: 1002,
		Msg:  "网络限流，请稍后再试",
	}
)
