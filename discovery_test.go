package zdpgo_consul

import (
	"fmt"
	"testing"
)

// 测试服务发现
func TestConsul_GetGrpcClientConn(t *testing.T) {
	c, err := New(ConsulConfig{
		Debug: true,
		Host:  "127.0.0.1",
		Port:  8500,
	})
	if err != nil {
		t.Error(err)
	}

	conn, err := c.GetGrpcClientConn("test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(conn)
}
