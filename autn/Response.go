package autn

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息
func (res *Response) Instance(code int, message string) Response {
	return Response{
		Code: code,
		Msg:  message,
		Data: nil,
	}
}

func (res *Response) InstanceData(code int, message string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  message,
		Data: data,
	}
}
