package zdpgo_consul

// ConsulConfig consul配置对象
type ConsulConfig struct {
	Debug       bool   // 是否为debug模式
	LogFilePath string // 日志存放路径
	Host        string // consul地址
	Port        uint16 // consul端口号
}

// ServiceConfig service服务配置对象
type ServiceConfig struct {
	Host string   // 地址
	Port uint16   // 端口号
	Name string   // 名称
	Id   string   // ID
	Tags []string // 标签列表
}
