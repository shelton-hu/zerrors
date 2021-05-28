package zerrors

import (
	"fmt"
	"runtime"
)

type Zerror struct {
	// 错误代码
	Code int
	// 错误信息
	Message string
	// 错误详情
	Detail string
	// 错误调用信息
	Caller string
	// 父错误
	parent *Zerror
}

// 根据错误代码code，新构建一个*Zerror，传入了errs，则根据errs类型，会有以下三种行为：
//     1. if len(errs) == 0 || errs[0] == nil {
//            return &Zerror{
//                Code:    code,
//                Message: msger.ErrorMessage(code),
//                Caller:  caller,
//            }
//        }
//     2. if errs[0].(type) == *Zerror {
//            return &Zerror{
//                Code:    code,
//                Message: msger.ErrorMessage(code),
//                Caller:  caller,
//                parent:  errs[0],
//            }
//        }
//     3. if errs[0].(type) != *Zerror {
//            return &Zerror{
//                Code:    code,
//                Message: msger.ErrorMessage(code),
//                Detail:  errs[0].Error(),
//                Caller:  caller,
//            }
//        }
func New(code int, errs ...error) *Zerror {
	return newZerror(code, 2, errs...)
}

func NewWithCallerSkip(code int, skip int, errs ...error) *Zerror {
	return newZerror(code, skip, errs...)
}

func newZerror(code int, skip int, errs ...error) *Zerror {
	var caller string
	if _, file, line, ok := runtime.Caller(skip); ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	nberr := &Zerror{
		Code:    code,
		Message: msger.ErrorMessage(code),
		Caller:  caller,
	}

	if len(errs) == 0 || errs[0] == nil {
		return nberr
	}

	parent, ok := errs[0].(*Zerror)
	if ok {
		nberr.parent = parent
	} else {
		nberr.Detail = errs[0].Error()
	}

	return nberr
}

// 装饰一个err，根据传进来的err类型，会有以下三种行为：
//     1. if err == nil {
//            return &Zerror{
//                Code:    defaultCode,
//                Message: defaultMessage,
//                Caller:  caller,
//            }
//        }
//     2. if err.(type) == *Zerror {
//            return &Zerror{
//                Code:    err.Code,
//                Message: err.Message,
//                Caller:  caller,
//                parent:  err,
//            }
//        }
//     3. if err.(type) != *Zerror {
//            return &Zerror{
//                Code:    defaultCode,
//                Message: defaultMessage,
//                Detail:  err.Error(),
//                Caller:  caller,
//            }
//        }
func Wrapper(err error) *Zerror {
	return wrapper(err, 2)
}

func WrapperWithCallerSkip(err error, skip int) *Zerror {
	return wrapper(err, skip)
}

func wrapper(err error, skip int) *Zerror {
	var caller string
	if _, file, line, ok := runtime.Caller(skip); ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	defaultCode, defaultMessage := msger.DefaultMessage()

	nberr := &Zerror{
		Code:    defaultCode,
		Message: defaultMessage,
		Caller:  caller,
	}

	if err == nil {
		return nberr
	}

	parent, ok := err.(*Zerror)
	if ok {
		nberr.Code = parent.Code
		nberr.Message = parent.Message
		nberr.parent = parent
	} else {
		nberr.Detail = err.Error()
	}

	return nberr
}

func (e *Zerror) Error() string {
	if e == nil {
		return ""
	}
	if e.Detail != "" {
		return e.Detail
	}
	return e.Message
}

func (e *Zerror) Parent() *Zerror {
	return e.parent
}
