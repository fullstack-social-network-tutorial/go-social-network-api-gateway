package configs

type Config struct {
	Address  `yaml:"address" json:"address"`
	Key      `yaml:"key" json:"key"`
	Inbound  `yaml:"inbound" json:"inbound"`
	Outbound `yaml:"outbound" json:"outbound"`
}

type Inbound struct {
	Auth string `yaml:"auth" json:"auth"`
}

type Outbound struct {
	Auth string `yaml:"auth" json:"auth"`
}

type Address struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type Key struct {
	ApiGateway string `yaml:"apiGateway" json:"apiGateway"`
}
