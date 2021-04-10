package observe

import "fmt"

type Observer interface {
	Update()
}

type BinaryObserver struct {
	subject *Subject
}
func (p *BinaryObserver) Update() {
	fmt.Println("Binary String: ", p.subject.GetState())
}
func (p *BinaryObserver) BinaryObserver(subject *Subject) {
	p.subject = subject
	p.subject.Attach(p)
}

type OctalObserver struct {
	subject *Subject
}
func (p *OctalObserver) Update() {
	fmt.Println("Octal String:", p.subject.GetState())
}
func (p *OctalObserver) OctalObserver(subject *Subject) {
	p.subject = subject
	p.subject.Attach(p)
}

type HexObserver struct {
	subject *Subject
}
func (h *HexObserver) HexaObserver(subject *Subject) {
	h.subject = subject
	h.subject.Attach(h)
}
func (h *HexObserver) Update() {
	fmt.Println("Hex String:", h.subject.GetState())
}