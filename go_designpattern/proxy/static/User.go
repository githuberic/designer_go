package static

import "fmt"

// User 用户
// @proxy IUser
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// 不实现细节
	fmt.Printf("username=%s,password=%s", username, password)
	return nil
}
