package good

type GoodUser struct {
	iID int
	sName string
}
func NewGoodUser(id int, name string) IUser {
	return &GoodUser{
		iID: id,
		sName: name,
	}
}

func (me *GoodUser) ID() int {
	return me.iID
}
func (me *GoodUser) Name() string {
	return me.sName
}
func (me *GoodUser) Study(course ICourse) {
	course.SetUser(me)
	course.Study()
}
