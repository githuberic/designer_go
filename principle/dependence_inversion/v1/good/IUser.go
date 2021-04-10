package good

type IUser interface {
	ID() int
	Name() string

	Study(ICourse)
}
