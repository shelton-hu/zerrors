package zerrors

type ErrorStack struct {
	// 错误代码
	Code int `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 错误详情
	Detail string `json:"detail"`
	// 错误调用信息
	Caller string `json:"caller"`
}

// 将错误栈转化成切片返回
func PrintErrorStack(e *Zerror) []*ErrorStack {
	stacks := []*ErrorStack{}
	for e != nil {
		stack := &ErrorStack{
			Code:    e.Code,
			Message: e.Message,
			Caller:  e.Caller,
		}
		if e.Detail != "" {
			stack.Detail = e.Detail
		}
		stacks = append(stacks, stack)
		e = e.parent
	}
	return stacks
}
