package e2

type Proxy struct {
	real RealSubject
}

func (this *Proxy) Do() string {
	res := "pre:"
	res += this.real.Do()
	res += ":after"
	return res
}
