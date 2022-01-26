package zdpgo_consul

import (
	"github.com/hashicorp/consul/api"
	"github.com/zhangdapeng520/zdpgo_zap"
)

// Consul consul核心对象
type Consul struct {
	config *ConsulConfig  // 配置对象
	log    *zdpgo_zap.Zap // 日志对象
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
	l := zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:        config.Debug,
		OpenGlobal:   true,
		OpenFileName: true,
		LogFilePath:  config.LogFilePath,
	})
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
