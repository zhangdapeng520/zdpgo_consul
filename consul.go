package zdpgo_consul

import (
	"github.com/zhangdapeng520/zdpgo_consul/api"
)

// Consul consul核心对象
type Consul struct {
	config *ConsulConfig // 配置对象
	client *api.Client   // consul客户端对象
}

// New 创建consul的实例
// @param config consul配置对象
func New(config ConsulConfig) (*Consul, error) {
	c := Consul{}

	// 初始化配置
	c.config = &config

	// 初始化客户端
	var err error
	c.client, err = c.Client()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
