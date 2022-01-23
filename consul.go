package zdpgo_consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/zhangdapeng520/zdpgo_log"
)

// Consul consul核心对象
type Consul struct {
	config *ConsulConfig  // 配置对象
	log    *zdpgo_log.Log // 日志对象
	client *api.Client    // consul客户端对象
}

// New 创建consul的实例
// @param config consul配置对象
func New(config ConsulConfig) *Consul {
	c := Consul{}

	// 初始化日志
	if config.LogFilePath == "" {
		config.LogFilePath = "zdpgo_consul.log"
	}
	logConfig := zdpgo_log.LogConfig{
		Debug: config.Debug,
		Path:  config.LogFilePath,
	}
	l := zdpgo_log.New(logConfig)
	c.log = l

	// 校验参数
	if config.Host == "" {
		c.log.Panic("consul主机地址不能为空")
	}
	if config.Port == 0 {
		c.log.Panic("consul端口号不能为空")
	}

	// 初始化配置
	c.config = &config

	// 初始化客户端
	c.client = c.Client()

	return &c
}
