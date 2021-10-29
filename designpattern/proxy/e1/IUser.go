package e1

// IUser IUser
type IUser interface {
	Login(username, password string) error
}
