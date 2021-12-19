package e3

import (
	"testing"
)

func TestVerify(t *testing.T) {
	proxy := &ProxyNetWork{}
	proxy.Conn("192.168.10.2")
	proxy.Close()
}
