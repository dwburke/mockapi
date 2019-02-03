package config

var Config *ConfigType

func init() {
	Config = &ConfigType{}
}

type Route struct {
	Method     string `yaml:"method"`
	Result     string `yaml:"result"`
	ResultType string `yaml:"result_type"`
}

type ConfigType struct {
	Routes map[string]*Route `yaml:"routes"`
}
