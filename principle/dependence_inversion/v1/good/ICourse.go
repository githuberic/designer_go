package good

type ICourse interface {
	ID() int
	Name() string

	SetUser(IUser)
	Study()
}
