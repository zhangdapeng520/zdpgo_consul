package zdpgo_consul

import (
	"fmt"
	"testing"
)

// 测试服务发现
func TestConsul_GetGrpcClientConn(t *testing.T) {
	c := New(ConsulConfig{
		Debug: true,
		Host:  "127.0.0.1",
		Port:  8500,
	})
	
	conn := c.GetGrpcClientConn("test")
	fmt.Println(conn)
}

