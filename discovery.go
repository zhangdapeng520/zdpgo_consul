package zgo_consul

import (
	"fmt"

	"google.golang.org/grpc"
	_ "zgo_consul/resolver"
)

// 获取GRPC的客户端连接
// 参数1：consulHost，consul的主机地址
// 参数2：consulPort，consul的端口号
// 参数3：serviceName，consul中的grpc服务名称
func GetGrpcClientConn(consulHost string, consulPort int, serviceName string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulHost, consulPort, serviceName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic(err)
	}
	return conn
}
