package zdpgo_consul

import (
	"fmt"

	_ "github.com/zhangdapeng520/zdpgo_consul/resolver"
	"google.golang.org/grpc"
)

// GetGrpcClientConn 获取GRPC的客户端连接
// @param serviceName，consul中的grpc服务名称
func (c *Consul) GetGrpcClientConn(serviceName string) *grpc.ClientConn {
	// 使用负载均衡的方式获取grpc客户端
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", c.config.Host, c.config.Port, serviceName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		c.log.Error("使用负载均衡的方式获取grpc客户端失败：", err)
	}

	// 返回grpc客户端
	return conn
}
