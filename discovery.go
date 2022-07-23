package zdpgo_consul

import (
	"fmt"

	_ "github.com/zhangdapeng520/zdpgo_consul/resolver"
	"google.golang.org/grpc"
)

// GetGrpcClientConn 获取GRPC的客户端连接
// @param serviceName，consul中的grpc服务名称
func (c *Consul) GetGrpcClientConn(serviceName string) (*grpc.ClientConn, error) {
	// 使用负载均衡的方式获取grpc客户端
	addr := fmt.Sprintf("consul://%s:%d/%s?wait=14s", c.config.Host, c.config.Port, serviceName)
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 基于轮询的负载均衡
	)
	if err != nil {
		return nil, err
	}

	// 返回grpc客户端
	return conn, nil
}
