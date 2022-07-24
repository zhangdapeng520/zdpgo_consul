package zdpgo_consul

// Config consul配置对象
type Config struct {
	Host string // consul地址
	Port uint16 // consul端口号
}

// ServiceConfig service服务配置对象
type ServiceConfig struct {
	Host string   // 地址
	Port uint16   // 端口号
	Name string   // 名称
	Id   string   // ID
	Tags []string // 标签列表
}

// WebConfig web服务配置对象
type WebConfig struct {
	Host      string   // 地址
	Port      uint16   // 端口号
	Name      string   // 名称
	Id        string   // ID
	HealthUrl string   // 健康检查地址
	Tags      []string // 标签列表
}

// DeregisterHTTPConfig 注销http服务的配置
type DeregisterHTTPConfig struct {
	ConsulHost string // consul主机地址
	ConsulPort uint16 // consul端口号
	ServerId   string // 服务id
}
