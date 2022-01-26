package zdpgo_consul

import (
	"testing"
)

// 测试服务注册
func TestConsul_Register(t *testing.T) {
	c := New(ConsulConfig{
		Debug: true,
		Host:  "127.0.0.1",
		Port:  8500,
	})

	c.RegisterGrpc(ServiceConfig{
		Host: "192.168.33.14",
		Port: 8888,
		Name: "test",
		Id:   "test",
		Tags: []string{"zhangdapeng", "service", "user"},
	})
}

// 测试服务注销
func TestConsul_DeRegister(t *testing.T) {
	c := New(ConsulConfig{
		Debug: true,
		Host:  "127.0.0.1",
		Port:  8500,
	})

	c.DeRegisterGrpc("test")
}
