package e1

type NameRepository struct {
	names []string
}

func NewNameRepository() *NameRepository {
	return &NameRepository{
		names: make([]string, 0),
	}
}

func (this *NameRepository) GetIterator() func() (string, bool) {
	index := 0
	return func() (name string, ok bool) {
		if index >= len(this.names) {
			return
		}
		name, ok = this.names[index], true
		index++
		return
	}
}

func (p *NameRepository) SetName(name string) {
	p.names = append(p.names, name)
}
