package zdpgo_consul

import (
	"github.com/zhangdapeng520/zdpgo_consul/api"
)

// Consul consul核心对象
type Consul struct {
	Config *Config     // 配置对象
	client *api.Client // consul客户端对象
}

// New 创建consul的实例
// @param config consul配置对象
func New(config *Config) (*Consul, error) {
	c := Consul{}

	// 初始化配置
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 8500
	}
	c.Config = config

	// 初始化客户端
	var err error
	c.client, err = c.Client()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
