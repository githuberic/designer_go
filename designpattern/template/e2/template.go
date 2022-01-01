package e2

import "fmt"

type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}
