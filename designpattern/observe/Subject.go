package observe

type Subject struct {
	observers []Observer
	state     int
}

func (p *Subject) GetState() int {
	return p.state
}

func (p *Subject) SetState(state int) {
	p.state = state
	p.NotifyAll()
}

func (p *Subject) Attach(observer Observer) {
	p.observers = append(p.observers, observer)
}

func (p *Subject) NotifyAll() {
	for _, observer := range p.observers {
		observer.Update()
	}
}
