package zdpgo_consul

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_consul/api"
)

// Client 获取consul客户端
func (c *Consul) Client() (*api.Client, error) {
	// 使用默认的配置
	cfg := api.DefaultConfig()

	// consul的地址
	cfg.Address = fmt.Sprintf("%s:%d", c.Config.Host, c.Config.Port)

	// 创建consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// 返回Consul客户端
	return client, nil
}

// RegisterGrpc 注册Grpc微服务到consul
// @param host 地址
// @param port 端口号
// @param name 名称
// @param id ID
// @param tags 标签列表
// @param isGrpc 是否为grpc
func (c *Consul) RegisterGrpc(config ServiceConfig) error {
	//生成对应的检查对象
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	check := &api.AgentServiceCheck{
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
		GRPC:                           addr,
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = config.Name
	registration.ID = config.Id
	registration.Port = int(config.Port)
	registration.Tags = config.Tags
	registration.Address = config.Host
	registration.Check = check

	// 注册服务
	err := c.client.Agent().ServiceRegister(registration)
	return err
}

// RegisterHTTP 注册HTTP服务
func (c *Consul) RegisterHTTP(config WebConfig) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.Config.Host, c.Config.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}

	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           config.HealthUrl,
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = config.Name
	registration.ID = config.Id
	registration.Port = int(config.Port)
	registration.Tags = config.Tags
	registration.Address = config.Host
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	return err
}

// DeRegister 注销服务
func (c *Consul) DeRegister(serviceId string) error {
	return c.client.Agent().ServiceDeregister(serviceId)
}
