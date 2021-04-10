package error_handler

import "fmt"

if err != nil {
return fmt.Errorf("something failed: %v", err)
}

// 将错误包装在另一个错误中，同时保留原始内容：
type authorizationError struct {
	operation string
	err       error // original error
}
func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}


// 标准的访问方法 最好使用一个接口，比如 causer接口中实现 Cause() 方法来暴露原始错误
type causer interface {
	Cause() error
}
func (e *authorizationError) Cause() error {
	return e.err
}


import "github.com/pkg/errors"
//错误包装
if err != nil {
return errors.Wrap(err, "read failed")
}
// Cause接口
switch err := errors.Cause(err).(type) {
case *MyError:
// handle specifically
default:
// unknown error
}