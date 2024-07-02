package internal

type Config struct {
	Server struct {
		Port int    `yaml:"port" mapstructure:"port"`
		Host string `yaml:"host" mapstructure:"host"`
	} `yaml:"server" mapstructure:"server"`
	Database struct {
		User     string `yaml:"user" mapstructure:"user"`
		Password string `yaml:"password" mapstructure:"password"`
		Name     string `yaml:"name" mapstructure:"name"`
	} `yaml:"database" mapstructure:"database"`
	EnvName string `yaml:"env_name" mapstructure:"env_name"`
}
