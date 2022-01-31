package zdpgo_consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// Client 获取consul客户端
func (c *Consul) Client() *api.Client {
	// 使用默认的配置
	cfg := api.DefaultConfig()

	// consul的地址
	cfg.Address = fmt.Sprintf("%s:%d", c.config.Host, c.config.Port)

	// 创建consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		c.log.Error("获取consul客户端失败：", err)
	}

	// 返回Consul客户端
	return client
}

// RegisterGrpc 注册Grpc微服务到consul
// @param host 地址
// @param port 端口号
// @param name 名称
// @param id ID
// @param tags 标签列表
// @param isGrpc 是否为grpc
func (c *Consul) RegisterGrpc(config ServiceConfig) {
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
	if err != nil {
		c.log.Error("注册服务到consul失败：", err)
	}
	c.log.Info("注册服务到consul成功")
}

// DeRegisterGrpc 从consul注销grpc服务
func (c *Consul) DeRegisterGrpc(id string) error {
	if err := c.client.Agent().ServiceDeregister(id); err != nil {
		c.log.Error("从consul注销grpc服务失败：", err)
		return err
	}
	return nil
}

// RegisterHTTP 注册HTTP服务
func (c *Consul) RegisterHTTP(config WebConfig) error {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", c.config.Host, c.config.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		c.log.Error("注册http服务失败", err)
	}

	// 生成对应的检查对象
	addr := fmt.Sprintf("http://%s:%d/health", config.Host, config.Port)
	c.log.Info("健康检查地址", "address", addr)
	check := &api.AgentServiceCheck{
		HTTP:                           addr,
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
	if err != nil {
		panic(err)
	}
	return nil
}

// DeRegisterHTTP 注销HTTP服务
func (c *Consul) DeRegisterHTTP(config DeregisterHTTPConfig) error {
	// 创建配置
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", config.ConsulHost, config.ConsulPort)

	// 创建consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		c.log.Error("创建consul客户端失败", "error", err.Error())
		return err
	}

	// 注销服务
	err = client.Agent().ServiceDeregister(config.ServerId)
	if err != nil {
		c.log.Error("注销HTTP服务失败", "error", err.Error())
	}
	return err
}
