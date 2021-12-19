package e1

import "container/list"

type Building struct {
	Apartments list.List
}
func (this Building) Accept(v Visitor) {
	v.VisitBuilding(this)
}
