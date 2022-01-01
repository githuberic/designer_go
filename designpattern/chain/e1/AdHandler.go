package e1

import (
	"fmt"
	"strings"
)

type AdHandler struct {
	handler Handler
}

func (p *AdHandler) Handle(content string) {
	fmt.Println("执行广告过滤。。。")
	newContent := strings.Replace(content, "广告", "**", 1)
	fmt.Println(newContent)
	p.next(p.handler, newContent)
}

func (p *AdHandler) next(handler Handler, content string) {
	if p.handler != nil {
		p.handler.Handle(content)
	}
}
