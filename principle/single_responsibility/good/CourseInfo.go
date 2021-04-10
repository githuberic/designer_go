package good

type CourseInfo struct {
	iID int
	sName string
}
func (p *CourseInfo) ID() int {
	return p.iID
}
func (p *CourseInfo) Name() string {
	return p.sName
}
