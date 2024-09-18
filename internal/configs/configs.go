package configs

type Config struct {
	Host `yaml:"host" json:"host"`
	Key  `yaml:"key" json:"key"`
}

type Host struct {
	Auth string `yaml:"auth" json:"auth"`
}

type Key struct {
	ApiGateway string `yaml:"apiGateway" json:"apiGateway"`
}
