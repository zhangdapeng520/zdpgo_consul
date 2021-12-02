package zgo_consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// 获取consul客户端
func GetConsulClient(consulHost string, consulPort int) *api.Client {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulHost, consulPort)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}

// 注册Grpc微服务到consul
func RegisterGrpc(
	client *api.Client, // consul端口号
	grpcHost string, // grpc地址
	grpcPort int, // grpc端口号
	grpcName string, // grpc名称
	grpcId string, // grpc ID
	grpcTags []string, // grpc 标签列表
) {
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", grpcHost, grpcPort),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = grpcName
	registration.ID = grpcId
	registration.Port = grpcPort
	registration.Tags = grpcTags
	registration.Address = grpcHost
	registration.Check = check

	// 注册服务
	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

// 从consul注销grpc服务
func DeRegister(client *api.Client, grpcId string) {
	if err := client.Agent().ServiceDeregister(grpcId); err != nil {
		panic(err)
	}
}
