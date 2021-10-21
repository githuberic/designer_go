package resource_clear_test

import (
	"io"
	"log"
	"testing"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func Open(str string) (io.Closer, error) {
	return nil, nil
}

func TestClose(t *testing.T) {
	r, err := Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer Close(r) // 使用defer关键字在函数退出时关闭文件。

	r, err = Open("b")
	if err != nil {
		log.Fatalf("error opening 'b'\n")
	}
	defer Close(r) // 使用defer关键字在函数退出时关闭文件。
}
